//    Copyright 2018 Google Inc.
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        https://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto/google/fhir/proto/stu3/fhirproto_extensions.proto

package fhirproto_extensions_go_proto

import (
	_ "github.com/google/fhir/go/proto/google/fhir/proto/annotations_go_proto"
	datatypes_go_proto "github.com/google/fhir/go/proto/google/fhir/proto/stu3/datatypes_go_proto"
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

// Whether or not the primitive is missing a value.
// See https://g.co/fhir/StructureDefinition/primitiveHasNoValue
type PrimitiveHasNoValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// xml:id (or equivalent in JSON)
	Id *datatypes_go_proto.String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Value of extension
	ValueBoolean *datatypes_go_proto.Boolean `protobuf:"bytes,3,opt,name=value_boolean,json=valueBoolean,proto3" json:"value_boolean,omitempty"`
}

func (x *PrimitiveHasNoValue) Reset() {
	*x = PrimitiveHasNoValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimitiveHasNoValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimitiveHasNoValue) ProtoMessage() {}

func (x *PrimitiveHasNoValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimitiveHasNoValue.ProtoReflect.Descriptor instead.
func (*PrimitiveHasNoValue) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescGZIP(), []int{0}
}

func (x *PrimitiveHasNoValue) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *PrimitiveHasNoValue) GetValueBoolean() *datatypes_go_proto.Boolean {
	if x != nil {
		return x.ValueBoolean
	}
	return nil
}

// Base64Binary rendering parameters.
// See https://g.co/fhir/StructureDefinition/base64Binary-separatorStride
type Base64BinarySeparatorStride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// xml:id (or equivalent in JSON)
	Id *datatypes_go_proto.String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The separator, usually a sequence of spaces.
	Separator *datatypes_go_proto.String `protobuf:"bytes,4,opt,name=separator,proto3" json:"separator,omitempty"`
	// The stride.
	Stride *datatypes_go_proto.PositiveInt `protobuf:"bytes,5,opt,name=stride,proto3" json:"stride,omitempty"`
}

func (x *Base64BinarySeparatorStride) Reset() {
	*x = Base64BinarySeparatorStride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base64BinarySeparatorStride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base64BinarySeparatorStride) ProtoMessage() {}

func (x *Base64BinarySeparatorStride) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base64BinarySeparatorStride.ProtoReflect.Descriptor instead.
func (*Base64BinarySeparatorStride) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescGZIP(), []int{1}
}

func (x *Base64BinarySeparatorStride) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Base64BinarySeparatorStride) GetSeparator() *datatypes_go_proto.String {
	if x != nil {
		return x.Separator
	}
	return nil
}

func (x *Base64BinarySeparatorStride) GetStride() *datatypes_go_proto.PositiveInt {
	if x != nil {
		return x.Stride
	}
	return nil
}

var File_proto_google_fhir_proto_stu3_fhirproto_extensions_proto protoreflect.FileDescriptor

var file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDesc = []byte{
	0x0a, 0x37, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x33, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x73, 0x74, 0x75, 0x33, 0x2e, 0x66, 0x68, 0x69, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x33, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89,
	0x02, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x4e,
	0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72,
	0x2e, 0x73, 0x74, 0x75, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x73, 0x74, 0x75, 0x33,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x52, 0x0c,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x3a, 0x7c, 0xc0, 0x9f,
	0xe3, 0xb6, 0x05, 0x02, 0x9a, 0xb5, 0x8e, 0x93, 0x06, 0x31, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f,
	0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0xb2, 0xfe, 0xe4, 0x97, 0x06,
	0x39, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x2e, 0x63, 0x6f, 0x2f, 0x66, 0x68,
	0x69, 0x72, 0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x48, 0x61, 0x73, 0x4e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb4, 0x03, 0x0a, 0x1b, 0x42,
	0x61, 0x73, 0x65, 0x36, 0x34, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x53, 0x65, 0x70, 0x61, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x69, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x73, 0x74, 0x75, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12, 0x6e, 0x0a, 0x09, 0x73, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x73, 0x74, 0x75, 0x33,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x30, 0xf0,
	0xd0, 0x87, 0xeb, 0x04, 0x01, 0xf2, 0xbe, 0xc0, 0xa4, 0x07, 0x24, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x28, 0x29, 0x20, 0x21, 0x3d,
	0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x28, 0x29, 0x52,
	0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x6d, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x73, 0x74, 0x75, 0x33, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x42,
	0x30, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0xf2, 0xbe, 0xc0, 0xa4, 0x07, 0x24, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x28, 0x29, 0x20,
	0x21, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x28,
	0x29, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x85, 0x01, 0xc0, 0x9f, 0xe3, 0xb6,
	0x05, 0x02, 0x9a, 0xb5, 0x8e, 0x93, 0x06, 0x31, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68,
	0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0xb2, 0xfe, 0xe4, 0x97, 0x06, 0x42, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x2e, 0x63, 0x6f, 0x2f, 0x66, 0x68, 0x69, 0x72,
	0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x2d, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x69, 0x64,
	0x65, 0x42, 0x7e, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x73, 0x74, 0x75, 0x33, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68,
	0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x33, 0x2f, 0x66, 0x68,
	0x69, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x98, 0xc6, 0xb0, 0xb5, 0x07,
	0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescOnce sync.Once
	file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescData = file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDesc
)

func file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescGZIP() []byte {
	file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescOnce.Do(func() {
		file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescData)
	})
	return file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDescData
}

var file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_goTypes = []interface{}{
	(*PrimitiveHasNoValue)(nil),            // 0: google.fhir.stu3.fhirproto.PrimitiveHasNoValue
	(*Base64BinarySeparatorStride)(nil),    // 1: google.fhir.stu3.fhirproto.Base64BinarySeparatorStride
	(*datatypes_go_proto.String)(nil),      // 2: google.fhir.stu3.proto.String
	(*datatypes_go_proto.Boolean)(nil),     // 3: google.fhir.stu3.proto.Boolean
	(*datatypes_go_proto.PositiveInt)(nil), // 4: google.fhir.stu3.proto.PositiveInt
}
var file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_depIdxs = []int32{
	2, // 0: google.fhir.stu3.fhirproto.PrimitiveHasNoValue.id:type_name -> google.fhir.stu3.proto.String
	3, // 1: google.fhir.stu3.fhirproto.PrimitiveHasNoValue.value_boolean:type_name -> google.fhir.stu3.proto.Boolean
	2, // 2: google.fhir.stu3.fhirproto.Base64BinarySeparatorStride.id:type_name -> google.fhir.stu3.proto.String
	2, // 3: google.fhir.stu3.fhirproto.Base64BinarySeparatorStride.separator:type_name -> google.fhir.stu3.proto.String
	4, // 4: google.fhir.stu3.fhirproto.Base64BinarySeparatorStride.stride:type_name -> google.fhir.stu3.proto.PositiveInt
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_init() }
func file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_init() {
	if File_proto_google_fhir_proto_stu3_fhirproto_extensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimitiveHasNoValue); i {
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
		file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base64BinarySeparatorStride); i {
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
			RawDescriptor: file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_goTypes,
		DependencyIndexes: file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_depIdxs,
		MessageInfos:      file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_msgTypes,
	}.Build()
	File_proto_google_fhir_proto_stu3_fhirproto_extensions_proto = out.File
	file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_rawDesc = nil
	file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_goTypes = nil
	file_proto_google_fhir_proto_stu3_fhirproto_extensions_proto_depIdxs = nil
}
