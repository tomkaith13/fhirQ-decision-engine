//    Copyright 2019 Google Inc.
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
// source: proto/google/fhir/proto/r4/core/resources/endpoint.proto

package endpoint_go_proto

import (
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/google/fhir/go/proto/google/fhir/proto/annotations_go_proto"
	codes_go_proto "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	datatypes_go_proto "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
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

// Auto-generated from StructureDefinition for Endpoint, last updated
// 2019-11-01T09:29:23.356+11:00. The technical details of an endpoint that can
// be used for electronic services. See
// http://hl7.org/fhir/StructureDefinition/Endpoint
type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Logical id of this artifact
	Id *datatypes_go_proto.Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Metadata about the resource
	Meta *datatypes_go_proto.Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *datatypes_go_proto.Uri `protobuf:"bytes,3,opt,name=implicit_rules,json=implicitRules,proto3" json:"implicit_rules,omitempty"`
	// Language of the resource content
	Language *datatypes_go_proto.Code `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *datatypes_go_proto.Narrative `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []*any.Any `protobuf:"bytes,6,rep,name=contained,proto3" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []*datatypes_go_proto.Extension `protobuf:"bytes,8,rep,name=extension,proto3" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []*datatypes_go_proto.Extension `protobuf:"bytes,9,rep,name=modifier_extension,json=modifierExtension,proto3" json:"modifier_extension,omitempty"`
	// Identifies this endpoint across multiple systems
	Identifier []*datatypes_go_proto.Identifier `protobuf:"bytes,10,rep,name=identifier,proto3" json:"identifier,omitempty"`
	Status     *Endpoint_StatusCode             `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	// Protocol/Profile/Standard to be used with this endpoint connection
	ConnectionType *datatypes_go_proto.Coding `protobuf:"bytes,12,opt,name=connection_type,json=connectionType,proto3" json:"connection_type,omitempty"`
	// A name that this endpoint can be identified by
	Name *datatypes_go_proto.String `protobuf:"bytes,13,opt,name=name,proto3" json:"name,omitempty"`
	// Organization that manages this endpoint (might not be the organization that
	// exposes the endpoint)
	ManagingOrganization *datatypes_go_proto.Reference `protobuf:"bytes,14,opt,name=managing_organization,json=managingOrganization,proto3" json:"managing_organization,omitempty"`
	// Contact details for source (e.g. troubleshooting)
	Contact []*datatypes_go_proto.ContactPoint `protobuf:"bytes,15,rep,name=contact,proto3" json:"contact,omitempty"`
	// Interval the endpoint is expected to be operational
	Period *datatypes_go_proto.Period `protobuf:"bytes,16,opt,name=period,proto3" json:"period,omitempty"`
	// The type of content that may be used at this endpoint (e.g. XDS Discharge
	// summaries)
	PayloadType     []*datatypes_go_proto.CodeableConcept `protobuf:"bytes,17,rep,name=payload_type,json=payloadType,proto3" json:"payload_type,omitempty"`
	PayloadMimeType []*Endpoint_PayloadMimeTypeCode       `protobuf:"bytes,18,rep,name=payload_mime_type,json=payloadMimeType,proto3" json:"payload_mime_type,omitempty"`
	// The technical base address for connecting to this endpoint
	Address *datatypes_go_proto.Url `protobuf:"bytes,19,opt,name=address,proto3" json:"address,omitempty"`
	// Usage depends on the channel type
	Header []*datatypes_go_proto.String `protobuf:"bytes,20,rep,name=header,proto3" json:"header,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetId() *datatypes_go_proto.Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Endpoint) GetMeta() *datatypes_go_proto.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Endpoint) GetImplicitRules() *datatypes_go_proto.Uri {
	if x != nil {
		return x.ImplicitRules
	}
	return nil
}

func (x *Endpoint) GetLanguage() *datatypes_go_proto.Code {
	if x != nil {
		return x.Language
	}
	return nil
}

