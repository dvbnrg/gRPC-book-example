// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tokenizer.proto

package practical_grpc_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TokenizeRequest struct {
	FileContents []byte `protobuf:"bytes,1,opt,name=file_contents,json=fileContents,proto3" json:"file_contents,omitempty"`
}

func (m *TokenizeRequest) Reset()                    { *m = TokenizeRequest{} }
func (m *TokenizeRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenizeRequest) ProtoMessage()               {}
func (*TokenizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *TokenizeRequest) GetFileContents() []byte {
	if m != nil {
		return m.FileContents
	}
	return nil
}

type TokenizeResponse struct {
	Words map[string]int64 `protobuf:"bytes,1,rep,name=words" json:"words,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *TokenizeResponse) Reset()                    { *m = TokenizeResponse{} }
func (m *TokenizeResponse) String() string            { return proto.CompactTextString(m) }
func (*TokenizeResponse) ProtoMessage()               {}
func (*TokenizeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *TokenizeResponse) GetWords() map[string]int64 {
	if m != nil {
		return m.Words
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenizeRequest)(nil), "practical.grpc.v1.TokenizeRequest")
	proto.RegisterType((*TokenizeResponse)(nil), "practical.grpc.v1.TokenizeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tokenizer service

type TokenizerClient interface {
	Tokenize(ctx context.Context, opts ...grpc.CallOption) (Tokenizer_TokenizeClient, error)
}

type tokenizerClient struct {
	cc *grpc.ClientConn
}

func NewTokenizerClient(cc *grpc.ClientConn) TokenizerClient {
	return &tokenizerClient{cc}
}

func (c *tokenizerClient) Tokenize(ctx context.Context, opts ...grpc.CallOption) (Tokenizer_TokenizeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Tokenizer_serviceDesc.Streams[0], c.cc, "/practical.grpc.v1.Tokenizer/Tokenize", opts...)
	if err != nil {
		return nil, err
	}
	x := &tokenizerTokenizeClient{stream}
	return x, nil
}

type Tokenizer_TokenizeClient interface {
	Send(*TokenizeRequest) error
	Recv() (*TokenizeResponse, error)
	grpc.ClientStream
}

type tokenizerTokenizeClient struct {
	grpc.ClientStream
}

func (x *tokenizerTokenizeClient) Send(m *TokenizeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tokenizerTokenizeClient) Recv() (*TokenizeResponse, error) {
	m := new(TokenizeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Tokenizer service

type TokenizerServer interface {
	Tokenize(Tokenizer_TokenizeServer) error
}

func RegisterTokenizerServer(s *grpc.Server, srv TokenizerServer) {
	s.RegisterService(&_Tokenizer_serviceDesc, srv)
}

func _Tokenizer_Tokenize_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TokenizerServer).Tokenize(&tokenizerTokenizeServer{stream})
}

type Tokenizer_TokenizeServer interface {
	Send(*TokenizeResponse) error
	Recv() (*TokenizeRequest, error)
	grpc.ServerStream
}

type tokenizerTokenizeServer struct {
	grpc.ServerStream
}

func (x *tokenizerTokenizeServer) Send(m *TokenizeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tokenizerTokenizeServer) Recv() (*TokenizeRequest, error) {
	m := new(TokenizeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Tokenizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "practical.grpc.v1.Tokenizer",
	HandlerType: (*TokenizerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tokenize",
			Handler:       _Tokenizer_Tokenize_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tokenizer.proto",
}

func init() { proto.RegisterFile("tokenizer.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc9, 0xcf, 0x4e,
	0xcd, 0xcb, 0xac, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x28, 0x4a,
	0x4c, 0x2e, 0xc9, 0x4c, 0x4e, 0xcc, 0xd1, 0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x33, 0x54, 0x32,
	0xe3, 0xe2, 0x0f, 0x81, 0xaa, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe6, 0xe2,
	0x4d, 0xcb, 0xcc, 0x49, 0x8d, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x29, 0x96, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x09, 0xe2, 0x01, 0x09, 0x3a, 0x43, 0xc5, 0x94, 0x26, 0x31, 0x72, 0x09, 0x20,
	0x34, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xb9, 0x70, 0xb1, 0x96, 0xe7, 0x17, 0xa5, 0x80,
	0x74, 0x30, 0x6b, 0x70, 0x1b, 0xe9, 0xe9, 0x61, 0xd8, 0xa7, 0x87, 0xae, 0x47, 0x2f, 0x1c, 0xa4,
	0xc1, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x08, 0xa2, 0x59, 0xca, 0x82, 0x8b, 0x0b, 0x21, 0x28, 0x24,
	0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0x76, 0x03, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a,
	0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0xe1, 0x58, 0x31, 0x59,
	0x30, 0x1a, 0xa5, 0x70, 0x71, 0xc2, 0xcc, 0x2f, 0x12, 0x0a, 0xe7, 0xe2, 0x80, 0x71, 0x84, 0x94,
	0xf0, 0xba, 0x04, 0xec, 0x6d, 0x29, 0x65, 0x22, 0x5c, 0xab, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06,
	0x0e, 0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x33, 0x2f, 0xa6, 0x5f, 0x01, 0x00,
	0x00,
}