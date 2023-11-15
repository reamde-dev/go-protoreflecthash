package protoreflecthash

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"testing"

	"github.com/benlaurie/objecthash/go/objecthash"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb2_latest "github.com/stackb/protoreflecthash/test_protos/generated/latest/proto2"
	pb3_latest "github.com/stackb/protoreflecthash/test_protos/generated/latest/proto3"
)

//go:embed testdata/protoset.pb
var testProtoset []byte

func TestHashNil(t *testing.T) {
	for name, tc := range map[string]struct {
		value interface{}
		want  string
	}{
		"nil": {
			want: "1b16b1df538ba12dc3f97edbb85caa7050d46c148134290feba80f8236c83db9",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashNil()
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashBool(t *testing.T) {
	for name, tc := range map[string]struct {
		value bool
		want  string
	}{
		"false": {
			value: false,
			want:  "c02c0b965e023abee808f2b548d8d5193a8b5229be6f3121a6f16e2d41a449b3",
		},
		"true": {
			value: true,
			want:  "7dc96f776c8423e57a2785489a3f9c43fb6e756876d6ad9a9cac4aa4e72ec193",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashBool(tc.value)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashInt(t *testing.T) {
	for name, tc := range map[string]struct {
		value int64
		want  string
	}{
		"zero": {
			value: 0,
			want:  "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		"positive": {
			value: 1,
			want:  "4cd9b7672d7fbee8fb51fb1e049f690342035f543a8efe734b7b5ffb0c154a45",
		},
		"min": {
			value: math.MinInt,
			want:  "2df43a3eaece5bb912a43ce16ebdf392e1dd9ce14c16255783ca1f5456d7d04f",
		},
		"max": {
			value: math.MaxInt,
			want:  "eda7a99bc44462f5181f63a88e2ab9d8d318d68c2c2bf9ff70d9e4ecc2a99df3",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashInt(tc.value)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashUint(t *testing.T) {
	for name, tc := range map[string]struct {
		value uint64
		want  string
	}{
		"zero": {
			value: 0,
			want:  "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		"positive": {
			value: 1,
			want:  "4cd9b7672d7fbee8fb51fb1e049f690342035f543a8efe734b7b5ffb0c154a45",
		},
		"max": {
			value: math.MaxUint,
			want:  "5b50a7751238c21772625d9807fc62e2d25ae5bd092d2018f0834d871c5db302",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashUint(tc.value)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashFloat(t *testing.T) {
	for name, tc := range map[string]struct {
		value float64
		want  string
	}{
		"zero": {
			value: 0,
			want:  "60101d8c9cb988411468e38909571f357daa67bff5a7b0a3f9ae295cd4aba33d",
		},
		"neg": {
			value: -1.0,
			want:  "f706daa44d7e40e21ea202c36119057924bb28a49949d8ddaa9c8c3c9367e602",
		},
		"pos": {
			value: +1.0,
			want:  "f01adc732390ab024d64080e0b173f0ee3a1610efbdd4ce2a13bbf8d9b26c639",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashFloat(tc.value)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashString(t *testing.T) {
	for name, tc := range map[string]struct {
		value string
		want  string
	}{
		"zero": {
			value: "",
			want:  "0bfe935e70c321c7ca3afc75ce0d0ca2f98b5422e008bb31c00c6d7f1f1c0ad6",
		},
		"ascii": {
			value: "bob",
			want:  "5ef421eb52293e5e3919d3c6f08413b873129dd859f4d0ff8273e13a494b9e9e",
		},
		"unicode": {
			value: "你好",
			want:  "462b68f5e3d75aed5f02841b4ffee068d4cf33ce1b155105b71a9e5f358026df",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashString(tc.value)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashBytes(t *testing.T) {
	for name, tc := range map[string]struct {
		value []byte
		want  string
	}{
		"zero": {
			value: []byte{},
			want:  "454349e422f05297191ead13e21d3db520e5abef52055e4964b82fb213f593a1",
		},
		"ascii": {
			value: []byte("bob"),
			want:  "aa75ac53926e8b0711ee730b4c0d8b232b167180f843da40d6e75871cd0785a5",
		},
		"unicode": {
			value: []byte("你好"),
			want:  "39fafdc74a5ee3ff86bd0b982304e58f4685767e87f5176307df9c9e1cf50925",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashBytes(tc.value)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashList(t *testing.T) {
	for name, tc := range map[string]struct {
		kind  protoreflect.Kind
		value protoreflect.List
		want  string
	}{
		"zero": {
			kind:  protoreflect.StringKind,
			value: stringList{},
			want:  "acac86c0e609ca906f632b0e2dacccb2b77d22b0621f20ebece1a4835b93f6f0",
		},
		"foobar": {
			kind:  protoreflect.StringKind,
			value: stringList{"foo", "bar"},
			want:  "32ae896c413cfdc79eec68be9139c86ded8b279238467c216cf2bec4d5f1e4a2",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashList(tc.kind, tc.value)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashMap(t *testing.T) {
	for name, tc := range map[string]struct {
		value        proto.Message
		mapFieldName string
		obj          interface{}
		json         string
		want         string
	}{
		"IntMaps.int_to_string": {
			value:        &pb3_latest.IntMaps{IntToString: map[int64]string{0: "ZERO"}},
			mapFieldName: "int_to_string",
			obj:          map[int64]string{0: "ZERO"},
			// json:         `{0:"ZERO"}`, // can't use json representation in this case
			want: "8cda73a524d09ce6fa10b071cacd4c725521b660ee4a546b6ebdbf139370e9b9",
		},
		"StringMaps.string_to_bool": {
			value:        &pb3_latest.StringMaps{StringToBool: map[string]bool{"true": true}},
			mapFieldName: "string_to_bool",
			obj:          map[string]bool{"true": true},
			json:         `{"true":true}`,
			want:         "d84d7d0593f90628672ccc4fbc89e31c51a847f45f39d98b95ea032c8de25e64",
		},
		"StringMaps.string_to_string": {
			value:        &pb3_latest.StringMaps{StringToString: map[string]string{"foo": "bar"}},
			mapFieldName: "string_to_string",
			obj:          map[string]string{"foo": "bar"},
			json:         `{"foo":"bar"}`,
			want:         "7ef5237c3027d6c58100afadf37796b3d351025cf28038280147d42fdc53b960",
		},
		"StringMaps.string_to_string_k123": {
			value:        &pb3_latest.StringMaps{StringToString: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}},
			mapFieldName: "string_to_string",
			obj:          map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
			json:         `{"k1":"v1","k2":"v2","k3":"v3"}`,
			want:         "ddd65f1f7568269a30df7cafc26044537dc2f02a1a0d830da61762fc3e687057",
		},
		"StringMaps.string_to_string_k213": {
			value:        &pb3_latest.StringMaps{StringToString: map[string]string{"k2": "v2", "k1": "v1", "k3": "v3"}},
			mapFieldName: "string_to_string",
			obj:          map[string]string{"k2": "v2", "k1": "v1", "k3": "v3"},
			json:         `{"k1":"v1","k2":"v2","k3":"v3"}`,
			want:         "ddd65f1f7568269a30df7cafc26044537dc2f02a1a0d830da61762fc3e687057",
		},
		//
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			msg := tc.value.ProtoReflect()
			fd := msg.Descriptor().Fields().ByName(protoreflect.Name(tc.mapFieldName))

			got := getHash(t, func() ([]byte, error) {
				return h.hashMap(fd.MapKey(), fd.MapValue(), msg.Get(fd).Map())
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}
			if tc.json != "" {
				if diff := cmp.Diff(tc.want, jsonHash(t, tc.json)); diff != "" {
					t.Errorf("jsonhash (-want +got):\n%s", diff)
				}
			}
			if tc.obj != nil {
				if diff := cmp.Diff(tc.want, objectHash(t, tc.obj)); diff != "" {
					t.Errorf("objecthash (-want +got):\n%s", diff)
				}
				if tc.json != "" {
					if diff := cmp.Diff(tc.json, jsonString(t, tc.obj)); diff != "" {
						t.Errorf("jsonstring (-want +got):\n%s", diff)
					}
				}
			}
		})
	}
}

func TestHashEnum(t *testing.T) {
	for name, tc := range map[string]struct {
		value protoreflect.EnumNumber
		want  string
	}{
		"zero": {
			value: 0,
			want:  "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		"earth": {
			value: pb3_latest.PlanetV1_EARTH_V1.Number(),
			want:  "9a83c6cb1126d93de4a30715b28f1f4b26b983c57fb39e6d826d7e893ae4ee74",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashEnum(tc.value)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, objectHash(t, tc.value)); diff != "" {
				t.Errorf("objecthash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashEmpty(t *testing.T) {
	emptyHash := "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4"
	emptyJson := "{}"

	if diff := cmp.Diff(emptyHash, jsonHash(t, emptyJson)); diff != "" {
		t.Errorf("jsonhash (-want +got):\n%s", diff)
	}

	for _, tc := range []struct {
		msg proto.Message
	}{
		{&pb2_latest.Empty{}},
		{&pb3_latest.Empty{}},

		// Empty repeated fields are ignored.
		{&pb2_latest.Repetitive{StringField: []string{}}},
		{&pb3_latest.Repetitive{StringField: []string{}}},

		// Empty map fields are ignored.
		{&pb2_latest.StringMaps{StringToString: map[string]string{}}},
		{&pb3_latest.StringMaps{StringToString: map[string]string{}}},

		// Proto3 scalar fields set to their default values are considered empty.
		{&pb3_latest.Simple{BoolField: false}},
		{&pb3_latest.Simple{BytesField: []byte{}}},
		{&pb3_latest.Simple{DoubleField: 0}},
		{&pb3_latest.Simple{DoubleField: 0.0}},
		{&pb3_latest.Simple{Fixed32Field: 0}},
		{&pb3_latest.Simple{Fixed64Field: 0}},
		{&pb3_latest.Simple{FloatField: 0}},
		{&pb3_latest.Simple{FloatField: 0.0}},
		{&pb3_latest.Simple{Int32Field: 0}},
		{&pb3_latest.Simple{Int64Field: 0}},
		{&pb3_latest.Simple{Sfixed32Field: 0}},
		{&pb3_latest.Simple{Sfixed64Field: 0}},
		{&pb3_latest.Simple{Sint32Field: 0}},
		{&pb3_latest.Simple{Sint64Field: 0}},
		{&pb3_latest.Simple{StringField: ""}},
		{&pb3_latest.Simple{Uint32Field: 0}},
		{&pb3_latest.Simple{Uint64Field: 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.msg), func(t *testing.T) {
			h := hasher{}

			got := getHash(t, func() ([]byte, error) {
				return h.hashMessage(tc.msg.ProtoReflect())
			})

			if diff := cmp.Diff(emptyHash, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHashIntegerFields(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"equivalence": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Fixed32Message{Values: []uint32{0, 1, 2}},
				&pb2_latest.Fixed64Message{Values: []uint64{0, 1, 2}},
				&pb2_latest.Int32Message{Values: []int32{0, 1, 2}},
				&pb2_latest.Int64Message{Values: []int64{0, 1, 2}},
				&pb2_latest.Sfixed32Message{Values: []int32{0, 1, 2}},
				&pb2_latest.Sfixed64Message{Values: []int64{0, 1, 2}},
				&pb2_latest.Sint32Message{Values: []int32{0, 1, 2}},
				&pb2_latest.Sint64Message{Values: []int64{0, 1, 2}},
				&pb2_latest.Uint32Message{Values: []uint32{0, 1, 2}},
				&pb2_latest.Uint64Message{Values: []uint64{0, 1, 2}},

				&pb3_latest.Fixed32Message{Values: []uint32{0, 1, 2}},
				&pb3_latest.Fixed64Message{Values: []uint64{0, 1, 2}},
				&pb3_latest.Int32Message{Values: []int32{0, 1, 2}},
				&pb3_latest.Int64Message{Values: []int64{0, 1, 2}},
				&pb3_latest.Sfixed32Message{Values: []int32{0, 1, 2}},
				&pb3_latest.Sfixed64Message{Values: []int64{0, 1, 2}},
				&pb3_latest.Sint32Message{Values: []int32{0, 1, 2}},
				&pb3_latest.Sint64Message{Values: []int64{0, 1, 2}},
				&pb3_latest.Uint32Message{Values: []uint32{0, 1, 2}},
				&pb3_latest.Uint64Message{Values: []uint64{0, 1, 2}},
			},
			obj: map[string][]int32{"values": {0, 1, 2}},
			// No equivalent JSON: JSON does not have an "integer" type. All numbers are floats.
			want: "42794fb0e73c2b5f427aa76486555d07589359054848396ddf173e9e0b4ab931",
		},
		"equivalence (with negatives)": {
			options: []Option{FieldNamesAsKeys()},
			protos:  []proto.Message{},
			obj:     map[string][]int32{"values": {-2, -1, 0, 1, 2}},
			// No equivalent JSON: JSON does not have an "integer" type. All numbers are floats.
			want: "6cb613a53b6086b88dbda40b30e902adb41288b0b1f7a627905beaa764ee49cb",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashFloatFields(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"float fields (hashing key field numbers)": {
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Values: []float64{-2, -1, 0, 1, 2}},
				&pb3_latest.DoubleMessage{Values: []float64{-2, -1, 0, 1, 2}},
				&pb2_latest.FloatMessage{Values: []float32{-2, -1, 0, 1, 2}},
				&pb3_latest.FloatMessage{Values: []float32{-2, -1, 0, 1, 2}},
			},
			// obj: map[string][]float64{"values": {-2, -1, 0, 1, 2}},
			// json: `{2: [-2, -1, 0, 1, 2]}`, skipping json as this is invalid json
			want: "08775d05cd028265e4956a95aef6c050a45652e9c59462da636a8460c5ed52f3",
		},
		"float fields (hashing key field strings)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Values: []float64{-2, -1, 0, 1, 2}},
				&pb3_latest.DoubleMessage{Values: []float64{-2, -1, 0, 1, 2}},
				&pb2_latest.FloatMessage{Values: []float32{-2, -1, 0, 1, 2}},
				&pb3_latest.FloatMessage{Values: []float32{-2, -1, 0, 1, 2}},
			},
			obj:  map[string][]float64{"values": {-2, -1, 0, 1, 2}},
			json: `{"values": [-2, -1, 0, 1, 2]}`,
			want: "586202dddb0e98bb8ce0b7289e29a9f7397b9b1996f3f8fe788f4cfb230b7ee8",
		},
		"float fields (fractions 32)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Values: []float64{0.0078125, 7.888609052210118e-31}},
				&pb3_latest.DoubleMessage{Values: []float64{0.0078125, 7.888609052210118e-31}},
				&pb2_latest.FloatMessage{Values: []float32{0.0078125, 7.888609052210118e-31}},
				&pb3_latest.FloatMessage{Values: []float32{0.0078125, 7.888609052210118e-31}},
			},
			obj:  map[string][]float64{"values": {0.0078125, 7.888609052210118e-31}},
			json: `{"values": [0.0078125, 7.888609052210118e-31]}`,
			want: "7b7cba0ed312bc6611f0523e7c46ce9a2ed9ecb798eb80e1cdf93c95faf503c7",
		},
		"float fields (fractions 64)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Values: []float64{-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},
				&pb3_latest.DoubleMessage{Values: []float64{-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},
				&pb2_latest.FloatMessage{Values: []float32{-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},
				&pb3_latest.FloatMessage{Values: []float32{-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},
			},
			json: `{"values": [-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625]}`,
			want: "ac261ff3d8b933998e3fea278539eb40b15811dd835d224e0150dce4794168b7",
		},
		"float fields (Non-equivalence of Floats using different representations)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.FloatMessage{Value: proto.Float32(0.1)},
				&pb3_latest.FloatMessage{Value: 0.1},
				// A float64 "0.1" is not equal to a float32 "0.1".
				// However, float32 "0.1" is equal to float64 "1.0000000149011612e-1".
				&pb2_latest.DoubleMessage{Value: proto.Float64(1.0000000149011612e-1)},
				&pb3_latest.DoubleMessage{Value: 1.0000000149011612e-1},
			},
			obj:  map[string]float32{"value": 0.1},
			json: `{"value": 1.0000000149011612e-1}`,
			want: "7081ed6a1e7ad8e7f981a2894a3bd6d3b0b0033b69c03cce84b61dd063f4efaa",
		},
		"float fields (There's no float32 number that is equivalent to a float64 '0.1'.)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(0.1)},
				&pb3_latest.DoubleMessage{Value: 0.1},
			},
			obj:  map[string]float64{"value": 0.1},
			json: `{"value": 0.1}`,
			want: "e175fbe785bae88b598d3ecaad8a64d2a998e9f673173a226868f2ef312a5225",
		},
		"float fields (Non-equivalence of Floats using different representations - decimal)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.FloatMessage{Value: proto.Float32(1.2163543e+25)},
				&pb3_latest.FloatMessage{Value: 1.2163543e+25},
				// The decimal representation of the equivalent 64-bit float is different.
				&pb2_latest.DoubleMessage{Value: proto.Float64(1.2163543234531120e+25)},
				&pb3_latest.DoubleMessage{Value: 1.2163543234531120e+25},
			},
			obj:  map[string]float32{"value": 1.2163543e+25},
			json: `{"value": 1.2163543234531120e+25}`,
			want: "bbb17cf7312f2ba5b0002d781f16d1ab50c3d25dc044ed3428750826a1c68653",
		},
		"float fields (no float32 number that is equivalent to a float64 '1e+25')": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(1e+25)},
				&pb3_latest.DoubleMessage{Value: 1e+25},
			},
			obj:  map[string]float64{"value": 1e+25},
			json: `{"value": 1e+25}`,
			want: "874beabbede24974a9f3f74e3448670e0c42c0aaba082f18b963b72253649362",
		},
		"float fields (proto2 unset)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(0)},
				&pb2_latest.FloatMessage{Value: proto.Float32(0)},
			},
			obj:  map[string]float64{"value": 0},
			json: `{"value":0}`,
			want: "94136b0850db069dfd7bee090fc7ede48aa7da53ae3cc8514140a493818c3b91",
		},
		"float fields (special NaN)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(math.NaN())},
				&pb3_latest.DoubleMessage{Value: math.NaN()},

				&pb2_latest.FloatMessage{Value: proto.Float32(float32(math.NaN()))},
				&pb3_latest.FloatMessage{Value: float32(math.NaN())},
			},
			obj: map[string]float64{"value": math.NaN()},
			// No equivalent JSON: JSON does not support special float values.
			// See: https://tools.ietf.org/html/rfc4627#section-2.4
			// json: `{"value": NaN}`,
			want: "16614de29b0823c41cabc993fa6c45da87e4e74c5d836edbcddcfaaf06ffafd1",
		},
		"float fields (special Inf(+))": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(math.Inf(1))},
				&pb3_latest.DoubleMessage{Value: math.Inf(1)},

				&pb2_latest.FloatMessage{Value: proto.Float32(float32(math.Inf(1)))},
				&pb3_latest.FloatMessage{Value: float32(math.Inf(1))},
			},
			obj: map[string]float64{"value": math.Inf(1)},
			// No equivalent JSON: JSON does not support special float values.
			// See: https://tools.ietf.org/html/rfc4627#section-2.4
			// json: `{"value": Inf}`,
			want: "c58cd512e86204e99cb6c11d83bb3daaccdd946e66383004cb9b7f87f762935c",
		},
		"float fields (special Inf(-))": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(math.Inf(-1))},
				&pb3_latest.DoubleMessage{Value: math.Inf(-1)},

				&pb2_latest.FloatMessage{Value: proto.Float32(float32(math.Inf(-1)))},
				&pb3_latest.FloatMessage{Value: float32(math.Inf(-1))},
			},
			obj: map[string]float64{"value": math.Inf(-1)},
			// No equivalent JSON: JSON does not support special float values.
			// See: https://tools.ietf.org/html/rfc4627#section-2.4
			// json: `{"value": Inf}`,
			want: "1a4ffd7e9dc1f915c5b3b821d9194ac7d6d2bdec947aa8c3b3b1e9017c651331",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashStringFields(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"unicode": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{StringField: proto.String("你好")},
				&pb3_latest.Simple{StringField: "你好"},
			},
			obj:  map[string]string{"string_field": "你好"},
			json: "{\"string_field\":\"你好\"}",
			want: "de0086ad683b5f8affffbbcbe57d09e5377aa47cb32f6f0b1bdecd2e54b9137d",
		},
		"esc": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{StringField: proto.String("\u03d3")},
				&pb3_latest.Simple{StringField: "\u03d3"},
			},
			obj:  map[string]string{"string_field": "\u03d3"},
			json: "{\"string_field\":\"\u03d3\"}",
			want: "12441188aebffcc3a1e625d825391678d8417c77e645fc992d1ab5b549c659a7",
		},
		"normalization": {
			// Note that this is the same character as above, but hashes differently
			// unless unicode normalisation is applied.
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{StringField: proto.String("\u03d2\u0301")},
				&pb3_latest.Simple{StringField: "\u03d2\u0301"},
			},
			obj:  map[string]string{"string_field": "\u03d2\u0301"},
			json: "{\"string_field\":\"\u03d2\u0301\"}",
			want: "1f33a91552e7a527fdf2de0d25f815590f1a3e2dc8340507d20d4ee42462d0a2",
		},
		"repeated empty": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{StringField: []string{""}},
				&pb3_latest.Repetitive{StringField: []string{""}},
			},
			obj:  map[string][]string{"string_field": {""}},
			json: "{\"string_field\":[\"\"]}",
			want: "63e64f0ed286e0d8f30735e6646ea9ef48174c23ba09a05288b4233c6e6a9419",
		},
		"repeated unicode": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{StringField: []string{"", "Test", "你好", "\u03d3"}},
				&pb3_latest.Repetitive{StringField: []string{"", "Test", "你好", "\u03d3"}},
			},
			obj:  map[string][]string{"string_field": {"", "Test", "你好", "\u03d3"}},
			json: "{\"string_field\":[\"\",\"Test\",\"你好\",\"\u03d3\"]}",
			want: "f76ae15a2685a5ec0e45f9ad7d75e492e6a17d31811480fbaf00af451fb4e98e",
		},
	} {
		tc.Check(name, t)
	}
}

