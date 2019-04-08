// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature            string   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Timestamp            uint32   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MessageRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *MessageRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MessageRequest) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type MessageResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "api.MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "api.MessageResponse")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0x87, 0x30,
	0x10, 0xc5, 0x53, 0xff, 0x44, 0xe5, 0x0c, 0x98, 0x9c, 0x0e, 0x0d, 0x71, 0x20, 0x24, 0x26, 0x4c,
	0x0c, 0x3a, 0x3a, 0xba, 0x99, 0xb8, 0xd4, 0x4f, 0x50, 0xf1, 0x82, 0x1d, 0x4a, 0x6b, 0xaf, 0xf0,
	0x01, 0xfc, 0xe4, 0xc6, 0xc2, 0x1f, 0xc2, 0xd6, 0xf7, 0x7b, 0x7d, 0x97, 0x77, 0x07, 0x85, 0x25,
	0x66, 0x3d, 0x50, 0xe7, 0x83, 0x8b, 0x0e, 0x4f, 0xda, 0x9b, 0xe6, 0x57, 0x40, 0xf9, 0xbe, 0x60,
	0x45, 0x3f, 0x13, 0x71, 0x44, 0x09, 0x57, 0xeb, 0x47, 0x29, 0x6a, 0xd1, 0xe6, 0xea, 0x2c, 0xf1,
	0x01, 0x72, 0x36, 0xc3, 0xa8, 0xe3, 0x14, 0x48, 0x5e, 0x24, 0x6f, 0x07, 0x58, 0xc1, 0xf5, 0xc4,
	0x14, 0x46, 0x6d, 0x49, 0x9e, 0x92, 0xb9, 0xe9, 0xff, 0x64, 0x34, 0x96, 0x38, 0x6a, 0xeb, 0x65,
	0x56, 0x8b, 0xb6, 0x50, 0x3b, 0x68, 0x1e, 0xe1, 0x76, 0xeb, 0xc0, 0xde, 0x8d, 0x4c, 0x88, 0x90,
	0xf5, 0xee, 0x6b, 0x69, 0x50, 0xa8, 0xf4, 0x7e, 0x7a, 0x83, 0x9b, 0xd7, 0x6f, 0x1d, 0x3f, 0x28,
	0xcc, 0xa6, 0x27, 0x7c, 0x81, 0x52, 0x51, 0x4f, 0x66, 0xa6, 0x35, 0x8c, 0x77, 0x9d, 0xf6, 0xa6,
	0x3b, 0xae, 0x53, 0xdd, 0x1f, 0xe1, 0x32, 0xff, 0xf3, 0x32, 0xdd, 0xe0, 0xf9, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0x88, 0x03, 0x7f, 0x14, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	ReceiveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ReceiveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/api.ChatService/ReceiveMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	ReceiveMessage(context.Context, *MessageRequest) (*MessageResponse, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) ReceiveMessage(ctx context.Context, req *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_ReceiveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ReceiveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChatService/ReceiveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ReceiveMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveMessage",
			Handler:    _ChatService_ReceiveMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}