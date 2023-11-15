package protoreflecthash

import (
	"bytes"
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const valueName = protoreflect.Name("value")

// Option modifies how hashes for protobufs is calculated.
type Option func(*hasher)

// ProtoHasher is an interface for hashers that are capable of returning an
// ObjectHash for protobufs.
type ProtoHasher interface {
	// HashProto returns the object hash of a given protocol buffer message.
	HashProto(msg protoreflect.Message) ([]byte, error)
}

// NewHasher creates a new ProtoHasher with the options specified in the
// argument.
func NewHasher(options ...Option) ProtoHasher {
	h := &hasher{}
	for _, opt := range options {
		opt(h)
	}
	return h
}

// MessageFullnameIdentifier is an option that uses the message descriptor
// fullname rather than a generic map identifier when hashing messages.
func MessageFullnameIdentifier() Option {
	return func(h *hasher) {
		h.messageFullnameIdentifier = true
	}
}

// FieldNamesAsKeys is an option that uses field names for key hashing rather
// than the field number.
func FieldNamesAsKeys() Option {
	return func(h *hasher) {
		h.fieldNamesAsKeys = true
	}
}

// SkipFields is an option that skips hashing of specific fields.
func SkipFields(fields ...string) Option {
	return func(h *hasher) {
		if h.skipFields == nil {
			h.skipFields = make(map[string]struct{})
		}
		for _, field := range fields {
			h.skipFields[field] = struct{}{}
		}
	}
}

type hasher struct {
	// Whether to use the proto field name as its key, as opposed to using the
	// tag number as the key.
	fieldNamesAsKeys bool
	// Whether to use the fullname of the message descriptor rather than 'm'
	// (mapIdentifier) for proto messages.
	messageFullnameIdentifier bool
	// Whether to skip hashing of specific fields. The full name of the field
	// is used as the key.
	skipFields map[string]struct{}
}

type fieldHashEntry struct {
	number int32
	khash  []byte
	vhash  []byte
}

// HashProto implements MessageHasher
func (h *hasher) HashProto(msg protoreflect.Message) ([]byte, error) {
	// Check if the value is nil.
	if msg == nil {
		return h.hashNil()
	}

	// Make sure the proto itself is actually valid (ie. can be marshalled).
	// If this fails, it probably means there are unset required fields or invalid
	// values.
	if _, err := proto.Marshal(msg.Interface()); err != nil {
		return nil, err
	}

	return h.hashMessage(msg)
}

func (h *hasher) hashMessage(msg protoreflect.Message) ([]byte, error) {
	if msg == nil {
		return hashNil()
	}

	md := msg.Descriptor()

	if hash, err, ok := h.hashWellKnownType(md, msg); ok {
		return hash, err
	}

	// TOOD(pcj): what is the correct handling of placeholder types?
	if md.IsPlaceholder() {
		return nil, nil
	}

	var hashes []*fieldHashEntry

	fieldHashes, err := h.hashFields(msg, md.Fields())
	if err != nil {
		return nil, fmt.Errorf("hashing fields: %w", err)
	}
	hashes = append(hashes, fieldHashes...)

	sort.Slice(hashes, func(i, j int) bool {
		return hashes[i].number < hashes[j].number
	})

	var buf bytes.Buffer
	for _, hash := range hashes {
		buf.Write(hash.khash)
		buf.Write(hash.vhash)
	}

	identifier := mapIdentifier
	if h.messageFullnameIdentifier {
		identifier = string(md.FullName())
	}

	return hash(identifier, buf.Bytes())
}

func (h *hasher) hashFields(msg protoreflect.Message, fields protoreflect.FieldDescriptors) ([]*fieldHashEntry, error) {
	hashes := make([]*fieldHashEntry, 0, fields.Len())

	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if !msg.Has(fd) {
			// if we are in this block and the field is a scalar one, it is
			// either a proto3 field that was never set or is the empty value
			// (indistinguishable) or this is a proto2 field that is nil.
			continue
		}
		if _, ok := h.skipFields[string(fd.FullName())]; ok {
			continue
		}
		hash, err := h.hashField(fd, msg.Get(fd))
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, hash)
	}

	return hashes, nil
}