// TestHashProto2DefaultFields checks that proto2 default field values are properly hashed.
func TestHashProto2DefaultFields(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"simple bool": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{BoolField: proto.Bool(false)},
			},
			obj:  map[string]bool{"bool_field": false},
			json: "{\"bool_field\":false}",
			want: "1ab5ecdbe4176473024f7efd080593b740d22d076d06ea6edd8762992b484a12",
		},
		"simple bytes": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{BytesField: []byte{}},
			},
			obj: map[string][]byte{"bytes_field": {}},
			// json: skipped - JSON does not have 'bytes' representation
			want: "10a0dbbfa097b731c7a505246ffa96a82f997b8c25892d76d3b8b1355e529e05",
		},
		"simple string": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{StringField: proto.String("")},
			},
			obj:  map[string]string{"string_field": ""},
			json: "{\"string_field\":\"\"}",
			want: "2d60c2941830ef4bb14424e47c6cd010f2b95e5e34291f429998288a60ac8c22",
		},
		"ints": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Fixed32Message{Value: proto.Uint32(0)},
				&pb2_latest.Fixed64Message{Value: proto.Uint64(0)},
				&pb2_latest.Int32Message{Value: proto.Int32(0)},
				&pb2_latest.Int64Message{Value: proto.Int64(0)},
				&pb2_latest.Sfixed32Message{Value: proto.Int32(0)},
				&pb2_latest.Sfixed64Message{Value: proto.Int64(0)},
				&pb2_latest.Sint32Message{Value: proto.Int32(0)},
				&pb2_latest.Sint64Message{Value: proto.Int64(0)},
				&pb2_latest.Uint32Message{Value: proto.Uint32(0)},
				&pb2_latest.Uint64Message{Value: proto.Uint64(0)},
			},
			obj:  map[string]int64{"value": 0},
			want: "49f031b73dad26859ffeea8a2bb170aaf7358d2277b00c7fc7ea8edcd37e53a1",
		},
		"floats": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(0.0)},
				&pb2_latest.FloatMessage{Value: proto.Float32(0.0)},
			},
			obj:  map[string]float64{"value": 0.0},
			json: "{\"value\":0.0}",
			want: "94136b0850db069dfd7bee090fc7ede48aa7da53ae3cc8514140a493818c3b91",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashOneofFields(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"empty": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Singleton{},
				&pb3_latest.Singleton{},

				&pb2_latest.Empty{},
				&pb3_latest.Empty{},
			},
			obj:  map[int64]string{},
			json: `{}`,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"one selected but empty": {
			protos: []proto.Message{
				// Only proto2 has empty values.
				&pb2_latest.Simple{BoolField: proto.Bool(false)},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheBool{}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheBool{}},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheBool{TheBool: false}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheBool{TheBool: false}},
			},
			obj:  map[int64]bool{1: false},
			want: "8a956cfa8e9b45b738cb8dc8a3dc7126dab3cbd2c07c80fa1ec312a1a31ed709",
		},
		"One of the options selected with content (empty string)": {
			protos: []proto.Message{
				// Only proto2 has empty values.
				&pb2_latest.Simple{StringField: proto.String("")},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheString{}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheString{}},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheString{TheString: ""}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheString{TheString: ""}},
			},
			obj:  map[int64]string{25: ""},
			want: "79cff9d2d0ee6c6071c82b58d1a2fcf056b58c4501606862489e5731644c755a",
		},
		"One of the options selected with content (ints)": {
			protos: []proto.Message{
				// Only proto2 has empty values.
				&pb2_latest.Simple{Int32Field: proto.Int32(0)},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheInt32{}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheInt32{}},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheInt32{TheInt32: 0}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheInt32{TheInt32: 0}},
			},
			obj:  map[int64]int32{13: 0},
			want: "bafd42680c987c47a76f72e08ed975877162efdb550d2c564c758dc7d988468f",
		},
		"One of the options selected with content (strings)": {
			protos: []proto.Message{
				&pb2_latest.Simple{StringField: proto.String("TEST!")},
				&pb3_latest.Simple{StringField: "TEST!"},
				//
				// For protobufs, it is legal (and backwards-compatible) to update a message by wrapping
				// an existing field within a oneof rule. Therefore, both objects (using old schem and
				// the new schema) should result in the same objecthash.
				//
				// Example:
				//
				// # Old schema:               | # New schema:
				// message Simple {            | message Singleton {
				//   string string_field = 25; |   oneof singleton {
				// }                           |     string the_string = 25;
				//                             |   }
				//                             | }
				//
				// The following examples demonstrate this equivalence.
				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheString{TheString: "TEST!"}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheString{TheString: "TEST!"}},
			},
			obj:  map[int64]string{25: "TEST!"},
			want: "336cdbca99fd46157bc47bcc456f0ac7f1ef3be7a79acf3535f671434b53944f",
		},
		"One of the options selected with content (equiv case ints)": {
			protos: []proto.Message{
				&pb2_latest.Simple{Int32Field: proto.Int32(99)},
				&pb3_latest.Simple{Int32Field: 99},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheInt32{TheInt32: 99}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheInt32{TheInt32: 99}},
			},
			obj:  map[int64]int32{13: 99},
			want: "65517521bc278528d25caf1643da0f094fd88dad50205c9743e3c984a7c53b7d",
		},
		"One of the options selected with content (nested)": {
			protos: []proto.Message{
				&pb2_latest.Simple{SingletonField: &pb2_latest.Singleton{}},
				&pb3_latest.Simple{SingletonField: &pb3_latest.Singleton{}},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheSingleton{TheSingleton: &pb2_latest.Singleton{}}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheSingleton{TheSingleton: &pb3_latest.Singleton{}}},
			},
			obj:  map[int64]map[int64]int64{35: {}},
			want: "4967c72525c764229f9fbf1294764c9aedc0d4f9f4c52e04a19c7f35ca65f517",
		},
		"One of the options selected with content (double nested)": {
			protos: []proto.Message{
				&pb2_latest.Simple{SingletonField: &pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheSingleton{TheSingleton: &pb2_latest.Singleton{}}}},
				&pb3_latest.Simple{SingletonField: &pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheSingleton{TheSingleton: &pb3_latest.Singleton{}}}},

				&pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheSingleton{TheSingleton: &pb2_latest.Singleton{Singleton: &pb2_latest.Singleton_TheSingleton{TheSingleton: &pb2_latest.Singleton{}}}}},
				&pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheSingleton{TheSingleton: &pb3_latest.Singleton{Singleton: &pb3_latest.Singleton_TheSingleton{TheSingleton: &pb3_latest.Singleton{}}}}},
			},
			obj:  map[int64]map[int64]map[int64]int64{35: {35: {}}},
			want: "8ea95bbda0f42073a61f46f9f375f48d5a7cb034fce56b44f958470fda5236d0",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashMapFields(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"boolean maps": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.BoolMaps{BoolToString: map[bool]string{true: "NOT FALSE", false: "NOT TRUE"}},
				&pb3_latest.BoolMaps{BoolToString: map[bool]string{true: "NOT FALSE", false: "NOT TRUE"}},
			},
			obj: map[string]map[bool]string{"bool_to_string": {true: "NOT FALSE", false: "NOT TRUE"}},
			// No equivalent JSON: JSON does not have an "integer" type. All numbers are floats.
			want: "d89d053bf7b37b4784832c72445661db99538fe1d490988575409a9040084f18",
		},
		"integer maps": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.IntMaps{IntToString: map[int64]string{0: "ZERO"}},
				&pb3_latest.IntMaps{IntToString: map[int64]string{0: "ZERO"}},
			},
			obj: map[string]map[int64]string{"int_to_string": {0: "ZERO"}},
			// No equivalent JSON object because JSON map keys must be strings.
			want: "53892192fb69cbd93ceb0552ca571b8505887f25d6f12822025341f16983a6af",
		},
		"string maps": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.StringMaps{StringToString: map[string]string{"foo": "bar"}},
				&pb3_latest.StringMaps{StringToString: map[string]string{"foo": "bar"}},
			},
			obj:  map[string]map[string]string{"string_to_string": {"foo": "bar"}},
			json: `{"string_to_string": {"foo": "bar"}}`,
			want: "cadfe560995647c63c20234a6409d2b1b8cf8dcf7d8e420ca33f23ff9ca9abfa",
		},
		"string maps (unicode)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.StringMaps{StringToString: map[string]string{
					"": "你好", "你好": "\u03d3", "\u03d3": "\u03d2\u0301"}},
				&pb3_latest.StringMaps{StringToString: map[string]string{
					"": "你好", "你好": "\u03d3", "\u03d3": "\u03d2\u0301"}},
			},
			obj:  map[string]map[string]string{"string_to_string": {"": "你好", "你好": "\u03d3", "\u03d3": "\u03d2\u0301"}},
			json: `{"string_to_string": {"": "你好", "你好": "\u03d3", "\u03d3": "\u03d2\u0301"}}`,
			want: "be8b5ae6d5986cde37ab8b395c66045fbb69a8b3b534fa34df7c19a640f4cd66",
		},
		"message maps": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.StringMaps{StringToSimple: map[string]*pb2_latest.Simple{"foo": {}}},
				&pb3_latest.StringMaps{StringToSimple: map[string]*pb3_latest.Simple{"foo": {}}},
			},
			obj:  map[string]map[string]map[string]string{"string_to_simple": {"foo": {}}},
			json: `{"string_to_simple": {"foo": {}}}`,
			want: "58057927bb1a123452a2d75071b55b08e426490ca42c3dd14e3be59183ac4751",
		},
	} {
		tc.Check(name, t)
	}
}

