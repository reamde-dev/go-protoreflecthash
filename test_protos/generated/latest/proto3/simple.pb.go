// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: test_protos/schema/proto3/simple.proto

package proto3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_schema_proto3_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_schema_proto3_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_simple_proto_rawDescGZIP(), []int{0}
}

type Simple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoolField       bool        `protobuf:"varint,1,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	BytesField      []byte      `protobuf:"bytes,3,opt,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
	DoubleField     float64     `protobuf:"fixed64,5,opt,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	Fixed32Field    uint32      `protobuf:"fixed32,7,opt,name=fixed32_field,json=fixed32Field,proto3" json:"fixed32_field,omitempty"`
	Fixed64Field    uint64      `protobuf:"fixed64,9,opt,name=fixed64_field,json=fixed64Field,proto3" json:"fixed64_field,omitempty"`
	FloatField      float32     `protobuf:"fixed32,11,opt,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	Int32Field      int32       `protobuf:"varint,13,opt,name=int32_field,json=int32Field,proto3" json:"int32_field,omitempty"`
	Int64Field      int64       `protobuf:"varint,15,opt,name=int64_field,json=int64Field,proto3" json:"int64_field,omitempty"`
	Sfixed32Field   int32       `protobuf:"fixed32,17,opt,name=sfixed32_field,json=sfixed32Field,proto3" json:"sfixed32_field,omitempty"`
	Sfixed64Field   int64       `protobuf:"fixed64,19,opt,name=sfixed64_field,json=sfixed64Field,proto3" json:"sfixed64_field,omitempty"`
	Sint32Field     int32       `protobuf:"zigzag32,21,opt,name=sint32_field,json=sint32Field,proto3" json:"sint32_field,omitempty"`
	Sint64Field     int64       `protobuf:"zigzag64,23,opt,name=sint64_field,json=sint64Field,proto3" json:"sint64_field,omitempty"`
	StringField     string      `protobuf:"bytes,25,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	Uint32Field     uint32      `protobuf:"varint,27,opt,name=uint32_field,json=uint32Field,proto3" json:"uint32_field,omitempty"`
	Uint64Field     uint64      `protobuf:"varint,29,opt,name=uint64_field,json=uint64Field,proto3" json:"uint64_field,omitempty"`
	SimpleField     *Simple     `protobuf:"bytes,31,opt,name=simple_field,json=simpleField,proto3" json:"simple_field,omitempty"`
	RepetitiveField *Repetitive `protobuf:"bytes,33,opt,name=repetitive_field,json=repetitiveField,proto3" json:"repetitive_field,omitempty"`
	SingletonField  *Singleton  `protobuf:"bytes,35,opt,name=singleton_field,json=singletonField,proto3" json:"singleton_field,omitempty"`
}

func (x *Simple) Reset() {
	*x = Simple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_schema_proto3_simple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simple) ProtoMessage() {}

func (x *Simple) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_schema_proto3_simple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simple.ProtoReflect.Descriptor instead.
func (*Simple) Descriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_simple_proto_rawDescGZIP(), []int{1}
}

func (x *Simple) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *Simple) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *Simple) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *Simple) GetFixed32Field() uint32 {
	if x != nil {
		return x.Fixed32Field
	}
	return 0
}

func (x *Simple) GetFixed64Field() uint64 {
	if x != nil {
		return x.Fixed64Field
	}
	return 0
}

func (x *Simple) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *Simple) GetInt32Field() int32 {
	if x != nil {
		return x.Int32Field
	}
	return 0
}

func (x *Simple) GetInt64Field() int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *Simple) GetSfixed32Field() int32 {
	if x != nil {
		return x.Sfixed32Field
	}
	return 0
}

func (x *Simple) GetSfixed64Field() int64 {
	if x != nil {
		return x.Sfixed64Field
	}
	return 0
}

func (x *Simple) GetSint32Field() int32 {
	if x != nil {
		return x.Sint32Field
	}
	return 0
}

func (x *Simple) GetSint64Field() int64 {
	if x != nil {
		return x.Sint64Field
	}
	return 0
}

func (x *Simple) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *Simple) GetUint32Field() uint32 {
	if x != nil {
		return x.Uint32Field
	}
	return 0
}

func (x *Simple) GetUint64Field() uint64 {
	if x != nil {
		return x.Uint64Field
	}
	return 0
}

func (x *Simple) GetSimpleField() *Simple {
	if x != nil {
		return x.SimpleField
	}
	return nil
}

func (x *Simple) GetRepetitiveField() *Repetitive {
	if x != nil {
		return x.RepetitiveField
	}
	return nil
}

func (x *Simple) GetSingletonField() *Singleton {
	if x != nil {
		return x.SingletonField
	}
	return nil
}

type Repetitive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoolField       []bool        `protobuf:"varint,1,rep,packed,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	BytesField      [][]byte      `protobuf:"bytes,3,rep,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
	DoubleField     []float64     `protobuf:"fixed64,5,rep,packed,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	Fixed32Field    []uint32      `protobuf:"fixed32,7,rep,packed,name=fixed32_field,json=fixed32Field,proto3" json:"fixed32_field,omitempty"`
	Fixed64Field    []uint64      `protobuf:"fixed64,9,rep,packed,name=fixed64_field,json=fixed64Field,proto3" json:"fixed64_field,omitempty"`
	FloatField      []float32     `protobuf:"fixed32,11,rep,packed,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	Int32Field      []int32       `protobuf:"varint,13,rep,packed,name=int32_field,json=int32Field,proto3" json:"int32_field,omitempty"`
	Int64Field      []int64       `protobuf:"varint,15,rep,packed,name=int64_field,json=int64Field,proto3" json:"int64_field,omitempty"`
	Sfixed32Field   []int32       `protobuf:"fixed32,17,rep,packed,name=sfixed32_field,json=sfixed32Field,proto3" json:"sfixed32_field,omitempty"`
	Sfixed64Field   []int64       `protobuf:"fixed64,19,rep,packed,name=sfixed64_field,json=sfixed64Field,proto3" json:"sfixed64_field,omitempty"`
	Sint32Field     []int32       `protobuf:"zigzag32,21,rep,packed,name=sint32_field,json=sint32Field,proto3" json:"sint32_field,omitempty"`
	Sint64Field     []int64       `protobuf:"zigzag64,23,rep,packed,name=sint64_field,json=sint64Field,proto3" json:"sint64_field,omitempty"`
	StringField     []string      `protobuf:"bytes,25,rep,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	Uint32Field     []uint32      `protobuf:"varint,27,rep,packed,name=uint32_field,json=uint32Field,proto3" json:"uint32_field,omitempty"`
	Uint64Field     []uint64      `protobuf:"varint,29,rep,packed,name=uint64_field,json=uint64Field,proto3" json:"uint64_field,omitempty"`
	SimpleField     []*Simple     `protobuf:"bytes,31,rep,name=simple_field,json=simpleField,proto3" json:"simple_field,omitempty"`
	RepetitiveField []*Repetitive `protobuf:"bytes,33,rep,name=repetitive_field,json=repetitiveField,proto3" json:"repetitive_field,omitempty"`
	SingletonField  []*Singleton  `protobuf:"bytes,35,rep,name=singleton_field,json=singletonField,proto3" json:"singleton_field,omitempty"`
}

func (x *Repetitive) Reset() {
	*x = Repetitive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_schema_proto3_simple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repetitive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repetitive) ProtoMessage() {}

func (x *Repetitive) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_schema_proto3_simple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repetitive.ProtoReflect.Descriptor instead.
func (*Repetitive) Descriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_simple_proto_rawDescGZIP(), []int{2}
}

func (x *Repetitive) GetBoolField() []bool {
	if x != nil {
		return x.BoolField
	}
	return nil
}

func (x *Repetitive) GetBytesField() [][]byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *Repetitive) GetDoubleField() []float64 {
	if x != nil {
		return x.DoubleField
	}
	return nil
}

func (x *Repetitive) GetFixed32Field() []uint32 {
	if x != nil {
		return x.Fixed32Field
	}
	return nil
}

func (x *Repetitive) GetFixed64Field() []uint64 {
	if x != nil {
		return x.Fixed64Field
	}
	return nil
}

func (x *Repetitive) GetFloatField() []float32 {
	if x != nil {
		return x.FloatField
	}
	return nil
}

func (x *Repetitive) GetInt32Field() []int32 {
	if x != nil {
		return x.Int32Field
	}
	return nil
}

func (x *Repetitive) GetInt64Field() []int64 {
	if x != nil {
		return x.Int64Field
	}
	return nil
}

func (x *Repetitive) GetSfixed32Field() []int32 {
	if x != nil {
		return x.Sfixed32Field
	}
	return nil
}

func (x *Repetitive) GetSfixed64Field() []int64 {
	if x != nil {
		return x.Sfixed64Field
	}
	return nil
}

func (x *Repetitive) GetSint32Field() []int32 {
	if x != nil {
		return x.Sint32Field
	}
	return nil
}

func (x *Repetitive) GetSint64Field() []int64 {
	if x != nil {
		return x.Sint64Field
	}
	return nil
}

func (x *Repetitive) GetStringField() []string {
	if x != nil {
		return x.StringField
	}
	return nil
}

func (x *Repetitive) GetUint32Field() []uint32 {
	if x != nil {
		return x.Uint32Field
	}
	return nil
}

func (x *Repetitive) GetUint64Field() []uint64 {
	if x != nil {
		return x.Uint64Field
	}
	return nil
}

func (x *Repetitive) GetSimpleField() []*Simple {
	if x != nil {
		return x.SimpleField
	}
	return nil
}

func (x *Repetitive) GetRepetitiveField() []*Repetitive {
	if x != nil {
		return x.RepetitiveField
	}
	return nil
}

func (x *Repetitive) GetSingletonField() []*Singleton {
	if x != nil {
		return x.SingletonField
	}
	return nil
}

type Singleton struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Singleton:
	//
	//	*Singleton_TheBool
	//	*Singleton_TheBytes
	//	*Singleton_TheDouble
	//	*Singleton_TheFixed32
	//	*Singleton_TheFixed64
	//	*Singleton_TheFloat
	//	*Singleton_TheInt32
	//	*Singleton_TheInt64
	//	*Singleton_TheSfixed32
	//	*Singleton_TheSfixed64
	//	*Singleton_TheSint32
	//	*Singleton_TheSint64
	//	*Singleton_TheString
	//	*Singleton_TheUint32
	//	*Singleton_TheUint64
	//	*Singleton_TheSimple
	//	*Singleton_TheRepetitive
	//	*Singleton_TheSingleton
	Singleton isSingleton_Singleton `protobuf_oneof:"singleton"`
}

func (x *Singleton) Reset() {
	*x = Singleton{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_schema_proto3_simple_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Singleton) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Singleton) ProtoMessage() {}

func (x *Singleton) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_schema_proto3_simple_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Singleton.ProtoReflect.Descriptor instead.
func (*Singleton) Descriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_simple_proto_rawDescGZIP(), []int{3}
}

func (m *Singleton) GetSingleton() isSingleton_Singleton {
	if m != nil {
		return m.Singleton
	}
	return nil
}

func (x *Singleton) GetTheBool() bool {
	if x, ok := x.GetSingleton().(*Singleton_TheBool); ok {
		return x.TheBool
	}
	return false
}

func (x *Singleton) GetTheBytes() []byte {
	if x, ok := x.GetSingleton().(*Singleton_TheBytes); ok {
		return x.TheBytes
	}
	return nil
}

func (x *Singleton) GetTheDouble() float64 {
	if x, ok := x.GetSingleton().(*Singleton_TheDouble); ok {
		return x.TheDouble
	}
	return 0
}

func (x *Singleton) GetTheFixed32() uint32 {
	if x, ok := x.GetSingleton().(*Singleton_TheFixed32); ok {
		return x.TheFixed32
	}
	return 0
}

func (x *Singleton) GetTheFixed64() uint64 {
	if x, ok := x.GetSingleton().(*Singleton_TheFixed64); ok {
		return x.TheFixed64
	}
	return 0
}

func (x *Singleton) GetTheFloat() float32 {
	if x, ok := x.GetSingleton().(*Singleton_TheFloat); ok {
		return x.TheFloat
	}
	return 0
}

func (x *Singleton) GetTheInt32() int32 {
	if x, ok := x.GetSingleton().(*Singleton_TheInt32); ok {
		return x.TheInt32
	}
	return 0
}

func (x *Singleton) GetTheInt64() int64 {
	if x, ok := x.GetSingleton().(*Singleton_TheInt64); ok {
		return x.TheInt64
	}
	return 0
}

func (x *Singleton) GetTheSfixed32() int32 {
	if x, ok := x.GetSingleton().(*Singleton_TheSfixed32); ok {
		return x.TheSfixed32
	}
	return 0
}

func (x *Singleton) GetTheSfixed64() int64 {
	if x, ok := x.GetSingleton().(*Singleton_TheSfixed64); ok {
		return x.TheSfixed64
	}
	return 0
}

func (x *Singleton) GetTheSint32() int32 {
	if x, ok := x.GetSingleton().(*Singleton_TheSint32); ok {
		return x.TheSint32
	}
	return 0
}

func (x *Singleton) GetTheSint64() int64 {
	if x, ok := x.GetSingleton().(*Singleton_TheSint64); ok {
		return x.TheSint64
	}
	return 0
}

func (x *Singleton) GetTheString() string {
	if x, ok := x.GetSingleton().(*Singleton_TheString); ok {
		return x.TheString
	}
	return ""
}

func (x *Singleton) GetTheUint32() uint32 {
	if x, ok := x.GetSingleton().(*Singleton_TheUint32); ok {
		return x.TheUint32
	}
	return 0
}

func (x *Singleton) GetTheUint64() uint64 {
	if x, ok := x.GetSingleton().(*Singleton_TheUint64); ok {
		return x.TheUint64
	}
	return 0
}

func (x *Singleton) GetTheSimple() *Simple {
	if x, ok := x.GetSingleton().(*Singleton_TheSimple); ok {
		return x.TheSimple
	}
	return nil
}

func (x *Singleton) GetTheRepetitive() *Repetitive {
	if x, ok := x.GetSingleton().(*Singleton_TheRepetitive); ok {
		return x.TheRepetitive
	}
	return nil
}

func (x *Singleton) GetTheSingleton() *Singleton {
	if x, ok := x.GetSingleton().(*Singleton_TheSingleton); ok {
		return x.TheSingleton
	}
	return nil
}

type isSingleton_Singleton interface {
	isSingleton_Singleton()
}

type Singleton_TheBool struct {
	TheBool bool `protobuf:"varint,1,opt,name=the_bool,json=theBool,proto3,oneof"`
}

type Singleton_TheBytes struct {
	TheBytes []byte `protobuf:"bytes,3,opt,name=the_bytes,json=theBytes,proto3,oneof"`
}

type Singleton_TheDouble struct {
	TheDouble float64 `protobuf:"fixed64,5,opt,name=the_double,json=theDouble,proto3,oneof"`
}

type Singleton_TheFixed32 struct {
	TheFixed32 uint32 `protobuf:"fixed32,7,opt,name=the_fixed32,json=theFixed32,proto3,oneof"`
}

type Singleton_TheFixed64 struct {
	TheFixed64 uint64 `protobuf:"fixed64,9,opt,name=the_fixed64,json=theFixed64,proto3,oneof"`
}

type Singleton_TheFloat struct {
	TheFloat float32 `protobuf:"fixed32,11,opt,name=the_float,json=theFloat,proto3,oneof"`
}

type Singleton_TheInt32 struct {
	TheInt32 int32 `protobuf:"varint,13,opt,name=the_int32,json=theInt32,proto3,oneof"`
}

type Singleton_TheInt64 struct {
	TheInt64 int64 `protobuf:"varint,15,opt,name=the_int64,json=theInt64,proto3,oneof"`
}

type Singleton_TheSfixed32 struct {
	TheSfixed32 int32 `protobuf:"fixed32,17,opt,name=the_sfixed32,json=theSfixed32,proto3,oneof"`
}

type Singleton_TheSfixed64 struct {
	TheSfixed64 int64 `protobuf:"fixed64,19,opt,name=the_sfixed64,json=theSfixed64,proto3,oneof"`
}

type Singleton_TheSint32 struct {
	TheSint32 int32 `protobuf:"zigzag32,21,opt,name=the_sint32,json=theSint32,proto3,oneof"`
}

type Singleton_TheSint64 struct {
	TheSint64 int64 `protobuf:"zigzag64,23,opt,name=the_sint64,json=theSint64,proto3,oneof"`
}

type Singleton_TheString struct {
	TheString string `protobuf:"bytes,25,opt,name=the_string,json=theString,proto3,oneof"`
}

type Singleton_TheUint32 struct {
	TheUint32 uint32 `protobuf:"varint,27,opt,name=the_uint32,json=theUint32,proto3,oneof"`
}

type Singleton_TheUint64 struct {
	TheUint64 uint64 `protobuf:"varint,29,opt,name=the_uint64,json=theUint64,proto3,oneof"`
}

type Singleton_TheSimple struct {
	TheSimple *Simple `protobuf:"bytes,31,opt,name=the_simple,json=theSimple,proto3,oneof"`
}

type Singleton_TheRepetitive struct {
	TheRepetitive *Repetitive `protobuf:"bytes,33,opt,name=the_repetitive,json=theRepetitive,proto3,oneof"`
}

type Singleton_TheSingleton struct {
	TheSingleton *Singleton `protobuf:"bytes,35,opt,name=the_singleton,json=theSingleton,proto3,oneof"`
}

func (*Singleton_TheBool) isSingleton_Singleton() {}

func (*Singleton_TheBytes) isSingleton_Singleton() {}

func (*Singleton_TheDouble) isSingleton_Singleton() {}

func (*Singleton_TheFixed32) isSingleton_Singleton() {}

func (*Singleton_TheFixed64) isSingleton_Singleton() {}

func (*Singleton_TheFloat) isSingleton_Singleton() {}

func (*Singleton_TheInt32) isSingleton_Singleton() {}

func (*Singleton_TheInt64) isSingleton_Singleton() {}

func (*Singleton_TheSfixed32) isSingleton_Singleton() {}

func (*Singleton_TheSfixed64) isSingleton_Singleton() {}

func (*Singleton_TheSint32) isSingleton_Singleton() {}

func (*Singleton_TheSint64) isSingleton_Singleton() {}

func (*Singleton_TheString) isSingleton_Singleton() {}

func (*Singleton_TheUint32) isSingleton_Singleton() {}

func (*Singleton_TheUint64) isSingleton_Singleton() {}

func (*Singleton_TheSimple) isSingleton_Singleton() {}

func (*Singleton_TheRepetitive) isSingleton_Singleton() {}

func (*Singleton_TheSingleton) isSingleton_Singleton() {}

var File_test_protos_schema_proto3_simple_proto protoreflect.FileDescriptor

var file_test_protos_schema_proto3_simple_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2f, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0xd8, 0x05, 0x0a, 0x06, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0f, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x0b, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x52, 0x65, 0x70,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x76, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x41, 0x0a, 0x0f, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x23, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xdc, 0x05, 0x0a, 0x0a,
	0x52, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f,
	0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x01,
	0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x07, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x06, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x11, 0x20, 0x03,
	0x28, 0x0f, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x13, 0x20, 0x03, 0x28, 0x10, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x15, 0x20, 0x03, 0x28, 0x11, 0x52, 0x0b,
	0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x17, 0x20, 0x03, 0x28,
	0x12, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x19,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x44, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x21, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x52, 0x65, 0x70, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x76, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x41, 0x0a, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x74, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xc4, 0x05, 0x0a, 0x09, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x08, 0x74, 0x68, 0x65, 0x5f,
	0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x74, 0x68,
	0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x74, 0x68, 0x65, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x5f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x65, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x68, 0x65, 0x5f, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x07, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x68,
	0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x68, 0x65, 0x5f,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x09, 0x20, 0x01, 0x28, 0x06, 0x48, 0x00, 0x52,
	0x0a, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1d, 0x0a, 0x09, 0x74,
	0x68, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00,
	0x52, 0x08, 0x74, 0x68, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x68,
	0x65, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x08, 0x74, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x68, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08,
	0x74, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x23, 0x0a, 0x0c, 0x74, 0x68, 0x65, 0x5f,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0f, 0x48, 0x00,
	0x52, 0x0b, 0x74, 0x68, 0x65, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x23, 0x0a,
	0x0c, 0x74, 0x68, 0x65, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x10, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x12, 0x1f, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x65, 0x53, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x1f, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x17, 0x20, 0x01, 0x28, 0x12, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x65, 0x53, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x12, 0x1f, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x65, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x5f, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x65,
	0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1f, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x5f, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68,
	0x65, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x36, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x5f, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x42, 0x0a, 0x0e, 0x74, 0x68, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x74, 0x68, 0x65, 0x52, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x74, 0x6f, 0x6e, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x74, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x68, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x74, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_protos_schema_proto3_simple_proto_rawDescOnce sync.Once
	file_test_protos_schema_proto3_simple_proto_rawDescData = file_test_protos_schema_proto3_simple_proto_rawDesc
)

func file_test_protos_schema_proto3_simple_proto_rawDescGZIP() []byte {
	file_test_protos_schema_proto3_simple_proto_rawDescOnce.Do(func() {
		file_test_protos_schema_proto3_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_protos_schema_proto3_simple_proto_rawDescData)
	})
	return file_test_protos_schema_proto3_simple_proto_rawDescData
}

var file_test_protos_schema_proto3_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_protos_schema_proto3_simple_proto_goTypes = []interface{}{
	(*Empty)(nil),      // 0: schema.proto3.Empty
	(*Simple)(nil),     // 1: schema.proto3.Simple
	(*Repetitive)(nil), // 2: schema.proto3.Repetitive
	(*Singleton)(nil),  // 3: schema.proto3.Singleton
}
var file_test_protos_schema_proto3_simple_proto_depIdxs = []int32{
	1, // 0: schema.proto3.Simple.simple_field:type_name -> schema.proto3.Simple
	2, // 1: schema.proto3.Simple.repetitive_field:type_name -> schema.proto3.Repetitive
	3, // 2: schema.proto3.Simple.singleton_field:type_name -> schema.proto3.Singleton
	1, // 3: schema.proto3.Repetitive.simple_field:type_name -> schema.proto3.Simple
	2, // 4: schema.proto3.Repetitive.repetitive_field:type_name -> schema.proto3.Repetitive
	3, // 5: schema.proto3.Repetitive.singleton_field:type_name -> schema.proto3.Singleton
	1, // 6: schema.proto3.Singleton.the_simple:type_name -> schema.proto3.Simple
	2, // 7: schema.proto3.Singleton.the_repetitive:type_name -> schema.proto3.Repetitive
	3, // 8: schema.proto3.Singleton.the_singleton:type_name -> schema.proto3.Singleton
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_test_protos_schema_proto3_simple_proto_init() }
func file_test_protos_schema_proto3_simple_proto_init() {
	if File_test_protos_schema_proto3_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_protos_schema_proto3_simple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_protos_schema_proto3_simple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simple); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_protos_schema_proto3_simple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repetitive); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_protos_schema_proto3_simple_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Singleton); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_test_protos_schema_proto3_simple_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Singleton_TheBool)(nil),
		(*Singleton_TheBytes)(nil),
		(*Singleton_TheDouble)(nil),
		(*Singleton_TheFixed32)(nil),
		(*Singleton_TheFixed64)(nil),
		(*Singleton_TheFloat)(nil),
		(*Singleton_TheInt32)(nil),
		(*Singleton_TheInt64)(nil),
		(*Singleton_TheSfixed32)(nil),
		(*Singleton_TheSfixed64)(nil),
		(*Singleton_TheSint32)(nil),
		(*Singleton_TheSint64)(nil),
		(*Singleton_TheString)(nil),
		(*Singleton_TheUint32)(nil),
		(*Singleton_TheUint64)(nil),
		(*Singleton_TheSimple)(nil),
		(*Singleton_TheRepetitive)(nil),
		(*Singleton_TheSingleton)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_protos_schema_proto3_simple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_protos_schema_proto3_simple_proto_goTypes,
		DependencyIndexes: file_test_protos_schema_proto3_simple_proto_depIdxs,
		MessageInfos:      file_test_protos_schema_proto3_simple_proto_msgTypes,
	}.Build()
	File_test_protos_schema_proto3_simple_proto = out.File
	file_test_protos_schema_proto3_simple_proto_rawDesc = nil
	file_test_protos_schema_proto3_simple_proto_goTypes = nil
	file_test_protos_schema_proto3_simple_proto_depIdxs = nil
}
