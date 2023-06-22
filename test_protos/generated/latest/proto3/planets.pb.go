// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: test_protos/schema/proto3/planets.proto

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

type PlanetV1 int32

const (
	PlanetV1_UNKNOWN_V1 PlanetV1 = 0
	PlanetV1_EARTH_V1   PlanetV1 = 3
	PlanetV1_MERCURY_V1 PlanetV1 = 1
	PlanetV1_VENUS_V1   PlanetV1 = 2
	PlanetV1_MARS_V1    PlanetV1 = 4
	PlanetV1_JUPITER_V1 PlanetV1 = 5
	PlanetV1_SATURN_V1  PlanetV1 = 6
)

// Enum value maps for PlanetV1.
var (
	PlanetV1_name = map[int32]string{
		0: "UNKNOWN_V1",
		3: "EARTH_V1",
		1: "MERCURY_V1",
		2: "VENUS_V1",
		4: "MARS_V1",
		5: "JUPITER_V1",
		6: "SATURN_V1",
	}
	PlanetV1_value = map[string]int32{
		"UNKNOWN_V1": 0,
		"EARTH_V1":   3,
		"MERCURY_V1": 1,
		"VENUS_V1":   2,
		"MARS_V1":    4,
		"JUPITER_V1": 5,
		"SATURN_V1":  6,
	}
)

func (x PlanetV1) Enum() *PlanetV1 {
	p := new(PlanetV1)
	*p = x
	return p
}

func (x PlanetV1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanetV1) Descriptor() protoreflect.EnumDescriptor {
	return file_test_protos_schema_proto3_planets_proto_enumTypes[0].Descriptor()
}

func (PlanetV1) Type() protoreflect.EnumType {
	return &file_test_protos_schema_proto3_planets_proto_enumTypes[0]
}

func (x PlanetV1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanetV1.Descriptor instead.
func (PlanetV1) EnumDescriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_planets_proto_rawDescGZIP(), []int{0}
}

type PlanetV2 int32

const (
	PlanetV2_UNKNOWN_V2 PlanetV2 = 0
	PlanetV2_EARTH_V2   PlanetV2 = 3
	PlanetV2_MERCURY_V2 PlanetV2 = 1
	PlanetV2_VENUS_V2   PlanetV2 = 2
	PlanetV2_MARS_V2    PlanetV2 = 4
	PlanetV2_JUPITER_V2 PlanetV2 = 5
	PlanetV2_SATURN_V2  PlanetV2 = 6
	PlanetV2_URANUS_V2  PlanetV2 = 7
	PlanetV2_NEPTUNE_V2 PlanetV2 = 8
	PlanetV2_PLUTO_V2   PlanetV2 = 9
)

// Enum value maps for PlanetV2.
var (
	PlanetV2_name = map[int32]string{
		0: "UNKNOWN_V2",
		3: "EARTH_V2",
		1: "MERCURY_V2",
		2: "VENUS_V2",
		4: "MARS_V2",
		5: "JUPITER_V2",
		6: "SATURN_V2",
		7: "URANUS_V2",
		8: "NEPTUNE_V2",
		9: "PLUTO_V2",
	}
	PlanetV2_value = map[string]int32{
		"UNKNOWN_V2": 0,
		"EARTH_V2":   3,
		"MERCURY_V2": 1,
		"VENUS_V2":   2,
		"MARS_V2":    4,
		"JUPITER_V2": 5,
		"SATURN_V2":  6,
		"URANUS_V2":  7,
		"NEPTUNE_V2": 8,
		"PLUTO_V2":   9,
	}
)

func (x PlanetV2) Enum() *PlanetV2 {
	p := new(PlanetV2)
	*p = x
	return p
}

func (x PlanetV2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanetV2) Descriptor() protoreflect.EnumDescriptor {
	return file_test_protos_schema_proto3_planets_proto_enumTypes[1].Descriptor()
}