// TestHashOtherTypes performs tests on types that do not have their own test file.
func TestHashOtherTypes(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"nil": {
			protos: []proto.Message{
				nil,
			},
			json: "null",
			obj:  nil,
			want: "1b16b1df538ba12dc3f97edbb85caa7050d46c148134290feba80f8236c83db9",
		},
		"booleans (true)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{BoolField: proto.Bool(true)},
				&pb3_latest.Simple{BoolField: true},
			},
			json: `{"bool_field": true}`,
			obj:  map[string]bool{"bool_field": true},
			want: "7b2ac6048e6c8797205505ea486539a5589583be43154da88785a5121e2d6899",
		},
		"booleans (false)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{BoolField: proto.Bool(false)},
				// proto3 scalar fields set to their default value are considered empty.
			},
			json: `{"bool_field": false}`,
			obj:  map[string]bool{"bool_field": false},
			want: "1ab5ecdbe4176473024f7efd080593b740d22d076d06ea6edd8762992b484a12",
		},
		"bytes": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Simple{BytesField: []byte{0, 0, 0}},
				&pb3_latest.Simple{BytesField: []byte{0, 0, 0}},
			},
			// No equivalent JSON: JSON does not have a "bytes" type.
			obj:  map[string][]byte{"bytes_field": []byte("\000\000\000")},
			want: "fdd59e1f3120117943124cb9c39da79ac47ea631343ff9154dffb0e64550789c",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashTimestamp(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"Empty/Zero Timestamps": {
			// The semantics of the Timestamp object imply that the distinction between
			// unset and zero happen at the message level, rather than the field level.
			//
			// As a result, an unset timestamp is one where the proto itself is nil,
			// while an explicitly set timestamp with unset fields is considered to be
			// explicitly set to 0.
			//
			// This is unlike normal proto3 messages, where unset/zero fields must be
			// considered to be unset, because they're indistinguishable in the general
			// case.
			protos: []proto.Message{
				&timestamppb.Timestamp{},
				&timestamppb.Timestamp{Seconds: 0, Nanos: 0},
			},
			// JSON treats all numbers as floats, so it is not possible to have an equivalent JSON string.
			obj:  []int64{0, 0},
			want: "3a82b649344529f03f52c1833f5aecc488a53b31461a1f54c305d149b12b8f53",
		},
		"Normal Timestamps": {
			protos: []proto.Message{
				&timestamppb.Timestamp{Seconds: 1525450021, Nanos: 123456789},
			},
			// JSON treats all numbers as floats, so it is not possible to have an equivalent JSON string.
			obj:  []int64{1525450021, 123456789},
			want: "1fd36770664df599ad44e4e4f06b1fad6ef7a4b3f316d79ca11bea668032a199",
		},
		"Timestamps within other protos (zero)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{TimestampField: &timestamppb.Timestamp{}},
				&pb2_latest.KnownTypes{TimestampField: &timestamppb.Timestamp{Seconds: 0, Nanos: 0}},

				&pb3_latest.KnownTypes{TimestampField: &timestamppb.Timestamp{}},
				&pb3_latest.KnownTypes{TimestampField: &timestamppb.Timestamp{Seconds: 0, Nanos: 0}},
			},
			// JSON treats all numbers as floats, so it is not possible to have an equivalent JSON string.
			obj:  map[string][]int64{"timestamp_field": {0, 0}},
			want: "8457fe431752dbc5c47301c2546fcf6f0ad8c5317092b443e187d18e312e497e",
		},
		"Timestamps within other protos (non-zero)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{TimestampField: &timestamppb.Timestamp{Seconds: 1525450021, Nanos: 123456789}},
				&pb3_latest.KnownTypes{TimestampField: &timestamppb.Timestamp{Seconds: 1525450021, Nanos: 123456789}},
			},
			// JSON treats all numbers as floats, so it is not possible to have an equivalent JSON string.
			obj:  map[string][]int64{"timestamp_field": {1525450021, 123456789}},
			want: "cf99942e3f8d1212f4ce263e206d64e29525b97b91368e71f9595bce83ac6a3e",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashDuration(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"Empty/Zero Duration": {
			// The semantics of the Duration object imply that the distinction between
			// unset and zero happen at the message level, rather than the field level.
			//
			// As a result, an unset timestamp is one where the proto itself is nil,
			// while an explicitly set timestamp with unset fields is considered to be
			// explicitly set to 0.
			//
			// This is unlike normal proto3 messages, where unset/zero fields must be
			// considered to be unset, because they're indistinguishable in the general
			// case.
			protos: []proto.Message{
				&durationpb.Duration{},
				&durationpb.Duration{Seconds: 0, Nanos: 0},
			},
			// JSON treats all numbers as floats, so it is not possible to have an equivalent JSON string.
			obj:  []int64{0, 0},
			want: "3a82b649344529f03f52c1833f5aecc488a53b31461a1f54c305d149b12b8f53",
		},
		"Normal Duration": {
			protos: []proto.Message{
				&durationpb.Duration{Seconds: 1525450021, Nanos: 123456789},
			},
			// JSON treats all numbers as floats, so it is not possible to have an equivalent JSON string.
			obj:  []int64{1525450021, 123456789},
			want: "1fd36770664df599ad44e4e4f06b1fad6ef7a4b3f316d79ca11bea668032a199",
		},
		"Durations within other protos (zero)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{DurationField: &durationpb.Duration{}},
				&pb2_latest.KnownTypes{DurationField: &durationpb.Duration{Seconds: 0, Nanos: 0}},

				&pb3_latest.KnownTypes{DurationField: &durationpb.Duration{}},
				&pb3_latest.KnownTypes{DurationField: &durationpb.Duration{Seconds: 0, Nanos: 0}},
			},
			// JSON treats all numbers as floats, so it is not possible to have an equivalent JSON string.
			obj:  map[string][]int64{"duration_field": {0, 0}},
			want: "80668dc83d8e5c0c9e24afba293e69cb1ce697772521f7a8ea3afc20a6dd617a",
		},
		"Durations within other protos (non-zero)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{DurationField: &durationpb.Duration{Seconds: 1525450021, Nanos: 123456789}},
				&pb3_latest.KnownTypes{DurationField: &durationpb.Duration{Seconds: 1525450021, Nanos: 123456789}},
			},
			// JSON treats all numbers as floats, so it is not possible to have an equivalent JSON string.
			obj:  map[string][]int64{"duration_field": {1525450021, 123456789}},
			want: "df5f2f19e08ba1027e491117ac0b5269d66439b0326a05de356f2840ba8354c1",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashBoolValue(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"false": {
			protos: []proto.Message{
				&wrapperspb.BoolValue{},
				&wrapperspb.BoolValue{Value: false},
			},
			json: `false`,
			obj:  false,
			want: "c02c0b965e023abee808f2b548d8d5193a8b5229be6f3121a6f16e2d41a449b3",
		},
		"true": {
			protos: []proto.Message{
				&wrapperspb.BoolValue{Value: true},
			},
			json: `true`,
			obj:  true,
			want: "7dc96f776c8423e57a2785489a3f9c43fb6e756876d6ad9a9cac4aa4e72ec193",
		},
		"BoolValue within other protos (not set)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{BoolValueField: nil},
				&pb3_latest.KnownTypes{BoolValueField: nil},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"BoolValue within other protos (set, default false)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{BoolValueField: &wrapperspb.BoolValue{}},
				&pb2_latest.KnownTypes{BoolValueField: &wrapperspb.BoolValue{Value: false}},
				&pb3_latest.KnownTypes{BoolValueField: &wrapperspb.BoolValue{}},
				&pb3_latest.KnownTypes{BoolValueField: &wrapperspb.BoolValue{Value: false}},
			},
			json: `{"bool_value_field": false}`,
			obj:  map[string]bool{"bool_value_field": false},
			want: "8ec24416eca90851428f5b63b7529d2ea7d24fe0e9b3ca11ea2ee851d0ce2280",
		},
		"BoolValue within other protos (set, true)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{BoolValueField: &wrapperspb.BoolValue{Value: true}},
				&pb3_latest.KnownTypes{BoolValueField: &wrapperspb.BoolValue{Value: true}},
			},
			json: `{"bool_value_field": true}`,
			obj:  map[string]bool{"bool_value_field": true},
			want: "3363c4b1d91d9469bbcca6c255245fba3fdc340722bd0b69c3d1a3dc84ce0d58",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashFloatValue(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"0.0": {
			protos: []proto.Message{
				&wrapperspb.FloatValue{},
				&wrapperspb.FloatValue{Value: 0.0},
			},
			json: `0.0`,
			obj:  0.0,
			want: "60101d8c9cb988411468e38909571f357daa67bff5a7b0a3f9ae295cd4aba33d",
		},
		"-1.0": {
			protos: []proto.Message{
				&wrapperspb.FloatValue{Value: -1.0},
			},
			json: `-1.0`,
			obj:  -1.0,
			want: "f706daa44d7e40e21ea202c36119057924bb28a49949d8ddaa9c8c3c9367e602",
		},
		"+1.0": {
			protos: []proto.Message{
				&wrapperspb.FloatValue{Value: 1.0},
			},
			json: `1.0`,
			obj:  1.0,
			want: "f01adc732390ab024d64080e0b173f0ee3a1610efbdd4ce2a13bbf8d9b26c639",
		},
		"FloatValue within other protos (not set)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{FloatValueField: nil},
				&pb3_latest.KnownTypes{FloatValueField: nil},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"FloatValue within other protos (set, default 0.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{FloatValueField: &wrapperspb.FloatValue{}},
				&pb2_latest.KnownTypes{FloatValueField: &wrapperspb.FloatValue{Value: 0.0}},
				&pb3_latest.KnownTypes{FloatValueField: &wrapperspb.FloatValue{}},
				&pb3_latest.KnownTypes{FloatValueField: &wrapperspb.FloatValue{Value: 0.0}},
			},
			json: `{"float_value_field": 0.0}`,
			obj:  map[string]float32{"float_value_field": 0.0},
			want: "75085520c0294c8467895b2bd9939cf4a6373629f95f155eb3c755c7debb326d",
		},
		"FloatValue within other protos (set, 1.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{FloatValueField: &wrapperspb.FloatValue{Value: 1.0}},
				&pb3_latest.KnownTypes{FloatValueField: &wrapperspb.FloatValue{Value: 1.0}},
			},
			json: `{"float_value_field": 1.0}`,
			obj:  map[string]float32{"float_value_field": 1.0},
			want: "a503b683c8668908da9abdc3ac2683fc5ee8c1aac3c6ccab020005b30e76ac9f",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashMessageFullnameIdentifier(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"empty (without option)": {
			protos: []proto.Message{
				&pb3_latest.Empty{},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"empty (with option)": {
			options: []Option{MessageFullnameIdentifier()},
			protos: []proto.Message{
				&pb3_latest.Empty{},
			},
			// json: `{}`, // skip json check: not possible to express typed structs
			obj:  nil,
			want: "f871a6b9331f109ef970f8130baffc29bb5753a2970f66af6004475171efdd9b",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashStringValue(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"empty": {
			protos: []proto.Message{
				&wrapperspb.StringValue{},
				&wrapperspb.StringValue{Value: ""},
			},
			json: `""`,
			obj:  "",
			want: "0bfe935e70c321c7ca3afc75ce0d0ca2f98b5422e008bb31c00c6d7f1f1c0ad6",
		},
		"你好": {
			protos: []proto.Message{
				&wrapperspb.StringValue{Value: "你好"},
			},
			json: `"你好"`,
			obj:  "你好",
			want: "462b68f5e3d75aed5f02841b4ffee068d4cf33ce1b155105b71a9e5f358026df",
		},
		"StringValue within other protos (not set)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{StringValueField: nil},
				&pb3_latest.KnownTypes{StringValueField: nil},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"StringValue within other protos (set, default empty)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{StringValueField: &wrapperspb.StringValue{}},
				&pb2_latest.KnownTypes{StringValueField: &wrapperspb.StringValue{Value: ""}},
				&pb3_latest.KnownTypes{StringValueField: &wrapperspb.StringValue{}},
				&pb3_latest.KnownTypes{StringValueField: &wrapperspb.StringValue{Value: ""}},
			},
			json: `{"string_value_field": ""}`,
			obj:  map[string]string{"string_value_field": ""},
			want: "2ce75d087e557a68b232652d48e6aac5f3fc457c597a0ed07a1b63a4c2d16039",
		},
		"StringValue within other protos (set, nonempty)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{StringValueField: &wrapperspb.StringValue{Value: "bob"}},
				&pb3_latest.KnownTypes{StringValueField: &wrapperspb.StringValue{Value: "bob"}},
			},
			json: `{"string_value_field": "bob"}`,
			obj:  map[string]string{"string_value_field": "bob"},
			want: "2ca37671e89e4a192caac37288ec665498eb7b2917b7a7354a23225a35765588",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashInt32Value(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"0": {
			protos: []proto.Message{
				&wrapperspb.Int32Value{},
				&wrapperspb.Int32Value{Value: 0},
			},
			// JSON skipped - represents numbers as floats
			obj:  0,
			want: "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		"-1": {
			protos: []proto.Message{
				&wrapperspb.Int32Value{Value: -1},
			},
			// JSON skipped - represents numbers as floats
			obj:  -1,
			want: "f105b11df43d5d321f5c773ef904af979024887b4d2b0fab699387f59e2ff01e",
		},
		"+1": {
			protos: []proto.Message{
				&wrapperspb.Int32Value{Value: 1},
			},
			// JSON skipped - represents numbers as floats
			obj:  1,
			want: "4cd9b7672d7fbee8fb51fb1e049f690342035f543a8efe734b7b5ffb0c154a45",
		},
		"Int32Value within other protos (not set)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Int32ValueField: nil},
				&pb3_latest.KnownTypes{Int32ValueField: nil},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"Int32Value within other protos (set, default 0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Int32ValueField: &wrapperspb.Int32Value{}},
				&pb2_latest.KnownTypes{Int32ValueField: &wrapperspb.Int32Value{Value: 0}},
				&pb3_latest.KnownTypes{Int32ValueField: &wrapperspb.Int32Value{}},
				&pb3_latest.KnownTypes{Int32ValueField: &wrapperspb.Int32Value{Value: 0}},
			},
			// JSON skipped - represents numbers as floats
			obj:  map[string]int32{"int32_value_field": 0},
			want: "f45c9b89d9a758f70fee58bad947bca07bd20a31119d927588e7bb11ef17180d",
		},
		"Int32Value within other protos (set, 1.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Int32ValueField: &wrapperspb.Int32Value{Value: 1}},
				&pb3_latest.KnownTypes{Int32ValueField: &wrapperspb.Int32Value{Value: 1}},
			},
			// JSON skipped - represents numbers as floats
			obj:  map[string]int32{"int32_value_field": 1},
			want: "95701cadfe534df4a6f4765625b01a45207a2a9b6479211361a935e33c565195",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashInt64Value(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"0": {
			protos: []proto.Message{
				&wrapperspb.Int64Value{},
				&wrapperspb.Int64Value{Value: 0},
			},
			// JSON skipped - represents numbers as floats
			obj:  0,
			want: "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		"-1": {
			protos: []proto.Message{
				&wrapperspb.Int64Value{Value: -1},
			},
			// JSON skipped - represents numbers as floats
			obj:  -1,
			want: "f105b11df43d5d321f5c773ef904af979024887b4d2b0fab699387f59e2ff01e",
		},
		"+1": {
			protos: []proto.Message{
				&wrapperspb.Int64Value{Value: 1},
			},
			// JSON skipped - represents numbers as floats
			obj:  1,
			want: "4cd9b7672d7fbee8fb51fb1e049f690342035f543a8efe734b7b5ffb0c154a45",
		},
		"Int64Value within other protos (not set)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Int64ValueField: nil},
				&pb3_latest.KnownTypes{Int64ValueField: nil},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"Int64Value within other protos (set, default 0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Int64ValueField: &wrapperspb.Int64Value{}},
				&pb2_latest.KnownTypes{Int64ValueField: &wrapperspb.Int64Value{Value: 0}},
				&pb3_latest.KnownTypes{Int64ValueField: &wrapperspb.Int64Value{}},
				&pb3_latest.KnownTypes{Int64ValueField: &wrapperspb.Int64Value{Value: 0}},
			},
			// JSON skipped - represents numbers as floats
			obj:  map[string]int64{"int64_value_field": 0},
			want: "8459ba1e83e7c72aeb9dcb564daf945f42fe3c1b8837b4266fac7754657160a1",
		},
		"Int64Value within other protos (set, 1.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Int64ValueField: &wrapperspb.Int64Value{Value: 1}},
				&pb3_latest.KnownTypes{Int64ValueField: &wrapperspb.Int64Value{Value: 1}},
			},
			// JSON skipped - represents numbers as floats
			obj:  map[string]int64{"int64_value_field": 1},
			want: "5b7702e61c5d070be2e42f01588be080635bf1244e232aa12107602eb9b2594d",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashUInt32Value(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"0": {
			protos: []proto.Message{
				&wrapperspb.UInt32Value{},
				&wrapperspb.UInt32Value{Value: 0},
			},
			// JSON skipped - represents numbers as floats
			obj:  uint32(0),
			want: "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		"1": {
			protos: []proto.Message{
				&wrapperspb.UInt32Value{Value: 1},
			},
			// JSON skipped - represents numbers as floats
			obj:  uint32(1),
			want: "4cd9b7672d7fbee8fb51fb1e049f690342035f543a8efe734b7b5ffb0c154a45",
		},
		"UInt32Value within other protos (not set)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint32ValueField: nil},
				&pb3_latest.KnownTypes{Uint32ValueField: nil},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"UInt32Value within other protos (set, default 0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint32ValueField: &wrapperspb.UInt32Value{}},
				&pb2_latest.KnownTypes{Uint32ValueField: &wrapperspb.UInt32Value{Value: 0}},
				&pb3_latest.KnownTypes{Uint32ValueField: &wrapperspb.UInt32Value{}},
				&pb3_latest.KnownTypes{Uint32ValueField: &wrapperspb.UInt32Value{Value: 0}},
			},
			// JSON skipped - represents numbers as floats
			obj:  map[string]int32{"uint32_value_field": 0},
			want: "7e3d86d713dec0db2344ff4eb01e40b4cc2c8393840422cf6a716f220b6f6b69",
		},
		"UInt32Value within other protos (set, 1.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint32ValueField: &wrapperspb.UInt32Value{Value: 1}},
				&pb3_latest.KnownTypes{Uint32ValueField: &wrapperspb.UInt32Value{Value: 1}},
			},
			// JSON skipped - represents numbers as floats
			obj:  map[string]int32{"uint32_value_field": 1},
			want: "8b68ebe1ad0456fba3267c950eade1e3d59b1b028af6686bae8d846d911e1b24",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashUInt64Value(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"0": {
			protos: []proto.Message{
				&wrapperspb.UInt64Value{},
				&wrapperspb.UInt64Value{Value: 0},
			},
			// JSON skipped - represents numbers as floats
			obj:  uint64(0),
			want: "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		"1": {
			protos: []proto.Message{
				&wrapperspb.UInt64Value{Value: 1},
			},
			// JSON skipped - represents numbers as floats
			obj:  uint64(1),
			want: "4cd9b7672d7fbee8fb51fb1e049f690342035f543a8efe734b7b5ffb0c154a45",
		},
		"UInt64Value within other protos (not set)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint64ValueField: nil},
				&pb3_latest.KnownTypes{Uint64ValueField: nil},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"UInt64Value within other protos (set, default 0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint64ValueField: &wrapperspb.UInt64Value{}},
				&pb2_latest.KnownTypes{Uint64ValueField: &wrapperspb.UInt64Value{Value: 0}},
				&pb3_latest.KnownTypes{Uint64ValueField: &wrapperspb.UInt64Value{}},
				&pb3_latest.KnownTypes{Uint64ValueField: &wrapperspb.UInt64Value{Value: 0}},
			},
			// JSON skipped - represents numbers as floats
			obj:  map[string]int64{"uint64_value_field": 0},
			want: "832f86706cc1b4136e174c5f0814e965388b01ecad751f1bd23c7523a684b1cc",
		},
		"UInt64Value within other protos (set, 1.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint64ValueField: &wrapperspb.UInt64Value{Value: 1}},
				&pb3_latest.KnownTypes{Uint64ValueField: &wrapperspb.UInt64Value{Value: 1}},
			},
			// JSON skipped - represents numbers as floats
			obj:  map[string]int64{"uint64_value_field": 1},
			want: "65ffdca739a0382cb77f89fbcfa681bc6214d70789c094bedf827b30b7283c08",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashStructValue(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"null": {
			protos: []proto.Message{
				&structpb.Value{Kind: &structpb.Value_NullValue{}},
			},
			json: "null",
			obj:  nil,
			want: "1b16b1df538ba12dc3f97edbb85caa7050d46c148134290feba80f8236c83db9",
		},
		"number": {
			protos: []proto.Message{
				&structpb.Value{Kind: &structpb.Value_NumberValue{}},
			},
			json: "0.0",
			obj:  0.0,
			want: "60101d8c9cb988411468e38909571f357daa67bff5a7b0a3f9ae295cd4aba33d",
		},
		"string": {
			protos: []proto.Message{
				&structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "bob"}},
			},
			json: `"bob"`,
			obj:  "bob",
			want: "5ef421eb52293e5e3919d3c6f08413b873129dd859f4d0ff8273e13a494b9e9e",
		},
		"bool": {
			protos: []proto.Message{
				&structpb.Value{Kind: &structpb.Value_BoolValue{BoolValue: true}},
			},
			json: `true`,
			obj:  true,
			want: "7dc96f776c8423e57a2785489a3f9c43fb6e756876d6ad9a9cac4aa4e72ec193",
		},
		"struct (empty)": {
			protos: []proto.Message{
				&structpb.Value{Kind: &structpb.Value_StructValue{
					StructValue: &structpb.Struct{},
				}},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"struct (non-empty)": {
			protos: []proto.Message{
				&structpb.Value{Kind: &structpb.Value_StructValue{
					StructValue: &structpb.Struct{
						Fields: map[string]*structpb.Value{
							"testing": {
								Kind: &structpb.Value_NullValue{},
							},
						},
					},
				}},
			},
			json: `{"testing": null}`,
			obj:  map[string]interface{}{"testing": nil},
			want: "c41bb94244de3717a2643ee980e392a69cc811aaae18a35370c306fb5f2123b0",
		},
		"list (empty)": {
			protos: []proto.Message{
				&structpb.Value{Kind: &structpb.Value_ListValue{
					ListValue: &structpb.ListValue{},
				}},
			},
			json: `[]`,
			obj:  nil,
			want: "acac86c0e609ca906f632b0e2dacccb2b77d22b0621f20ebece1a4835b93f6f0",
		},
		"list (non-empty)": {
			protos: []proto.Message{
				&structpb.Value{Kind: &structpb.Value_ListValue{
					ListValue: &structpb.ListValue{
						Values: []*structpb.Value{
							{Kind: &structpb.Value_StringValue{StringValue: "foo"}},
							{Kind: &structpb.Value_StringValue{StringValue: "bar"}},
						},
					},
				}},
			},
			json: `["foo", "bar"]`,
			obj:  []string{"foo", "bar"},
			want: "32ae896c413cfdc79eec68be9139c86ded8b279238467c216cf2bec4d5f1e4a2",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashDoubleValue(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"0.0": {
			protos: []proto.Message{
				&wrapperspb.DoubleValue{},
				&wrapperspb.DoubleValue{Value: 0.0},
			},
			json: `0.0`,
			obj:  0.0,
			want: "60101d8c9cb988411468e38909571f357daa67bff5a7b0a3f9ae295cd4aba33d",
		},
		"-1.0": {
			protos: []proto.Message{
				&wrapperspb.DoubleValue{Value: -1.0},
			},
			json: `-1.0`,
			obj:  -1.0,
			want: "f706daa44d7e40e21ea202c36119057924bb28a49949d8ddaa9c8c3c9367e602",
		},
		"+1.0": {
			protos: []proto.Message{
				&wrapperspb.DoubleValue{Value: 1.0},
			},
			json: `1.0`,
			obj:  1.0,
			want: "f01adc732390ab024d64080e0b173f0ee3a1610efbdd4ce2a13bbf8d9b26c639",
		},
		"DoubleValue within other protos (not set)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{DoubleValueField: nil},
				&pb3_latest.KnownTypes{DoubleValueField: nil},
			},
			json: `{}`,
			obj:  nil,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"DoubleValue within other protos (set, default 0.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{DoubleValueField: &wrapperspb.DoubleValue{}},
				&pb2_latest.KnownTypes{DoubleValueField: &wrapperspb.DoubleValue{Value: 0.0}},
				&pb3_latest.KnownTypes{DoubleValueField: &wrapperspb.DoubleValue{}},
				&pb3_latest.KnownTypes{DoubleValueField: &wrapperspb.DoubleValue{Value: 0.0}},
			},
			json: `{"double_value_field": 0.0}`,
			obj:  map[string]float64{"double_value_field": 0.0},
			want: "d593d09e840e41b2f5169561acf24a6b094f0dfb6850cf2a6dcea612f8990a41",
		},
		"DoubleValue within other protos (set, 1.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.KnownTypes{DoubleValueField: &wrapperspb.DoubleValue{Value: 1.0}},
				&pb3_latest.KnownTypes{DoubleValueField: &wrapperspb.DoubleValue{Value: 1.0}},
			},
			json: `{"double_value_field": 1.0}`,
			obj:  map[string]float64{"double_value_field": 1.0},
			want: "20ed14f30c84aaf5066a1e419eadb0214c7ab84d522f0150a1de709bae006cc2",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashRepeatedFields(t *testing.T) {
	for name, tc := range map[string]hashTestCase{
		"empty lists": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{
					BoolField:       []bool{},
					BytesField:      [][]byte{},
					DoubleField:     []float64{},
					Fixed32Field:    []uint32{},
					Fixed64Field:    []uint64{},
					FloatField:      []float32{},
					Int32Field:      []int32{},
					Int64Field:      []int64{},
					Sfixed32Field:   []int32{},
					Sfixed64Field:   []int64{},
					Sint32Field:     []int32{},
					Sint64Field:     []int64{},
					StringField:     []string{},
					Uint32Field:     []uint32{},
					Uint64Field:     []uint64{},
					SimpleField:     []*pb2_latest.Simple{},
					RepetitiveField: []*pb2_latest.Repetitive{},
					SingletonField:  []*pb2_latest.Singleton{},
				},
				&pb3_latest.Repetitive{
					BoolField:       []bool{},
					BytesField:      [][]byte{},
					DoubleField:     []float64{},
					Fixed32Field:    []uint32{},
					Fixed64Field:    []uint64{},
					FloatField:      []float32{},
					Int32Field:      []int32{},
					Int64Field:      []int64{},
					Sfixed32Field:   []int32{},
					Sfixed64Field:   []int64{},
					Sint32Field:     []int32{},
					Sint64Field:     []int64{},
					StringField:     []string{},
					Uint32Field:     []uint32{},
					Uint64Field:     []uint64{},
					SimpleField:     []*pb3_latest.Simple{},
					RepetitiveField: []*pb3_latest.Repetitive{},
					SingletonField:  []*pb3_latest.Singleton{},
				},
			},
			obj:  map[string]interface{}{},
			json: `{}`,
			want: "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4",
		},
		"Lists with strings (empty)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{StringField: []string{""}},
				&pb3_latest.Repetitive{StringField: []string{""}},
			},
			obj:  map[string][]string{"string_field": {""}},
			json: "{\"string_field\": [\"\"]}",
			want: "63e64f0ed286e0d8f30735e6646ea9ef48174c23ba09a05288b4233c6e6a9419",
		},
		"Lists with strings (non-empty)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{StringField: []string{"foo"}},
				&pb3_latest.Repetitive{StringField: []string{"foo"}},
			},
			obj:  map[string][]string{"string_field": {"foo"}},
			json: "{\"string_field\": [\"foo\"]}",
			want: "54c0b7c6e7c9ff0bb6076a2caeccbc96fad77f49b17b7ec9bc17dfe98a7b343e",
		},
		"Lists with strings (non-empty, multiple)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{StringField: []string{"foo", "bar"}},
				&pb3_latest.Repetitive{StringField: []string{"foo", "bar"}},
			},
			json: "{\"string_field\": [\"foo\", \"bar\"]}",
			obj:  map[string][]string{"string_field": {"foo", "bar"}},
			want: "a971a061d199ddf37a365d617f9cd4530efb15e933e0dbaf6602b2908b792056",
		},
		"lists with ints (0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{Int64Field: []int64{0}},
				&pb3_latest.Repetitive{Int64Field: []int64{0}},
			},
			obj:  map[string][]int64{"int64_field": {0}},
			want: "b7e7afd1c1c7beeec4dcc0ced0ec4af2c850add686a12987e8f0b6fcb603733a",
		},
		"lists with ints (span)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{Int64Field: []int64{-2, -1, 0, 1, 2}},
				&pb3_latest.Repetitive{Int64Field: []int64{-2, -1, 0, 1, 2}},
			},
			obj:  map[string][]int64{"int64_field": {-2, -1, 0, 1, 2}},
			want: "44e78ff73bdf5d0da5141e110b22bab240483ba17c40f83553a0e6bbfa671e22",
		},
		"lists with ints (large)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{Int64Field: []int64{123456789012345, 678901234567890}},
				&pb3_latest.Repetitive{Int64Field: []int64{123456789012345, 678901234567890}},
			},
			obj:  map[string][]int64{"int64_field": {123456789012345, 678901234567890}},
			want: "b0ce1b7dfa71b33a16571fea7f3f27341bf5980b040e9d949a8019f3143ecbc7",
		},
		"lists with floats (0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{FloatField: []float32{0}},
				&pb3_latest.Repetitive{FloatField: []float32{0}},
			},
			json: "{\"float_field\": [0]}",
			obj:  map[string][]float32{"float_field": {0}},
			want: "63b09f87ed057a88b38e2a69b6dde327d9e2624384542853327d6b90c83046f9",
		},
		"lists with floats (0.0)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{FloatField: []float32{0.0}},
				&pb3_latest.Repetitive{FloatField: []float32{0.0}},
			},
			json: "{\"float_field\": [0.0]}",
			obj:  map[string][]float32{"float_field": {0.0}},
			want: "63b09f87ed057a88b38e2a69b6dde327d9e2624384542853327d6b90c83046f9",
		},
		"lists with floats (span)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{FloatField: []float32{-2, -1, 0, 1, 2}},
				&pb3_latest.Repetitive{FloatField: []float32{-2, -1, 0, 1, 2}},
			},
			json: "{\"float_field\": [-2, -1, 0, 1, 2]}",
			obj:  map[string][]float32{"float_field": {-2, -1, 0, 1, 2}},
			want: "68b2552f2f33b5dd38c9be0aeee127170c86d8d2b3ab7daebdc2ea124226593f",
		},
		"lists with floats (span 2)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{FloatField: []float32{1, 2, 3}},
				&pb3_latest.Repetitive{FloatField: []float32{1, 2, 3}},
			},
			json: "{\"float_field\": [1, 2, 3]}",
			obj:  map[string][]float32{"float_field": {1, 2, 3}},
			want: "f26c1502d1f9f7bf672cf669290348f9bfdea0af48261f2822aad01927fe1749",
		},
		"lists with floats (span with decimals)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{DoubleField: []float64{1.2345, -10.1234}},
				&pb3_latest.Repetitive{DoubleField: []float64{1.2345, -10.1234}},
			},
			json: "{\"double_field\": [1.2345, -10.1234]}",
			obj:  map[string][]float64{"double_field": {1.2345, -10.1234}},
			want: "2e60f6cdebfeb5e705666e9b0ff0ec652320ae27d77ad89bd4c7ddc632d0b93c",
		},
		"lists with floats (span with decimals 2)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{DoubleField: []float64{1.0, 1.5, 0.0001, 1000.9999999, 2.0, -23.1234, 2.32542}},
				&pb3_latest.Repetitive{DoubleField: []float64{1.0, 1.5, 0.0001, 1000.9999999, 2.0, -23.1234, 2.32542}},
			},
			json: "{\"double_field\": [1.0, 1.5, 0.0001, 1000.9999999, 2.0, -23.1234, 2.32542]}",
			obj:  map[string][]float64{"double_field": {1.0, 1.5, 0.0001, 1000.9999999, 2.0, -23.1234, 2.32542}},
			want: "09a46866ca2c6d406513cd6e25feb6eda7aef4d25259f5ec16bf72f1f8bbcdac",
		},
		"lists with floats (span with decimals large)": {
			options: []Option{FieldNamesAsKeys()},
			protos: []proto.Message{
				&pb2_latest.Repetitive{DoubleField: []float64{123456789012345, 678901234567890}},
				&pb3_latest.Repetitive{DoubleField: []float64{123456789012345, 678901234567890}},
			},
			json: "{\"double_field\": [123456789012345, 678901234567890]}",
			obj:  map[string][]float64{"double_field": {123456789012345, 678901234567890}},
			want: "067d25d39b8514b6b905e0eba2d19242bcf4441e2367527dbceac7a9dd0108a0",
		},
	} {
		tc.Check(name, t)
	}
}