func (h *hasher) hashField(fd protoreflect.FieldDescriptor, value protoreflect.Value) (*fieldHashEntry, error) {
	khash, err := h.hashFieldKey(fd)
	if err != nil {
		return nil, fmt.Errorf("hashing field key %d (%s): %w", fd.Number(), fd.FullName(), err)
	}

	vhash, err := h.hashFieldValue(fd, value)
	if err != nil {
		return nil, fmt.Errorf("hashing field value %d (%s): %w", fd.Number(), fd.FullName(), err)
	}

	return &fieldHashEntry{
		number: int32(fd.Number()),
		khash:  khash,
		vhash:  vhash,
	}, nil
}

func (h *hasher) hashFieldKey(fd protoreflect.FieldDescriptor) ([]byte, error) {
	if h.fieldNamesAsKeys {
		return hashUnicode(string(fd.Name()))
	}
	return hashInt64(int64(fd.Number()))
}

func (h *hasher) hashFieldValue(fd protoreflect.FieldDescriptor, value protoreflect.Value) ([]byte, error) {
	if fd.IsList() {
		return h.hashList(fd.Kind(), value.List())
	}
	if fd.IsMap() {
		return h.hashMap(fd.MapKey(), fd.MapValue(), value.Map())
	}
	return h.hashValue(fd.Kind(), value)
}

func (h *hasher) hashValue(kind protoreflect.Kind, value protoreflect.Value) ([]byte, error) {
	switch kind {
	case
		protoreflect.BoolKind:
		return h.hashBool(value.Bool())
	case
		protoreflect.EnumKind:
		return h.hashEnum(value.Enum())
	case
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Fixed64Kind:
		return h.hashUint(value.Uint())
	case
		protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sfixed64Kind:
		return h.hashInt(value.Int())
	case
		protoreflect.FloatKind,
		protoreflect.DoubleKind:
		return h.hashFloat(value.Float())
	case
		protoreflect.StringKind:
		return h.hashString(value.String())
	case
		protoreflect.BytesKind:
		return h.hashBytes(value.Bytes())
	case
		protoreflect.MessageKind:
		return h.hashMessage(value.Message())
	case
		protoreflect.GroupKind:
		return nil, fmt.Errorf("protoreflect.GroupKind: not implemented: %T", value)
	}
	return nil, fmt.Errorf("unexpected field kind: %v (%T)", kind, value)
}

func (h *hasher) hashNil() ([]byte, error) {
	return hashNil()
}

func (h *hasher) hashBool(value bool) ([]byte, error) {
	return hashBool(value)
}

func (h *hasher) hashEnum(value protoreflect.EnumNumber) ([]byte, error) {
	return hashInt64(int64(value))
}

func (h *hasher) hashInt(value int64) ([]byte, error) {
	return hashInt64(value)
}

func (h *hasher) hashUint(value uint64) ([]byte, error) {
	return hashUint64(value)
}

func (h *hasher) hashFloat(value float64) ([]byte, error) {
	return hashFloat(value)
}

func (h *hasher) hashString(value string) ([]byte, error) {
	return hashUnicode(value)
}

func (h *hasher) hashBytes(value []byte) ([]byte, error) {
	return hashBytes(value)
}

func (h *hasher) hashList(kind protoreflect.Kind, list protoreflect.List) ([]byte, error) {
	var buf bytes.Buffer

	for i := 0; i < list.Len(); i++ {
		value := list.Get(i)
		data, err := h.hashValue(kind, value)
		if err != nil {
			return nil, fmt.Errorf("hashing list item %d: %w", i, err)
		}
		buf.Write(data)
	}

	return hash(listIdentifier, buf.Bytes())
}