func (PlanetV2) Type() protoreflect.EnumType {
	return &file_test_protos_schema_proto3_planets_proto_enumTypes[1]
}

func (x PlanetV2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanetV2.Descriptor instead.
func (PlanetV2) EnumDescriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_planets_proto_rawDescGZIP(), []int{1}
}

type PlanetV3 int32

const (
	PlanetV3_UNKNOWN_V3          PlanetV3 = 0
	PlanetV3_EARTH_V3            PlanetV3 = 3
	PlanetV3_MERCURY_V3          PlanetV3 = 1
	PlanetV3_VENUS_V3            PlanetV3 = 2
	PlanetV3_MARS_V3             PlanetV3 = 4
	PlanetV3_JUPITER_V3          PlanetV3 = 5
	PlanetV3_SATURN_V3           PlanetV3 = 6
	PlanetV3_URANUS_V3           PlanetV3 = 7
	PlanetV3_NEPTUNE_V3          PlanetV3 = 8
	PlanetV3_DEPRECATED_PLUTO_V3 PlanetV3 = 9
)

// Enum value maps for PlanetV3.
var (
	PlanetV3_name = map[int32]string{
		0: "UNKNOWN_V3",
		3: "EARTH_V3",
		1: "MERCURY_V3",
		2: "VENUS_V3",
		4: "MARS_V3",
		5: "JUPITER_V3",
		6: "SATURN_V3",
		7: "URANUS_V3",
		8: "NEPTUNE_V3",
		9: "DEPRECATED_PLUTO_V3",
	}
	PlanetV3_value = map[string]int32{
		"UNKNOWN_V3":          0,
		"EARTH_V3":            3,
		"MERCURY_V3":          1,
		"VENUS_V3":            2,
		"MARS_V3":             4,
		"JUPITER_V3":          5,
		"SATURN_V3":           6,
		"URANUS_V3":           7,
		"NEPTUNE_V3":          8,
		"DEPRECATED_PLUTO_V3": 9,
	}
)

func (x PlanetV3) Enum() *PlanetV3 {
	p := new(PlanetV3)
	*p = x
	return p
}

func (x PlanetV3) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanetV3) Descriptor() protoreflect.EnumDescriptor {
	return file_test_protos_schema_proto3_planets_proto_enumTypes[2].Descriptor()
}

func (PlanetV3) Type() protoreflect.EnumType {
	return &file_test_protos_schema_proto3_planets_proto_enumTypes[2]
}

func (x PlanetV3) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanetV3.Descriptor instead.
func (PlanetV3) EnumDescriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_planets_proto_rawDescGZIP(), []int{2}
}

type MyFavoritePlanetsV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planets []PlanetV1 `protobuf:"varint,1,rep,packed,name=planets,proto3,enum=schema.proto3.PlanetV1" json:"planets,omitempty"`
}

func (x *MyFavoritePlanetsV1) Reset() {
	*x = MyFavoritePlanetsV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_schema_proto3_planets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyFavoritePlanetsV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyFavoritePlanetsV1) ProtoMessage() {}

func (x *MyFavoritePlanetsV1) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_schema_proto3_planets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyFavoritePlanetsV1.ProtoReflect.Descriptor instead.
func (*MyFavoritePlanetsV1) Descriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_planets_proto_rawDescGZIP(), []int{0}
}

func (x *MyFavoritePlanetsV1) GetPlanets() []PlanetV1 {
	if x != nil {
		return x.Planets
	}
	return nil
}

type MyFavoritePlanetsV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planets []PlanetV2 `protobuf:"varint,1,rep,packed,name=planets,proto3,enum=schema.proto3.PlanetV2" json:"planets,omitempty"`
}

func (x *MyFavoritePlanetsV2) Reset() {
	*x = MyFavoritePlanetsV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_schema_proto3_planets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyFavoritePlanetsV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyFavoritePlanetsV2) ProtoMessage() {}

