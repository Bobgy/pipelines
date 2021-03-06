// Copyright 2019 Google LLC
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
// source: backend/api/visualization.proto

package go_client // import "github.com/kubeflow/pipelines/backend/api/go_client"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Visualization_Type int32

const (
	Visualization_ROC_CURVE Visualization_Type = 0
)

var Visualization_Type_name = map[int32]string{
	0: "ROC_CURVE",
}
var Visualization_Type_value = map[string]int32{
	"ROC_CURVE": 0,
}

func (x Visualization_Type) String() string {
	return proto.EnumName(Visualization_Type_name, int32(x))
}
func (Visualization_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_visualization_86dbab7fab4100b0, []int{1, 0}
}

type CreateVisualizationRequest struct {
	Visualization        *Visualization `protobuf:"bytes,1,opt,name=visualization,proto3" json:"visualization,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateVisualizationRequest) Reset()         { *m = CreateVisualizationRequest{} }
func (m *CreateVisualizationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVisualizationRequest) ProtoMessage()    {}
func (*CreateVisualizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_visualization_86dbab7fab4100b0, []int{0}
}
func (m *CreateVisualizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVisualizationRequest.Unmarshal(m, b)
}
func (m *CreateVisualizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVisualizationRequest.Marshal(b, m, deterministic)
}
func (dst *CreateVisualizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVisualizationRequest.Merge(dst, src)
}
func (m *CreateVisualizationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVisualizationRequest.Size(m)
}
func (m *CreateVisualizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVisualizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVisualizationRequest proto.InternalMessageInfo

func (m *CreateVisualizationRequest) GetVisualization() *Visualization {
	if m != nil {
		return m.Visualization
	}
	return nil
}

type Visualization struct {
	Type                 Visualization_Type `protobuf:"varint,1,opt,name=type,proto3,enum=api.Visualization_Type" json:"type,omitempty"`
	Source               string             `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Arguments            string             `protobuf:"bytes,3,opt,name=arguments,proto3" json:"arguments,omitempty"`
	Html                 string             `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`
	Error                string             `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Visualization) Reset()         { *m = Visualization{} }
func (m *Visualization) String() string { return proto.CompactTextString(m) }
func (*Visualization) ProtoMessage()    {}
func (*Visualization) Descriptor() ([]byte, []int) {
	return fileDescriptor_visualization_86dbab7fab4100b0, []int{1}
}
func (m *Visualization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Visualization.Unmarshal(m, b)
}
func (m *Visualization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Visualization.Marshal(b, m, deterministic)
}
func (dst *Visualization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Visualization.Merge(dst, src)
}
func (m *Visualization) XXX_Size() int {
	return xxx_messageInfo_Visualization.Size(m)
}
func (m *Visualization) XXX_DiscardUnknown() {
	xxx_messageInfo_Visualization.DiscardUnknown(m)
}

var xxx_messageInfo_Visualization proto.InternalMessageInfo

func (m *Visualization) GetType() Visualization_Type {
	if m != nil {
		return m.Type
	}
	return Visualization_ROC_CURVE
}

func (m *Visualization) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Visualization) GetArguments() string {
	if m != nil {
		return m.Arguments
	}
	return ""
}

func (m *Visualization) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *Visualization) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateVisualizationRequest)(nil), "api.CreateVisualizationRequest")
	proto.RegisterType((*Visualization)(nil), "api.Visualization")
	proto.RegisterEnum("api.Visualization_Type", Visualization_Type_name, Visualization_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VisualizationServiceClient is the client API for VisualizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VisualizationServiceClient interface {
	CreateVisualization(ctx context.Context, in *CreateVisualizationRequest, opts ...grpc.CallOption) (*Visualization, error)
}

type visualizationServiceClient struct {
	cc *grpc.ClientConn
}

func NewVisualizationServiceClient(cc *grpc.ClientConn) VisualizationServiceClient {
	return &visualizationServiceClient{cc}
}

func (c *visualizationServiceClient) CreateVisualization(ctx context.Context, in *CreateVisualizationRequest, opts ...grpc.CallOption) (*Visualization, error) {
	out := new(Visualization)
	err := c.cc.Invoke(ctx, "/api.VisualizationService/CreateVisualization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisualizationServiceServer is the server API for VisualizationService service.
type VisualizationServiceServer interface {
	CreateVisualization(context.Context, *CreateVisualizationRequest) (*Visualization, error)
}

func RegisterVisualizationServiceServer(s *grpc.Server, srv VisualizationServiceServer) {
	s.RegisterService(&_VisualizationService_serviceDesc, srv)
}

func _VisualizationService_CreateVisualization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVisualizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisualizationServiceServer).CreateVisualization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VisualizationService/CreateVisualization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisualizationServiceServer).CreateVisualization(ctx, req.(*CreateVisualizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VisualizationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.VisualizationService",
	HandlerType: (*VisualizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVisualization",
			Handler:    _VisualizationService_CreateVisualization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/api/visualization.proto",
}

func init() {
	proto.RegisterFile("backend/api/visualization.proto", fileDescriptor_visualization_86dbab7fab4100b0)
}

var fileDescriptor_visualization_86dbab7fab4100b0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x6b, 0xd7, 0x0d, 0x64, 0x8a, 0xf9, 0xd8, 0x16, 0x6a, 0x59, 0x46, 0x89, 0x7c, 0xaa,
	0x04, 0xb1, 0xd5, 0xe4, 0x82, 0xb8, 0xd1, 0x88, 0x23, 0x42, 0x72, 0x21, 0x87, 0x5e, 0xaa, 0xb5,
	0x3b, 0xdd, 0xac, 0xea, 0xec, 0x2e, 0xfb, 0x91, 0xaa, 0x5c, 0x90, 0x90, 0x78, 0x01, 0x38, 0xf0,
	0x20, 0x3c, 0x0a, 0xaf, 0xc0, 0x83, 0xa0, 0x6e, 0x23, 0x88, 0xd5, 0xf6, 0x64, 0xcf, 0xcc, 0x6f,
	0xfe, 0x3b, 0xfb, 0xdf, 0x81, 0x41, 0x4d, 0x9b, 0x73, 0x14, 0xa7, 0x25, 0x55, 0xbc, 0x5c, 0x72,
	0xe3, 0x68, 0xcb, 0x3f, 0x53, 0xcb, 0xa5, 0x28, 0x94, 0x96, 0x56, 0x92, 0x4d, 0xaa, 0x78, 0x9a,
	0x31, 0x29, 0x59, 0x8b, 0x1e, 0xa2, 0x42, 0x48, 0xeb, 0x09, 0x73, 0x8d, 0xa4, 0x7b, 0xeb, 0x1a,
	0xa8, 0xb5, 0xd4, 0xab, 0xc2, 0x4b, 0xff, 0x69, 0x46, 0x0c, 0xc5, 0xc8, 0x5c, 0x50, 0xc6, 0x50,
	0x97, 0x52, 0xf9, 0xd6, 0x9b, 0x32, 0xf9, 0x0c, 0xd2, 0xa9, 0x46, 0x6a, 0x71, 0xb6, 0x3e, 0x46,
	0x85, 0x9f, 0x1c, 0x1a, 0x4b, 0x5e, 0x41, 0xdc, 0x19, 0x2f, 0x09, 0x86, 0xc1, 0xfe, 0xf6, 0x98,
	0x14, 0x54, 0xf1, 0xa2, 0xdb, 0xd1, 0x05, 0xf3, 0x5f, 0x01, 0xc4, 0x1d, 0x80, 0xbc, 0x80, 0xc8,
	0x5e, 0x2a, 0xf4, 0x12, 0x0f, 0xc7, 0x7b, 0x37, 0x25, 0x8a, 0x0f, 0x97, 0x0a, 0x2b, 0x0f, 0x91,
	0x67, 0xd0, 0x33, 0xd2, 0xe9, 0x06, 0x93, 0x70, 0x18, 0xec, 0xf7, 0xab, 0x55, 0x44, 0x32, 0xe8,
	0x53, 0xcd, 0xdc, 0x02, 0x85, 0x35, 0xc9, 0xa6, 0x2f, 0xfd, 0x4f, 0x10, 0x02, 0xd1, 0xdc, 0x2e,
	0xda, 0x24, 0xf2, 0x05, 0xff, 0x4f, 0x76, 0x61, 0xcb, 0xbb, 0x93, 0x6c, 0xf9, 0xe4, 0x75, 0x90,
	0x3f, 0x85, 0xe8, 0xea, 0x34, 0x12, 0x43, 0xbf, 0x7a, 0x3f, 0x3d, 0x99, 0x7e, 0xac, 0x66, 0x6f,
	0x1f, 0x6f, 0x8c, 0x7f, 0x06, 0xb0, 0xdb, 0x99, 0xe9, 0x08, 0xf5, 0x92, 0x37, 0x48, 0xbe, 0xc0,
	0xce, 0x2d, 0x36, 0x91, 0x81, 0xbf, 0xc5, 0xdd, 0x06, 0xa6, 0xb7, 0x38, 0x95, 0x4f, 0xbe, 0xfe,
	0xfe, 0xf3, 0x23, 0x1c, 0xe5, 0xd9, 0xd5, 0xd3, 0x99, 0x72, 0x79, 0x50, 0xa3, 0xa5, 0x07, 0xdd,
	0x3d, 0x30, 0xaf, 0xbb, 0x7e, 0x1e, 0x7e, 0x0b, 0xbe, 0xbf, 0x79, 0x57, 0x65, 0x70, 0xef, 0x14,
	0xcf, 0xa8, 0x6b, 0x2d, 0x79, 0x42, 0x1e, 0x41, 0x9c, 0x6e, 0x7b, 0xfd, 0x23, 0x4b, 0xad, 0x33,
	0xc7, 0x03, 0x78, 0x0e, 0xbd, 0x43, 0xa4, 0x1a, 0x35, 0xd9, 0xb9, 0x1f, 0xa6, 0x31, 0x75, 0x76,
	0x2e, 0xf5, 0x4a, 0x62, 0x18, 0xd6, 0x0f, 0x00, 0xfe, 0x01, 0x1b, 0xc7, 0x13, 0xc6, 0xed, 0xdc,
	0xd5, 0x45, 0x23, 0x17, 0xe5, 0xb9, 0xab, 0xf1, 0xac, 0x95, 0x17, 0xa5, 0xe2, 0x0a, 0x5b, 0x2e,
	0xd0, 0x94, 0xeb, 0xfb, 0xc5, 0xe4, 0x49, 0xd3, 0x72, 0x14, 0xb6, 0xee, 0xf9, 0xb5, 0x99, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xb5, 0x35, 0x55, 0xc3, 0x02, 0x00, 0x00,
}
