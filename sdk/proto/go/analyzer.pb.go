// Code generated by protoc-gen-go.
// source: analyzer.proto
// DO NOT EDIT!

/*
Package pulumirpc is a generated protocol buffer package.

It is generated from these files:
	analyzer.proto
	engine.proto
	languages.proto
	provider.proto

It has these top-level messages:
	AnalyzeRequest
	AnalyzeResponse
	AnalyzeFailure
	LogRequest
	RunRequest
	RunResponse
	NewResourceRequest
	NewResourceResponse
	ConfigureRequest
	CheckRequest
	CheckResponse
	CheckFailure
	DiffRequest
	DiffResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
*/
package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

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

type AnalyzeRequest struct {
	Type       string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *AnalyzeRequest) Reset()                    { *m = AnalyzeRequest{} }
func (m *AnalyzeRequest) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeRequest) ProtoMessage()               {}
func (*AnalyzeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AnalyzeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnalyzeRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type AnalyzeResponse struct {
	Failures []*AnalyzeFailure `protobuf:"bytes,1,rep,name=failures" json:"failures,omitempty"`
}

func (m *AnalyzeResponse) Reset()                    { *m = AnalyzeResponse{} }
func (m *AnalyzeResponse) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeResponse) ProtoMessage()               {}
func (*AnalyzeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AnalyzeResponse) GetFailures() []*AnalyzeFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type AnalyzeFailure struct {
	Property string `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Reason   string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *AnalyzeFailure) Reset()                    { *m = AnalyzeFailure{} }
func (m *AnalyzeFailure) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeFailure) ProtoMessage()               {}
func (*AnalyzeFailure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AnalyzeFailure) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *AnalyzeFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*AnalyzeRequest)(nil), "pulumirpc.AnalyzeRequest")
	proto.RegisterType((*AnalyzeResponse)(nil), "pulumirpc.AnalyzeResponse")
	proto.RegisterType((*AnalyzeFailure)(nil), "pulumirpc.AnalyzeFailure")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Analyzer service

type AnalyzerClient interface {
	// Analyze analyzes a single resource object, and returns any errors that it finds.
	Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
}

type analyzerClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzerClient(cc *grpc.ClientConn) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.Analyzer/Analyze", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analyzer service

type AnalyzerServer interface {
	// Analyze analyzes a single resource object, and returns any errors that it finds.
	Analyze(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error)
}

func RegisterAnalyzerServer(s *grpc.Server, srv AnalyzerServer) {
	s.RegisterService(&_Analyzer_serviceDesc, srv)
}

func _Analyzer_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Analyzer/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Analyze(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _Analyzer_Analyze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyzer.proto",
}

func init() { proto.RegisterFile("analyzer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0xad, 0xca, 0xd9, 0xbe, 0x83, 0x13, 0xde, 0xa0, 0xb5, 0x38, 0x94, 0x4e, 0x9d, 0x72,
	0x50, 0x11, 0x67, 0x45, 0xc4, 0xc9, 0x21, 0xce, 0x0e, 0xbd, 0xe3, 0xdd, 0x51, 0xa8, 0x4d, 0x7c,
	0x49, 0x86, 0xfa, 0xd7, 0x0b, 0x49, 0x0c, 0x37, 0x74, 0x7b, 0x8f, 0xef, 0xeb, 0xaf, 0xbf, 0x17,
	0xd8, 0xf4, 0x53, 0x3f, 0xce, 0xbf, 0xc4, 0x42, 0xb3, 0xb2, 0x0a, 0x0b, 0xed, 0x46, 0xf7, 0x3d,
	0xb0, 0xde, 0x57, 0xf7, 0x47, 0xa5, 0x8e, 0x23, 0x6d, 0x7d, 0xb0, 0x73, 0x87, 0xad, 0xb1, 0xec,
	0xf6, 0x36, 0x14, 0x9b, 0x2f, 0xd8, 0x3c, 0x87, 0x4f, 0x25, 0xfd, 0x38, 0x32, 0x16, 0x11, 0x2e,
	0xed, 0xac, 0xa9, 0xcc, 0xea, 0xac, 0x2d, 0xa4, 0x9f, 0xf1, 0x09, 0x40, 0xb3, 0xd2, 0xc4, 0x76,
	0x20, 0x53, 0x9e, 0xd7, 0x59, 0xbb, 0xee, 0x6e, 0x45, 0x00, 0x8b, 0x7f, 0xb0, 0xf8, 0xf4, 0x60,
	0x79, 0x52, 0x6d, 0xde, 0xe1, 0x3a, 0xe1, 0x8d, 0x56, 0x93, 0x21, 0x7c, 0x84, 0xfc, 0xd0, 0x0f,
	0xa3, 0x63, 0x32, 0x65, 0x56, 0x5f, 0xb4, 0xeb, 0xee, 0x4e, 0x24, 0x5b, 0x11, 0xdb, 0x6f, 0xa1,
	0x21, 0x53, 0xb5, 0x79, 0x4d, 0xa2, 0x31, 0xc3, 0x0a, 0xf2, 0xf8, 0xa7, 0x39, 0xca, 0xa6, 0x1d,
	0x6f, 0x60, 0xc5, 0xd4, 0x1b, 0x35, 0x79, 0xd9, 0x42, 0xc6, 0xad, 0xfb, 0x80, 0x3c, 0x52, 0x18,
	0x5f, 0xe0, 0x2a, 0xce, 0xb8, 0x60, 0x10, 0x9f, 0xa3, 0xaa, 0x96, 0xa2, 0x70, 0x4a, 0x73, 0xb6,
	0x5b, 0xf9, 0xe3, 0x1f, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x30, 0x0e, 0x75, 0x80, 0x01,
	0x00, 0x00,
}