func (x *MyFavoritePlanetsV2) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_schema_proto3_planets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyFavoritePlanetsV2.ProtoReflect.Descriptor instead.
func (*MyFavoritePlanetsV2) Descriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_planets_proto_rawDescGZIP(), []int{1}
}

func (x *MyFavoritePlanetsV2) GetPlanets() []PlanetV2 {
	if x != nil {
		return x.Planets
	}
	return nil
}

type MyFavoritePlanetsV3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planets []PlanetV3 `protobuf:"varint,1,rep,packed,name=planets,proto3,enum=schema.proto3.PlanetV3" json:"planets,omitempty"`
}

func (x *MyFavoritePlanetsV3) Reset() {
	*x = MyFavoritePlanetsV3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_schema_proto3_planets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyFavoritePlanetsV3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyFavoritePlanetsV3) ProtoMessage() {}

func (x *MyFavoritePlanetsV3) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_schema_proto3_planets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyFavoritePlanetsV3.ProtoReflect.Descriptor instead.
func (*MyFavoritePlanetsV3) Descriptor() ([]byte, []int) {
	return file_test_protos_schema_proto3_planets_proto_rawDescGZIP(), []int{2}
}

func (x *MyFavoritePlanetsV3) GetPlanets() []PlanetV3 {
	if x != nil {
		return x.Planets
	}
	return nil
}

var File_test_protos_schema_proto3_planets_proto protoreflect.FileDescriptor

var file_test_protos_schema_proto3_planets_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x22, 0x48, 0x0a, 0x13, 0x4d, 0x79, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x56, 0x31, 0x12,
	0x31, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x56, 0x31, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x73, 0x22, 0x48, 0x0a, 0x13, 0x4d, 0x79, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x56, 0x32, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x56, 0x32, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x22, 0x48, 0x0a, 0x13,
	0x4d, 0x79, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x73, 0x56, 0x33, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x56, 0x33, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x2a, 0x72, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x56, 0x31, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x56, 0x31,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x41, 0x52, 0x54, 0x48, 0x5f, 0x56, 0x31, 0x10, 0x03,
	0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x52, 0x43, 0x55, 0x52, 0x59, 0x5f, 0x56, 0x31, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x4e, 0x55, 0x53, 0x5f, 0x56, 0x31, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x4d, 0x41, 0x52, 0x53, 0x5f, 0x56, 0x31, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4a,
	0x55, 0x50, 0x49, 0x54, 0x45, 0x52, 0x5f, 0x56, 0x31, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x41, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x56, 0x31, 0x10, 0x06, 0x2a, 0x9f, 0x01, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x56, 0x32, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x56, 0x32, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x41, 0x52, 0x54, 0x48,
	0x5f, 0x56, 0x32, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x52, 0x43, 0x55, 0x52, 0x59,
	0x5f, 0x56, 0x32, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x4e, 0x55, 0x53, 0x5f, 0x56,
	0x32, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x52, 0x53, 0x5f, 0x56, 0x32, 0x10, 0x04,
	0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x55, 0x50, 0x49, 0x54, 0x45, 0x52, 0x5f, 0x56, 0x32, 0x10, 0x05,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x41, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x56, 0x32, 0x10, 0x06, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x52, 0x41, 0x4e, 0x55, 0x53, 0x5f, 0x56, 0x32, 0x10, 0x07, 0x12, 0x0e,
	0x0a, 0x0a, 0x4e, 0x45, 0x50, 0x54, 0x55, 0x4e, 0x45, 0x5f, 0x56, 0x32, 0x10, 0x08, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x4c, 0x55, 0x54, 0x4f, 0x5f, 0x56, 0x32, 0x10, 0x09, 0x2a, 0xaa, 0x01, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x56, 0x33, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x56, 0x33, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x41, 0x52,
	0x54, 0x48, 0x5f, 0x56, 0x33, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x52, 0x43, 0x55,
	0x52, 0x59, 0x5f, 0x56, 0x33, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x4e, 0x55, 0x53,
	0x5f, 0x56, 0x33, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x52, 0x53, 0x5f, 0x56, 0x33,
	0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x55, 0x50, 0x49, 0x54, 0x45, 0x52, 0x5f, 0x56, 0x33,
	0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x41, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x56, 0x33, 0x10,
	0x06, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x52, 0x41, 0x4e, 0x55, 0x53, 0x5f, 0x56, 0x33, 0x10, 0x07,
	0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x45, 0x50, 0x54, 0x55, 0x4e, 0x45, 0x5f, 0x56, 0x33, 0x10, 0x08,
	0x12, 0x17, 0x0a, 0x13, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x50,
	0x4c, 0x55, 0x54, 0x4f, 0x5f, 0x56, 0x33, 0x10, 0x09, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x68, 0x61, 0x73, 0x68, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_protos_schema_proto3_planets_proto_rawDescOnce sync.Once
	file_test_protos_schema_proto3_planets_proto_rawDescData = file_test_protos_schema_proto3_planets_proto_rawDesc
)