func (x *Endpoint) GetText() *datatypes_go_proto.Narrative {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *Endpoint) GetContained() []*any.Any {
	if x != nil {
		return x.Contained
	}
	return nil
}

func (x *Endpoint) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *Endpoint) GetModifierExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.ModifierExtension
	}
	return nil
}

func (x *Endpoint) GetIdentifier() []*datatypes_go_proto.Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *Endpoint) GetStatus() *Endpoint_StatusCode {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Endpoint) GetConnectionType() *datatypes_go_proto.Coding {
	if x != nil {
		return x.ConnectionType
	}
	return nil
}

func (x *Endpoint) GetName() *datatypes_go_proto.String {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Endpoint) GetManagingOrganization() *datatypes_go_proto.Reference {
	if x != nil {
		return x.ManagingOrganization
	}
	return nil
}

func (x *Endpoint) GetContact() []*datatypes_go_proto.ContactPoint {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Endpoint) GetPeriod() *datatypes_go_proto.Period {
	if x != nil {
		return x.Period
	}
	return nil
}

func (x *Endpoint) GetPayloadType() []*datatypes_go_proto.CodeableConcept {
	if x != nil {
		return x.PayloadType
	}
	return nil
}

func (x *Endpoint) GetPayloadMimeType() []*Endpoint_PayloadMimeTypeCode {
	if x != nil {
		return x.PayloadMimeType
	}
	return nil
}

func (x *Endpoint) GetAddress() *datatypes_go_proto.Url {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Endpoint) GetHeader() []*datatypes_go_proto.String {
	if x != nil {
		return x.Header
	}
	return nil
}