func (h *hasher) hashMap(kd, fd protoreflect.FieldDescriptor, m protoreflect.Map) ([]byte, error) {

	var mapHashEntries []hashMapEntry

	var errValue error
	var errKey protoreflect.MapKey
	m.Range(func(mk protoreflect.MapKey, v protoreflect.Value) bool {
		khash, err := h.hashFieldValue(kd, mk.Value())
		if err != nil {
			errKey = mk
			errValue = err
			return false
		}

		vhash, err := h.hashFieldValue(fd, v)
		if err != nil {
			errKey = mk
			errValue = err
			return false
		}

		mapHashEntries = append(mapHashEntries, hashMapEntry{
			khash: khash,
			vhash: vhash,
		})

		return true
	})
	if errValue != nil {
		return nil, fmt.Errorf("hashing map key %v: %w", errKey, errValue)
	}

	sort.Sort(byKHash(mapHashEntries))

	var buf bytes.Buffer
	for _, e := range mapHashEntries {
		buf.Write(e.khash[:])
		buf.Write(e.vhash[:])
	}

	return hash(mapIdentifier, buf.Bytes())
}

func (h *hasher) hashWellKnownType(md protoreflect.MessageDescriptor, msg protoreflect.Message) (hash []byte, err error, ok bool) {
	fullName := md.FullName()
	switch fullName {
	case protoreflect.FullName("google.protobuf.Any"):
		hash, err = h.hashGoogleProtobufAny(md, msg)
	case protoreflect.FullName("google.protobuf.BoolValue"):
		hash, err = h.hashGoogleProtobufBoolValue(md, msg)
	case protoreflect.FullName("google.protobuf.DoubleValue"):
		hash, err = h.hashGoogleProtobufDoubleValue(md, msg)
	case protoreflect.FullName("google.protobuf.Duration"):
		hash, err = h.hashGoogleProtobufDuration(md, msg)
	case protoreflect.FullName("google.protobuf.FloatValue"):
		hash, err = h.hashGoogleProtobufFloatValue(md, msg)
	case protoreflect.FullName("google.protobuf.Int32Value"):
		hash, err = h.hashGoogleProtobufInt32Value(md, msg)
	case protoreflect.FullName("google.protobuf.ListValue"):
		hash, err = h.hashGoogleProtobufListValue(md, msg)
	case protoreflect.FullName("google.protobuf.Int64Value"):
		hash, err = h.hashGoogleProtobufInt64Value(md, msg)
	case protoreflect.FullName("google.protobuf.NullValue"):
		hash, err = h.hashGoogleProtobufNullValue(md, msg)
	case protoreflect.FullName("google.protobuf.StringValue"):
		hash, err = h.hashGoogleProtobufStringValue(md, msg)
	case protoreflect.FullName("google.protobuf.Struct"):
		hash, err = h.hashGoogleProtobufStruct(md, msg)
	case protoreflect.FullName("google.protobuf.Timestamp"):
		hash, err = h.hashGoogleProtobufTimestamp(md, msg)
	case protoreflect.FullName("google.protobuf.UInt32Value"):
		hash, err = h.hashGoogleProtobufUint32Value(md, msg)
	case protoreflect.FullName("google.protobuf.UInt64Value"):
		hash, err = h.hashGoogleProtobufUint64Value(md, msg)
	case protoreflect.FullName("google.protobuf.Value"):
		hash, err = h.hashGoogleProtobufValue(md, msg)
	default:
		return nil, nil, false // no special handling needed, use hashMessage
	}
	return hash, err, true
}

func (h *hasher) hashGoogleProtobufAny(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	// files := protoregistry.GlobalFiles - create option to set files explictly?
	typeUrl := msg.Get(md.Fields().ByName("type_url")).String()
	// TODO: lookup at a type server?
	return nil, fmt.Errorf("protoreflecthash does not support hashing of Any type: " + typeUrl)
}

func (h *hasher) hashGoogleProtobufDuration(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashFieldsByName(md, msg, "seconds", "nanos")
}

func (h *hasher) hashGoogleProtobufTimestamp(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashFieldsByName(md, msg, "seconds", "nanos")
}

func (h *hasher) hashFieldsByName(md protoreflect.MessageDescriptor, msg protoreflect.Message, names ...string) ([]byte, error) {
	var buf bytes.Buffer

	for _, name := range names {
		value := msg.Get(md.Fields().ByName(protoreflect.Name(name)))
		data, err := h.hashValue(protoreflect.Int32Kind, value)
		if err != nil {
			return nil, fmt.Errorf("hashing %s: %w", md.FullName(), err)
		}
		buf.Write(data)
	}

	return hash(listIdentifier, buf.Bytes())
}