func TestHashProtoReflectMessage(t *testing.T) {
	files := unmarshalProtoRegistryFiles(t, testProtoset)

	t.Run("integers.proto", func(t *testing.T) {
		for name, tc := range map[string]struct {
			md              protoreflect.MessageDescriptor
			json            string
			want            string
			skipEquivalence bool
		}{
			"Int32MessageZero": {
				md:              mdByPath(t, files, "test_protos/schema/proto3/integers.proto", "Int32Message"),
				json:            `{"values": [0, 1, 2]}`,
				want:            "ec28f92dbcce2dc9e38b48cd7725337ca7df40d729b8523a5b3512f7449e8156",
				skipEquivalence: true, // No equivalent JSON: JSON does not have an "integer" type. All numbers are floats.
			},
		} {
			t.Run(name, func(t *testing.T) {
				h := hasher{}

				got := getHash(t, func() ([]byte, error) {
					return h.hashMessage(unmarshalJson(t, tc.md, tc.json))
				})

				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("protohash (-want +got):\n%s", diff)
				}
				if !tc.skipEquivalence {
					if diff := cmp.Diff(tc.want, jsonHash(t, tc.json)); diff != "" {
						t.Errorf("jsonhash (-want +got):\n%s", diff)
					}
				}
			})
		}
	})

}