// active | suspended | error | off | entered-in-error | test
type Endpoint_StatusCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     codes_go_proto.EndpointStatusCode_Value `protobuf:"varint,1,opt,name=value,proto3,enum=google.fhir.r4.core.EndpointStatusCode_Value" json:"value,omitempty"`
	Id        *datatypes_go_proto.String              `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Extension []*datatypes_go_proto.Extension         `protobuf:"bytes,3,rep,name=extension,proto3" json:"extension,omitempty"`
}

func (x *Endpoint_StatusCode) Reset() {
	*x = Endpoint_StatusCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint_StatusCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint_StatusCode) ProtoMessage() {}

func (x *Endpoint_StatusCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint_StatusCode.ProtoReflect.Descriptor instead.
func (*Endpoint_StatusCode) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Endpoint_StatusCode) GetValue() codes_go_proto.EndpointStatusCode_Value {
	if x != nil {
		return x.Value
	}
	return codes_go_proto.EndpointStatusCode_INVALID_UNINITIALIZED
}

func (x *Endpoint_StatusCode) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Endpoint_StatusCode) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

// Mimetype to send. If not specified, the content could be anything
// (including no payload, if the connectionType defined this)
type Endpoint_PayloadMimeTypeCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *datatypes_go_proto.String      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Extension []*datatypes_go_proto.Extension `protobuf:"bytes,3,rep,name=extension,proto3" json:"extension,omitempty"`
	// This valueset is not enumerable, and so is represented as a string.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Endpoint_PayloadMimeTypeCode) Reset() {
	*x = Endpoint_PayloadMimeTypeCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint_PayloadMimeTypeCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint_PayloadMimeTypeCode) ProtoMessage() {}

func (x *Endpoint_PayloadMimeTypeCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint_PayloadMimeTypeCode.ProtoReflect.Descriptor instead.
func (*Endpoint_PayloadMimeTypeCode) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Endpoint_PayloadMimeTypeCode) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Endpoint_PayloadMimeTypeCode) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *Endpoint_PayloadMimeTypeCode) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_google_fhir_proto_r4_core_resources_endpoint_proto protoreflect.FileDescriptor

var file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDesc = []byte{
	0x0a, 0x38, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x34, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x34, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x34, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x0e, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x27, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0e, 0x69, 0x6d, 0x70, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72,
	0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x72, 0x69, 0x52, 0x0d, 0x69, 0x6d, 0x70, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x61, 0x72, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72,
	0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x4c, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x67, 0x0a, 0x15, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x12,
	0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x4f, 0x0a, 0x0c, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72,
	0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x0b,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x11, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x69, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x55, 0x72, 0x6c, 0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0xa8, 0x02, 0x0a, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x6a, 0xc0, 0x9f, 0xe3, 0xb6,
	0x05, 0x01, 0x8a, 0xf9, 0x83, 0xb2, 0x05, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68,
	0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x53, 0x65, 0x74, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2d, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x9a, 0xb5, 0x8e, 0x93, 0x06, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f,
	0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x82, 0x02, 0x0a, 0x13, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2b,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x64, 0xc0, 0x9f, 0xe3, 0xb6, 0x05, 0x01, 0x8a, 0xf9, 0x83, 0xb2, 0x05, 0x26, 0x68, 0x74, 0x74,
	0x70, 0x3a, 0x2f, 0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72,
	0x2f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x74, 0x2f, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x9a, 0xb5, 0x8e, 0x93, 0x06, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f,
	0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x3a, 0x3c, 0xc0, 0x9f, 0xe3,
	0xb6, 0x05, 0x03, 0xb2, 0xfe, 0xe4, 0x97, 0x06, 0x30, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f,
	0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x42,
	0x78, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68,
	0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x66, 0x68, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x72, 0x34, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x98, 0xc6, 0xb0, 0xb5, 0x07, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescOnce sync.Once
	file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescData = file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDesc
)

func file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescGZIP() []byte {
	file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescOnce.Do(func() {
		file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescData)
	})
	return file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDescData
}

var file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_goTypes = []interface{}{
	(*Endpoint)(nil),                             // 0: google.fhir.r4.core.Endpoint
	(*Endpoint_StatusCode)(nil),                  // 1: google.fhir.r4.core.Endpoint.StatusCode
	(*Endpoint_PayloadMimeTypeCode)(nil),         // 2: google.fhir.r4.core.Endpoint.PayloadMimeTypeCode
	(*datatypes_go_proto.Id)(nil),                // 3: google.fhir.r4.core.Id
	(*datatypes_go_proto.Meta)(nil),              // 4: google.fhir.r4.core.Meta
	(*datatypes_go_proto.Uri)(nil),               // 5: google.fhir.r4.core.Uri
	(*datatypes_go_proto.Code)(nil),              // 6: google.fhir.r4.core.Code
	(*datatypes_go_proto.Narrative)(nil),         // 7: google.fhir.r4.core.Narrative
	(*any.Any)(nil),                              // 8: google.protobuf.Any
	(*datatypes_go_proto.Extension)(nil),         // 9: google.fhir.r4.core.Extension
	(*datatypes_go_proto.Identifier)(nil),        // 10: google.fhir.r4.core.Identifier
	(*datatypes_go_proto.Coding)(nil),            // 11: google.fhir.r4.core.Coding
	(*datatypes_go_proto.String)(nil),            // 12: google.fhir.r4.core.String
	(*datatypes_go_proto.Reference)(nil),         // 13: google.fhir.r4.core.Reference
	(*datatypes_go_proto.ContactPoint)(nil),      // 14: google.fhir.r4.core.ContactPoint
	(*datatypes_go_proto.Period)(nil),            // 15: google.fhir.r4.core.Period
	(*datatypes_go_proto.CodeableConcept)(nil),   // 16: google.fhir.r4.core.CodeableConcept
	(*datatypes_go_proto.Url)(nil),               // 17: google.fhir.r4.core.Url
	(codes_go_proto.EndpointStatusCode_Value)(0), // 18: google.fhir.r4.core.EndpointStatusCode.Value
}
var file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_depIdxs = []int32{
	3,  // 0: google.fhir.r4.core.Endpoint.id:type_name -> google.fhir.r4.core.Id
	4,  // 1: google.fhir.r4.core.Endpoint.meta:type_name -> google.fhir.r4.core.Meta
	5,  // 2: google.fhir.r4.core.Endpoint.implicit_rules:type_name -> google.fhir.r4.core.Uri
	6,  // 3: google.fhir.r4.core.Endpoint.language:type_name -> google.fhir.r4.core.Code
	7,  // 4: google.fhir.r4.core.Endpoint.text:type_name -> google.fhir.r4.core.Narrative
	8,  // 5: google.fhir.r4.core.Endpoint.contained:type_name -> google.protobuf.Any
	9,  // 6: google.fhir.r4.core.Endpoint.extension:type_name -> google.fhir.r4.core.Extension
	9,  // 7: google.fhir.r4.core.Endpoint.modifier_extension:type_name -> google.fhir.r4.core.Extension
	10, // 8: google.fhir.r4.core.Endpoint.identifier:type_name -> google.fhir.r4.core.Identifier
	1,  // 9: google.fhir.r4.core.Endpoint.status:type_name -> google.fhir.r4.core.Endpoint.StatusCode
	11, // 10: google.fhir.r4.core.Endpoint.connection_type:type_name -> google.fhir.r4.core.Coding
	12, // 11: google.fhir.r4.core.Endpoint.name:type_name -> google.fhir.r4.core.String
	13, // 12: google.fhir.r4.core.Endpoint.managing_organization:type_name -> google.fhir.r4.core.Reference
	14, // 13: google.fhir.r4.core.Endpoint.contact:type_name -> google.fhir.r4.core.ContactPoint
	15, // 14: google.fhir.r4.core.Endpoint.period:type_name -> google.fhir.r4.core.Period
	16, // 15: google.fhir.r4.core.Endpoint.payload_type:type_name -> google.fhir.r4.core.CodeableConcept
	2,  // 16: google.fhir.r4.core.Endpoint.payload_mime_type:type_name -> google.fhir.r4.core.Endpoint.PayloadMimeTypeCode
	17, // 17: google.fhir.r4.core.Endpoint.address:type_name -> google.fhir.r4.core.Url
	12, // 18: google.fhir.r4.core.Endpoint.header:type_name -> google.fhir.r4.core.String
	18, // 19: google.fhir.r4.core.Endpoint.StatusCode.value:type_name -> google.fhir.r4.core.EndpointStatusCode.Value
	12, // 20: google.fhir.r4.core.Endpoint.StatusCode.id:type_name -> google.fhir.r4.core.String
	9,  // 21: google.fhir.r4.core.Endpoint.StatusCode.extension:type_name -> google.fhir.r4.core.Extension
	12, // 22: google.fhir.r4.core.Endpoint.PayloadMimeTypeCode.id:type_name -> google.fhir.r4.core.String
	9,  // 23: google.fhir.r4.core.Endpoint.PayloadMimeTypeCode.extension:type_name -> google.fhir.r4.core.Extension
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	24, // [24:24] is the sub-list for extension extendee
	0,  // [0:24] is the sub-list for field type_name
}

func init() { file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_init() }
func file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_init() {
	if File_proto_google_fhir_proto_r4_core_resources_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint_StatusCode); i {
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
		file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint_PayloadMimeTypeCode); i {
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
			RawDescriptor: file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_goTypes,
		DependencyIndexes: file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_depIdxs,
		MessageInfos:      file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_msgTypes,
	}.Build()
	File_proto_google_fhir_proto_r4_core_resources_endpoint_proto = out.File
	file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_rawDesc = nil
	file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_goTypes = nil
	file_proto_google_fhir_proto_r4_core_resources_endpoint_proto_depIdxs = nil
}