func file_test_protos_schema_proto3_planets_proto_rawDescGZIP() []byte {
	file_test_protos_schema_proto3_planets_proto_rawDescOnce.Do(func() {
		file_test_protos_schema_proto3_planets_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_protos_schema_proto3_planets_proto_rawDescData)
	})
	return file_test_protos_schema_proto3_planets_proto_rawDescData
}

var file_test_protos_schema_proto3_planets_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_test_protos_schema_proto3_planets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_protos_schema_proto3_planets_proto_goTypes = []interface{}{
	(PlanetV1)(0),               // 0: schema.proto3.PlanetV1
	(PlanetV2)(0),               // 1: schema.proto3.PlanetV2
	(PlanetV3)(0),               // 2: schema.proto3.PlanetV3
	(*MyFavoritePlanetsV1)(nil), // 3: schema.proto3.MyFavoritePlanetsV1
	(*MyFavoritePlanetsV2)(nil), // 4: schema.proto3.MyFavoritePlanetsV2
	(*MyFavoritePlanetsV3)(nil), // 5: schema.proto3.MyFavoritePlanetsV3
}
var file_test_protos_schema_proto3_planets_proto_depIdxs = []int32{
	0, // 0: schema.proto3.MyFavoritePlanetsV1.planets:type_name -> schema.proto3.PlanetV1
	1, // 1: schema.proto3.MyFavoritePlanetsV2.planets:type_name -> schema.proto3.PlanetV2
	2, // 2: schema.proto3.MyFavoritePlanetsV3.planets:type_name -> schema.proto3.PlanetV3
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_test_protos_schema_proto3_planets_proto_init() }
func file_test_protos_schema_proto3_planets_proto_init() {
	if File_test_protos_schema_proto3_planets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_protos_schema_proto3_planets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyFavoritePlanetsV1); i {
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
		file_test_protos_schema_proto3_planets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyFavoritePlanetsV2); i {
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
		file_test_protos_schema_proto3_planets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyFavoritePlanetsV3); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_protos_schema_proto3_planets_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_protos_schema_proto3_planets_proto_goTypes,
		DependencyIndexes: file_test_protos_schema_proto3_planets_proto_depIdxs,
		EnumInfos:         file_test_protos_schema_proto3_planets_proto_enumTypes,
		MessageInfos:      file_test_protos_schema_proto3_planets_proto_msgTypes,
	}.Build()
	File_test_protos_schema_proto3_planets_proto = out.File
	file_test_protos_schema_proto3_planets_proto_rawDesc = nil
	file_test_protos_schema_proto3_planets_proto_goTypes = nil
	file_test_protos_schema_proto3_planets_proto_depIdxs = nil
}