func TestSkipFields(t *testing.T) {
	for name, tc := range map[string]struct {
		value proto.Message
		want  string
	}{
		"all_ok": {
			value: &pb3_latest.Simple{
				StringField: "bar",
			},
			want: "07c17fb86517d0751756424f3a6f474a5b7a9a5c70913bf15df8656a3a82027a",
		},
		"skipped": {
			value: &pb3_latest.Simple{
				StringField: "bar",
				Int64Field:  123,
			},
			want: "07c17fb86517d0751756424f3a6f474a5b7a9a5c70913bf15df8656a3a82027a",
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := hasher{
				skipFields: map[string]struct{}{
					"schema.proto3.Simple.int64_field": {},
				},
			}
			got, err := h.hashMessage(tc.value.ProtoReflect())
			if err != nil {
				t.Errorf("hashMessage() error = %v", err)
			}

			if diff := cmp.Diff(tc.want, fmt.Sprintf("%x", got)); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}
		})
	}
}

func unmarshalJson(t *testing.T, md protoreflect.MessageDescriptor, json string) protoreflect.Message {
	msg := dynamicpb.NewMessage(md)
	if err := protojson.Unmarshal([]byte(json), msg); err != nil {
		t.Fatal(err)
	}
	return msg
}

func fdByPath(t *testing.T, files *protoregistry.Files, filename string) protoreflect.FileDescriptor {
	fd, err := files.FindFileByPath(filename)
	if err != nil {
		t.Fatal(err)
	}
	return fd
}

