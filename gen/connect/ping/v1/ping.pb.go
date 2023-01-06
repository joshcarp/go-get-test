// Copyright 2021-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: connect/ping/v1/ping.proto

package pingv1_test

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *PingRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *PingResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type FailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *FailRequest) Reset() {
	*x = FailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailRequest) ProtoMessage() {}

func (x *FailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailRequest.ProtoReflect.Descriptor instead.
func (*FailRequest) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{2}
}

func (x *FailRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type FailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FailResponse) Reset() {
	*x = FailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailResponse) ProtoMessage() {}

func (x *FailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailResponse.ProtoReflect.Descriptor instead.
func (*FailResponse) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{3}
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{4}
}

func (x *SumRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{5}
}

func (x *SumResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type CountUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CountUpRequest) Reset() {
	*x = CountUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUpRequest) ProtoMessage() {}

func (x *CountUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUpRequest.ProtoReflect.Descriptor instead.
func (*CountUpRequest) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{6}
}

func (x *CountUpRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CountUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CountUpResponse) Reset() {
	*x = CountUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUpResponse) ProtoMessage() {}

func (x *CountUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUpResponse.ProtoReflect.Descriptor instead.
func (*CountUpResponse) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{7}
}

func (x *CountUpResponse) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CumSumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CumSumRequest) Reset() {
	*x = CumSumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CumSumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CumSumRequest) ProtoMessage() {}

func (x *CumSumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CumSumRequest.ProtoReflect.Descriptor instead.
func (*CumSumRequest) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{8}
}

func (x *CumSumRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CumSumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *CumSumResponse) Reset() {
	*x = CumSumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ping_v1_ping_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CumSumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CumSumResponse) ProtoMessage() {}

func (x *CumSumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ping_v1_ping_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CumSumResponse.ProtoReflect.Descriptor instead.
func (*CumSumResponse) Descriptor() ([]byte, []int) {
	return file_connect_ping_v1_ping_proto_rawDescGZIP(), []int{9}
}

func (x *CumSumResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_connect_ping_v1_ping_proto protoreflect.FileDescriptor

var file_connect_ping_v1_ping_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x39, 0x0a,
	0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x46, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x1f, 0x0a,
	0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x28,
	0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x75, 0x6d, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0e,
	0x43, 0x75, 0x6d, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d,
	0x32, 0x84, 0x03, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x45, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x12,
	0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x50, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x12,
	0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x06, 0x43, 0x75, 0x6d, 0x53, 0x75, 0x6d,
	0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x75, 0x6d, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x75, 0x6d, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0xba, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x73, 0x68, 0x63, 0x61, 0x72,
	0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x74, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x58, 0xaa, 0x02,
	0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x69, 0x6e, 0x67,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connect_ping_v1_ping_proto_rawDescOnce sync.Once
	file_connect_ping_v1_ping_proto_rawDescData = file_connect_ping_v1_ping_proto_rawDesc
)

func file_connect_ping_v1_ping_proto_rawDescGZIP() []byte {
	file_connect_ping_v1_ping_proto_rawDescOnce.Do(func() {
		file_connect_ping_v1_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_connect_ping_v1_ping_proto_rawDescData)
	})
	return file_connect_ping_v1_ping_proto_rawDescData
}

var file_connect_ping_v1_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_connect_ping_v1_ping_proto_goTypes = []interface{}{
	(*PingRequest)(nil),     // 0: connect.ping.v1.PingRequest
	(*PingResponse)(nil),    // 1: connect.ping.v1.PingResponse
	(*FailRequest)(nil),     // 2: connect.ping.v1.FailRequest
	(*FailResponse)(nil),    // 3: connect.ping.v1.FailResponse
	(*SumRequest)(nil),      // 4: connect.ping.v1.SumRequest
	(*SumResponse)(nil),     // 5: connect.ping.v1.SumResponse
	(*CountUpRequest)(nil),  // 6: connect.ping.v1.CountUpRequest
	(*CountUpResponse)(nil), // 7: connect.ping.v1.CountUpResponse
	(*CumSumRequest)(nil),   // 8: connect.ping.v1.CumSumRequest
	(*CumSumResponse)(nil),  // 9: connect.ping.v1.CumSumResponse
}
var file_connect_ping_v1_ping_proto_depIdxs = []int32{
	0, // 0: connect.ping.v1.PingService.Ping:input_type -> connect.ping.v1.PingRequest
	2, // 1: connect.ping.v1.PingService.Fail:input_type -> connect.ping.v1.FailRequest
	4, // 2: connect.ping.v1.PingService.Sum:input_type -> connect.ping.v1.SumRequest
	6, // 3: connect.ping.v1.PingService.CountUp:input_type -> connect.ping.v1.CountUpRequest
	8, // 4: connect.ping.v1.PingService.CumSum:input_type -> connect.ping.v1.CumSumRequest
	1, // 5: connect.ping.v1.PingService.Ping:output_type -> connect.ping.v1.PingResponse
	3, // 6: connect.ping.v1.PingService.Fail:output_type -> connect.ping.v1.FailResponse
	5, // 7: connect.ping.v1.PingService.Sum:output_type -> connect.ping.v1.SumResponse
	7, // 8: connect.ping.v1.PingService.CountUp:output_type -> connect.ping.v1.CountUpResponse
	9, // 9: connect.ping.v1.PingService.CumSum:output_type -> connect.ping.v1.CumSumResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_connect_ping_v1_ping_proto_init() }
func file_connect_ping_v1_ping_proto_init() {
	if File_connect_ping_v1_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connect_ping_v1_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailRequest); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailResponse); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUpRequest); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUpResponse); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CumSumRequest); i {
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
		file_connect_ping_v1_ping_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CumSumResponse); i {
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
			RawDescriptor: file_connect_ping_v1_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connect_ping_v1_ping_proto_goTypes,
		DependencyIndexes: file_connect_ping_v1_ping_proto_depIdxs,
		MessageInfos:      file_connect_ping_v1_ping_proto_msgTypes,
	}.Build()
	File_connect_ping_v1_ping_proto = out.File
	file_connect_ping_v1_ping_proto_rawDesc = nil
	file_connect_ping_v1_ping_proto_goTypes = nil
	file_connect_ping_v1_ping_proto_depIdxs = nil
}