func (h *hasher) hashGoogleProtobufDoubleValue(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashFloat(msg.Get(md.Fields().ByName(valueName)).Float())
}

func (h *hasher) hashGoogleProtobufFloatValue(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashFloat(msg.Get(md.Fields().ByName(valueName)).Float())
}

func (h *hasher) hashGoogleProtobufInt32Value(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashInt(msg.Get(md.Fields().ByName(valueName)).Int())
}

func (h *hasher) hashGoogleProtobufInt64Value(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashInt(msg.Get(md.Fields().ByName(valueName)).Int())
}

func (h *hasher) hashGoogleProtobufUint32Value(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashUint(msg.Get(md.Fields().ByName(valueName)).Uint())
}

func (h *hasher) hashGoogleProtobufUint64Value(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashUint(msg.Get(md.Fields().ByName(valueName)).Uint())
}

func (h *hasher) hashGoogleProtobufBoolValue(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashBool(msg.Get(md.Fields().ByName(valueName)).Bool())
}

func (h *hasher) hashGoogleProtobufStringValue(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return h.hashString(msg.Get(md.Fields().ByName(valueName)).String())
}

func (h *hasher) hashGoogleProtobufValue(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	od := md.Oneofs().ByName("kind")
	fd := msg.WhichOneof(od)
	if fd == nil {
		return nil, fmt.Errorf("invalid struct value: one value must be populated")
	}
	value := msg.Get(fd)

	switch fd.Name() {
	case "null_value":
		return hashNil()
	case "number_value":
		return h.hashFloat(value.Float())
	case "string_value":
		return h.hashString(value.String())
	case "bool_value":
		return h.hashBool(value.Bool())
	case "struct_value":
		return h.hashGoogleProtobufStruct(value.Message().Descriptor(), value.Message())
	case "list_value":
		return h.hashGoogleProtobufListValue(value.Message().Descriptor(), value.Message())
	default:
		return nil, fmt.Errorf("unexpected struct value kind: %s", fd.Name())
	}
}

func (h *hasher) hashGoogleProtobufListValue(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	list := msg.Get(md.Fields().ByName("values")).List()

	var buf bytes.Buffer
	for i := 0; i < list.Len(); i++ {
		value := list.Get(i)
		data, err := h.hashMessage(value.Message())
		if err != nil {
			return nil, fmt.Errorf("hashing list item %d: %w", i, err)
		}
		buf.Write(data)
	}

	return hash(listIdentifier, buf.Bytes())
}

func (h *hasher) hashGoogleProtobufNullValue(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	return hashNil()
}

func (h *hasher) hashGoogleProtobufStruct(md protoreflect.MessageDescriptor, msg protoreflect.Message) ([]byte, error) {
	m := msg.Get(md.Fields().ByName("fields")).Map()
	var mapHashEntries []hashMapEntry

	var errValue error
	var errKey protoreflect.MapKey
	m.Range(func(mk protoreflect.MapKey, v protoreflect.Value) bool {
		khash, err := h.hashString(mk.String())
		if err != nil {
			errKey = mk
			errValue = err
			return false
		}

		vhash, err := h.hashMessage(v.Message())
		if err != nil {
			errKey = mk
			errValue = err
			return false
		}

		mapHashEntries = append(mapHashEntries, hashMapEntry{
			khash: khash,
			vhash: vhash,
		})

		return true
	})
	if errValue != nil {
		return nil, fmt.Errorf("hashing map key %v: %w", errKey, errValue)
	}

	sort.Sort(byKHash(mapHashEntries))

	var buf bytes.Buffer
	for _, e := range mapHashEntries {
		buf.Write(e.khash[:])
		buf.Write(e.vhash[:])
	}

	return hash(mapIdentifier, buf.Bytes())
}

type hashMapEntry struct {
	khash []byte
	vhash []byte
}

type byKHash []hashMapEntry

func (h byKHash) Len() int {
	return len(h)
}

func (h byKHash) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h byKHash) Less(i, j int) bool {
	return bytes.Compare(h[i].khash[:], h[j].khash[:]) < 0
}