func mdByPath(t *testing.T, files *protoregistry.Files, filename, name string) protoreflect.MessageDescriptor {
	fd := fdByPath(t, files, filename)
	md := fd.Messages().ByName(protoreflect.Name(name))
	if md == nil {
		t.Fatal(filename, "| message descriptor not found:", name)
	}
	return md
}

func getHash(t *testing.T, fn func() ([]byte, error)) string {
	hash, err := fn()
	if err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%x", hash)
}

func objectHash(t *testing.T, value interface{}) string {
	objh, err := objecthash.ObjectHash(value)
	if err != nil {
		t.Fatal(err)
	}
	hash := fmt.Sprintf("%x", objh)
	if err != nil {
		t.Fatal(err)
	}
	return hash
}

func jsonString(t *testing.T, value interface{}) string {
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func jsonHash(t *testing.T, value string) string {
	objh, err := objecthash.CommonJSONHash(value)
	if err != nil {
		t.Fatal(err)
	}
	hash := fmt.Sprintf("%x", objh)
	if err != nil {
		t.Fatal(err)
	}
	return hash
}

func unmarshalFileDescriptorSet(t *testing.T, data []byte) *descriptorpb.FileDescriptorSet {
	var dpb descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(data, &dpb); err != nil {
		t.Fatalf("unmarshaling protoset file: %v", err)
	}
	return &dpb
}

func unmarshalProtoRegistryFiles(t *testing.T, data []byte) *protoregistry.Files {
	descriptor := unmarshalFileDescriptorSet(t, data)
	files, err := protodesc.NewFiles(descriptor)
	if err != nil {
		t.Fatal(err)
	}
	return files
}

type stringList []string

// Len reports the number of entries in the List.
// Get, Set, and Truncate panic with out of bound indexes.
func (ss stringList) Len() int {
	return len(ss)
}

// Get retrieves the value at the given index.
// It never returns an invalid value.
func (ss stringList) Get(i int) protoreflect.Value {
	return protoreflect.ValueOf(ss[i])
}

// Set stores a value for the given index.
// When setting a composite type, it is unspecified whether the set
// value aliases the source's memory in any way.
//
// Set is a mutating operation and unsafe for concurrent use.
func (ss stringList) Set(i int, v protoreflect.Value) {
	ss[i] = v.String()
}

// Append appends the provided value to the end of the list.
// When appending a composite type, it is unspecified whether the appended
// value aliases the source's memory in any way.
//
// Append is a mutating operation and unsafe for concurrent use.
func (ss stringList) Append(protoreflect.Value) {
	log.Panicln("not implemented")
}

// AppendMutable appends a new, empty, mutable message value to the end
// of the list and returns it.
// It panics if the list does not contain a message type.
func (ss stringList) AppendMutable() protoreflect.Value {
	log.Panicln("not implemented")
	return protoreflect.Value{}
}

// Truncate truncates the list to a smaller length.
//
// Truncate is a mutating operation and unsafe for concurrent use.
func (ss stringList) Truncate(len int) {
	log.Panicln("not implemented")
}

// NewElement returns a new value for a list element.
// For enums, this returns the first enum value.
// For other scalars, this returns the zero value.
// For messages, this returns a new, empty, mutable value.
func (ss stringList) NewElement() protoreflect.Value {
	log.Panicln("not implemented")
	return protoreflect.Value{}
}

// IsValid reports whether the list is valid.
//
// An invalid list is an empty, read-only value.
//
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (ss stringList) IsValid() bool {
	return true
}

type hashTestCase struct {
	options []Option
	protos  []proto.Message
	obj     interface{}
	json    string
	want    string
}

func (tc *hashTestCase) Check(name string, t *testing.T) {
	for _, msg := range tc.protos {
		t.Run(fmt.Sprintf("%s: %+v", name, msg), func(t *testing.T) {
			h := NewHasher(tc.options...)

			got := getHash(t, func() ([]byte, error) {
				var prm protoreflect.Message
				if msg != nil {
					prm = msg.ProtoReflect()
				}
				return h.HashProto(prm)
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("protohash (-want +got):\n%s", diff)
			}

			if tc.json != "" {
				if diff := cmp.Diff(tc.want, jsonHash(t, tc.json)); diff != "" {
					t.Errorf("jsonhash (-want +got):\n%s", diff)
				}
			}
			if tc.obj != nil {
				if diff := cmp.Diff(tc.want, objectHash(t, tc.obj)); diff != "" {
					t.Errorf("objecthash (-want +got):\n%s", diff)
				}
			}
		})
	}
}
