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
// source: proto/google/fhir/proto/r4/core/resources/enrollment_response.proto

package enrollment_response_go_proto

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

// Auto-generated from StructureDefinition for EnrollmentResponse, last updated
// 2019-11-01T09:29:23.356+11:00. EnrollmentResponse resource. See
// http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
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
	// Business Identifier
	Identifier []*datatypes_go_proto.Identifier `protobuf:"bytes,10,rep,name=identifier,proto3" json:"identifier,omitempty"`
	Status     *EnrollmentResponse_StatusCode   `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	// Claim reference
	Request *datatypes_go_proto.Reference   `protobuf:"bytes,12,opt,name=request,proto3" json:"request,omitempty"`
	Outcome *EnrollmentResponse_OutcomeCode `protobuf:"bytes,13,opt,name=outcome,proto3" json:"outcome,omitempty"`
	// Disposition Message
	Disposition *datatypes_go_proto.String `protobuf:"bytes,14,opt,name=disposition,proto3" json:"disposition,omitempty"`
	// Creation date
	Created *datatypes_go_proto.DateTime `protobuf:"bytes,15,opt,name=created,proto3" json:"created,omitempty"`
	// Insurer
	Organization *datatypes_go_proto.Reference `protobuf:"bytes,16,opt,name=organization,proto3" json:"organization,omitempty"`
	// Responsible practitioner
	RequestProvider *datatypes_go_proto.Reference `protobuf:"bytes,17,opt,name=request_provider,json=requestProvider,proto3" json:"request_provider,omitempty"`
}

func (x *EnrollmentResponse) Reset() {
	*x = EnrollmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollmentResponse) ProtoMessage() {}

func (x *EnrollmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollmentResponse.ProtoReflect.Descriptor instead.
func (*EnrollmentResponse) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescGZIP(), []int{0}
}

func (x *EnrollmentResponse) GetId() *datatypes_go_proto.Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EnrollmentResponse) GetMeta() *datatypes_go_proto.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *EnrollmentResponse) GetImplicitRules() *datatypes_go_proto.Uri {
	if x != nil {
		return x.ImplicitRules
	}
	return nil
}

func (x *EnrollmentResponse) GetLanguage() *datatypes_go_proto.Code {
	if x != nil {
		return x.Language
	}
	return nil
}

func (x *EnrollmentResponse) GetText() *datatypes_go_proto.Narrative {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *EnrollmentResponse) GetContained() []*any.Any {
	if x != nil {
		return x.Contained
	}
	return nil
}

func (x *EnrollmentResponse) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *EnrollmentResponse) GetModifierExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.ModifierExtension
	}
	return nil
}

func (x *EnrollmentResponse) GetIdentifier() []*datatypes_go_proto.Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *EnrollmentResponse) GetStatus() *EnrollmentResponse_StatusCode {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *EnrollmentResponse) GetRequest() *datatypes_go_proto.Reference {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *EnrollmentResponse) GetOutcome() *EnrollmentResponse_OutcomeCode {
	if x != nil {
		return x.Outcome
	}
	return nil
}

func (x *EnrollmentResponse) GetDisposition() *datatypes_go_proto.String {
	if x != nil {
		return x.Disposition
	}
	return nil
}

func (x *EnrollmentResponse) GetCreated() *datatypes_go_proto.DateTime {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *EnrollmentResponse) GetOrganization() *datatypes_go_proto.Reference {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *EnrollmentResponse) GetRequestProvider() *datatypes_go_proto.Reference {
	if x != nil {
		return x.RequestProvider
	}
	return nil
}

// active | cancelled | draft | entered-in-error
type EnrollmentResponse_StatusCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     codes_go_proto.FinancialResourceStatusCode_Value `protobuf:"varint,1,opt,name=value,proto3,enum=google.fhir.r4.core.FinancialResourceStatusCode_Value" json:"value,omitempty"`
	Id        *datatypes_go_proto.String                       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Extension []*datatypes_go_proto.Extension                  `protobuf:"bytes,3,rep,name=extension,proto3" json:"extension,omitempty"`
}

func (x *EnrollmentResponse_StatusCode) Reset() {
	*x = EnrollmentResponse_StatusCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollmentResponse_StatusCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollmentResponse_StatusCode) ProtoMessage() {}

func (x *EnrollmentResponse_StatusCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollmentResponse_StatusCode.ProtoReflect.Descriptor instead.
func (*EnrollmentResponse_StatusCode) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EnrollmentResponse_StatusCode) GetValue() codes_go_proto.FinancialResourceStatusCode_Value {
	if x != nil {
		return x.Value
	}
	return codes_go_proto.FinancialResourceStatusCode_INVALID_UNINITIALIZED
}

func (x *EnrollmentResponse_StatusCode) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EnrollmentResponse_StatusCode) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

// queued | complete | error | partial
type EnrollmentResponse_OutcomeCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     codes_go_proto.ClaimProcessingCode_Value `protobuf:"varint,1,opt,name=value,proto3,enum=google.fhir.r4.core.ClaimProcessingCode_Value" json:"value,omitempty"`
	Id        *datatypes_go_proto.String               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Extension []*datatypes_go_proto.Extension          `protobuf:"bytes,3,rep,name=extension,proto3" json:"extension,omitempty"`
}

func (x *EnrollmentResponse_OutcomeCode) Reset() {
	*x = EnrollmentResponse_OutcomeCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollmentResponse_OutcomeCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollmentResponse_OutcomeCode) ProtoMessage() {}

func (x *EnrollmentResponse_OutcomeCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollmentResponse_OutcomeCode.ProtoReflect.Descriptor instead.
func (*EnrollmentResponse_OutcomeCode) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescGZIP(), []int{0, 1}
}

func (x *EnrollmentResponse_OutcomeCode) GetValue() codes_go_proto.ClaimProcessingCode_Value {
	if x != nil {
		return x.Value
	}
	return codes_go_proto.ClaimProcessingCode_INVALID_UNINITIALIZED
}

func (x *EnrollmentResponse_OutcomeCode) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EnrollmentResponse_OutcomeCode) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

var File_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto protoreflect.FileDescriptor

var file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDesc = []byte{
	0x0a, 0x43, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x34, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68,
	0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x34, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x34, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c,
	0x0e, 0x0a, 0x12, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e,
	0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3f, 0x0a,
	0x0e, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x72, 0x69, 0x52,
	0x0d, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x35,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72,
	0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69,
	0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x61, 0x72, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x3c, 0x0a,
	0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72,
	0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x12, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x51, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x17, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x11,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x07, 0x6f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x56, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x12, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0c, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68,
	0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x42, 0x3a, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0c, 0x50, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x10, 0x50, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0xf2, 0xff,
	0xfc, 0xc2, 0x06, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x1a, 0xab, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x4c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2b,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x64, 0xc0, 0x9f, 0xe3, 0xb6, 0x05,
	0x01, 0x8a, 0xf9, 0x83, 0xb2, 0x05, 0x26, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68, 0x6c,
	0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x65, 0x74, 0x2f, 0x66, 0x6d, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x9a, 0xb5, 0x8e,
	0x93, 0x06, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x1a,
	0xad, 0x02, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x44, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e,
	0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x68, 0x69, 0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x6d, 0xc0, 0x9f, 0xe3, 0xb6, 0x05, 0x01, 0x8a, 0xf9, 0x83, 0xb2, 0x05, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69,
	0x72, 0x2f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x6d, 0x69, 0x74,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x9a, 0xb5, 0x8e,
	0x93, 0x06, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x3a,
	0x46, 0xc0, 0x9f, 0xe3, 0xb6, 0x05, 0x03, 0xb2, 0xfe, 0xe4, 0x97, 0x06, 0x3a, 0x68, 0x74, 0x74,
	0x70, 0x3a, 0x2f, 0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72,
	0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x42, 0x83, 0x01,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69,
	0x72, 0x2e, 0x72, 0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x60, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x34, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x98, 0xc6, 0xb0,
	0xb5, 0x07, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescOnce sync.Once
	file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescData = file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDesc
)

func file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescGZIP() []byte {
	file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescOnce.Do(func() {
		file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescData)
	})
	return file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDescData
}

var file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_goTypes = []interface{}{
	(*EnrollmentResponse)(nil),                            // 0: google.fhir.r4.core.EnrollmentResponse
	(*EnrollmentResponse_StatusCode)(nil),                 // 1: google.fhir.r4.core.EnrollmentResponse.StatusCode
	(*EnrollmentResponse_OutcomeCode)(nil),                // 2: google.fhir.r4.core.EnrollmentResponse.OutcomeCode
	(*datatypes_go_proto.Id)(nil),                         // 3: google.fhir.r4.core.Id
	(*datatypes_go_proto.Meta)(nil),                       // 4: google.fhir.r4.core.Meta
	(*datatypes_go_proto.Uri)(nil),                        // 5: google.fhir.r4.core.Uri
	(*datatypes_go_proto.Code)(nil),                       // 6: google.fhir.r4.core.Code
	(*datatypes_go_proto.Narrative)(nil),                  // 7: google.fhir.r4.core.Narrative
	(*any.Any)(nil),                                       // 8: google.protobuf.Any
	(*datatypes_go_proto.Extension)(nil),                  // 9: google.fhir.r4.core.Extension
	(*datatypes_go_proto.Identifier)(nil),                 // 10: google.fhir.r4.core.Identifier
	(*datatypes_go_proto.Reference)(nil),                  // 11: google.fhir.r4.core.Reference
	(*datatypes_go_proto.String)(nil),                     // 12: google.fhir.r4.core.String
	(*datatypes_go_proto.DateTime)(nil),                   // 13: google.fhir.r4.core.DateTime
	(codes_go_proto.FinancialResourceStatusCode_Value)(0), // 14: google.fhir.r4.core.FinancialResourceStatusCode.Value
	(codes_go_proto.ClaimProcessingCode_Value)(0),         // 15: google.fhir.r4.core.ClaimProcessingCode.Value
}
var file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_depIdxs = []int32{
	3,  // 0: google.fhir.r4.core.EnrollmentResponse.id:type_name -> google.fhir.r4.core.Id
	4,  // 1: google.fhir.r4.core.EnrollmentResponse.meta:type_name -> google.fhir.r4.core.Meta
	5,  // 2: google.fhir.r4.core.EnrollmentResponse.implicit_rules:type_name -> google.fhir.r4.core.Uri
	6,  // 3: google.fhir.r4.core.EnrollmentResponse.language:type_name -> google.fhir.r4.core.Code
	7,  // 4: google.fhir.r4.core.EnrollmentResponse.text:type_name -> google.fhir.r4.core.Narrative
	8,  // 5: google.fhir.r4.core.EnrollmentResponse.contained:type_name -> google.protobuf.Any
	9,  // 6: google.fhir.r4.core.EnrollmentResponse.extension:type_name -> google.fhir.r4.core.Extension
	9,  // 7: google.fhir.r4.core.EnrollmentResponse.modifier_extension:type_name -> google.fhir.r4.core.Extension
	10, // 8: google.fhir.r4.core.EnrollmentResponse.identifier:type_name -> google.fhir.r4.core.Identifier
	1,  // 9: google.fhir.r4.core.EnrollmentResponse.status:type_name -> google.fhir.r4.core.EnrollmentResponse.StatusCode
	11, // 10: google.fhir.r4.core.EnrollmentResponse.request:type_name -> google.fhir.r4.core.Reference
	2,  // 11: google.fhir.r4.core.EnrollmentResponse.outcome:type_name -> google.fhir.r4.core.EnrollmentResponse.OutcomeCode
	12, // 12: google.fhir.r4.core.EnrollmentResponse.disposition:type_name -> google.fhir.r4.core.String
	13, // 13: google.fhir.r4.core.EnrollmentResponse.created:type_name -> google.fhir.r4.core.DateTime
	11, // 14: google.fhir.r4.core.EnrollmentResponse.organization:type_name -> google.fhir.r4.core.Reference
	11, // 15: google.fhir.r4.core.EnrollmentResponse.request_provider:type_name -> google.fhir.r4.core.Reference
	14, // 16: google.fhir.r4.core.EnrollmentResponse.StatusCode.value:type_name -> google.fhir.r4.core.FinancialResourceStatusCode.Value
	12, // 17: google.fhir.r4.core.EnrollmentResponse.StatusCode.id:type_name -> google.fhir.r4.core.String
	9,  // 18: google.fhir.r4.core.EnrollmentResponse.StatusCode.extension:type_name -> google.fhir.r4.core.Extension
	15, // 19: google.fhir.r4.core.EnrollmentResponse.OutcomeCode.value:type_name -> google.fhir.r4.core.ClaimProcessingCode.Value
	12, // 20: google.fhir.r4.core.EnrollmentResponse.OutcomeCode.id:type_name -> google.fhir.r4.core.String
	9,  // 21: google.fhir.r4.core.EnrollmentResponse.OutcomeCode.extension:type_name -> google.fhir.r4.core.Extension
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_init() }
func file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_init() {
	if File_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollmentResponse); i {
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
		file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollmentResponse_StatusCode); i {
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
		file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollmentResponse_OutcomeCode); i {
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
			RawDescriptor: file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_goTypes,
		DependencyIndexes: file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_depIdxs,
		MessageInfos:      file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_msgTypes,
	}.Build()
	File_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto = out.File
	file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_rawDesc = nil
	file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_goTypes = nil
	file_proto_google_fhir_proto_r4_core_resources_enrollment_response_proto_depIdxs = nil
}
