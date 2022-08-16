// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insecure/attacker.proto

package insecure

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Protocol int32

const (
	Protocol_UNKNOWN   Protocol = 0
	Protocol_UNICAST   Protocol = 1
	Protocol_MULTICAST Protocol = 2
	Protocol_PUBLISH   Protocol = 3
)

var Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "UNICAST",
	2: "MULTICAST",
	3: "PUBLISH",
}

var Protocol_value = map[string]int32{
	"UNKNOWN":   0,
	"UNICAST":   1,
	"MULTICAST": 2,
	"PUBLISH":   3,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}

func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de1530a0da5eef1a, []int{0}
}

// Message represents the messages exchanged between the CorruptibleNetwork (server) and Orchestrator (client).
// This is a wrapper for both egress and ingress messages.
type Message struct {
	Egress               *EgressMessage  `protobuf:"bytes,1,opt,name=Egress,proto3" json:"Egress,omitempty"`
	Ingress              *IngressMessage `protobuf:"bytes,2,opt,name=Ingress,proto3" json:"Ingress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_de1530a0da5eef1a, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetEgress() *EgressMessage {
	if m != nil {
		return m.Egress
	}
	return nil
}

func (m *Message) GetIngress() *IngressMessage {
	if m != nil {
		return m.Ingress
	}
	return nil
}

// EgressMessage represents the message exchanged between the CorruptibleConduitFactory and Attacker services.
// CorruptOriginID for EgressMessage represents the corrupt node id where the message is coming from.
type EgressMessage struct {
	ChannelID            string   `protobuf:"bytes,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	CorruptOriginID      []byte   `protobuf:"bytes,2,opt,name=CorruptOriginID,proto3" json:"CorruptOriginID,omitempty"`
	TargetNum            uint32   `protobuf:"varint,3,opt,name=TargetNum,proto3" json:"TargetNum,omitempty"`
	TargetIDs            [][]byte `protobuf:"bytes,4,rep,name=TargetIDs,proto3" json:"TargetIDs,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Protocol             Protocol `protobuf:"varint,6,opt,name=protocol,proto3,enum=corruptible.Protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EgressMessage) Reset()         { *m = EgressMessage{} }
func (m *EgressMessage) String() string { return proto.CompactTextString(m) }
func (*EgressMessage) ProtoMessage()    {}
func (*EgressMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_de1530a0da5eef1a, []int{1}
}

func (m *EgressMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EgressMessage.Unmarshal(m, b)
}
func (m *EgressMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EgressMessage.Marshal(b, m, deterministic)
}
func (m *EgressMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EgressMessage.Merge(m, src)
}
func (m *EgressMessage) XXX_Size() int {
	return xxx_messageInfo_EgressMessage.Size(m)
}
func (m *EgressMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EgressMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EgressMessage proto.InternalMessageInfo

func (m *EgressMessage) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *EgressMessage) GetCorruptOriginID() []byte {
	if m != nil {
		return m.CorruptOriginID
	}
	return nil
}

func (m *EgressMessage) GetTargetNum() uint32 {
	if m != nil {
		return m.TargetNum
	}
	return 0
}

func (m *EgressMessage) GetTargetIDs() [][]byte {
	if m != nil {
		return m.TargetIDs
	}
	return nil
}

func (m *EgressMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *EgressMessage) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_UNKNOWN
}

// OriginID for IngressMessage represents the node id where the message is coming from - that node could be corrupt or honest.
type IngressMessage struct {
	ChannelID            string   `protobuf:"bytes,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	OriginID             []byte   `protobuf:"bytes,2,opt,name=OriginID,proto3" json:"OriginID,omitempty"`
	CorruptTargetID      []byte   `protobuf:"bytes,3,opt,name=CorruptTargetID,proto3" json:"CorruptTargetID,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngressMessage) Reset()         { *m = IngressMessage{} }
func (m *IngressMessage) String() string { return proto.CompactTextString(m) }
func (*IngressMessage) ProtoMessage()    {}
func (*IngressMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_de1530a0da5eef1a, []int{2}
}

func (m *IngressMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngressMessage.Unmarshal(m, b)
}
func (m *IngressMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngressMessage.Marshal(b, m, deterministic)
}
func (m *IngressMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressMessage.Merge(m, src)
}
func (m *IngressMessage) XXX_Size() int {
	return xxx_messageInfo_IngressMessage.Size(m)
}
func (m *IngressMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressMessage.DiscardUnknown(m)
}

var xxx_messageInfo_IngressMessage proto.InternalMessageInfo

func (m *IngressMessage) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *IngressMessage) GetOriginID() []byte {
	if m != nil {
		return m.OriginID
	}
	return nil
}

func (m *IngressMessage) GetCorruptTargetID() []byte {
	if m != nil {
		return m.CorruptTargetID
	}
	return nil
}

func (m *IngressMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("corruptible.Protocol", Protocol_name, Protocol_value)
	proto.RegisterType((*Message)(nil), "corruptible.Message")
	proto.RegisterType((*EgressMessage)(nil), "corruptible.EgressMessage")
	proto.RegisterType((*IngressMessage)(nil), "corruptible.IngressMessage")
}

func init() { proto.RegisterFile("insecure/attacker.proto", fileDescriptor_de1530a0da5eef1a) }

var fileDescriptor_de1530a0da5eef1a = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x4d, 0xc9, 0x9f, 0x49, 0x53, 0xa2, 0x15, 0x14, 0xe3, 0x72, 0x88, 0x7c, 0xb2,
	0x38, 0x38, 0x60, 0xc4, 0x9d, 0xd6, 0x29, 0xc2, 0xd0, 0xba, 0x96, 0x9b, 0x08, 0x89, 0xdb, 0xc6,
	0x5d, 0x8c, 0x85, 0xbb, 0x1b, 0xed, 0xae, 0x0f, 0x79, 0x09, 0x9e, 0x84, 0x87, 0xe2, 0x51, 0x90,
	0xd7, 0xde, 0x3a, 0xae, 0x40, 0x3d, 0xce, 0x7c, 0x33, 0xdf, 0xce, 0x6f, 0x34, 0x0b, 0x2f, 0x72,
	0x26, 0x69, 0x5a, 0x0a, 0xba, 0x20, 0x4a, 0x91, 0xf4, 0x27, 0x15, 0xde, 0x56, 0x70, 0xc5, 0xf1,
	0x24, 0xe5, 0x42, 0x94, 0x5b, 0x95, 0x6f, 0x0a, 0x6a, 0x9f, 0x66, 0x9c, 0x67, 0x05, 0x5d, 0x68,
	0x69, 0x53, 0x7e, 0x5f, 0xd0, 0xbb, 0xad, 0xda, 0xd5, 0x95, 0x8e, 0x82, 0xe1, 0x15, 0x95, 0x92,
	0x64, 0x14, 0xfb, 0x30, 0xb8, 0xc8, 0x04, 0x95, 0xd2, 0x42, 0x73, 0xe4, 0x4e, 0x7c, 0xdb, 0xdb,
	0x73, 0xf1, 0x6a, 0xa9, 0xa9, 0x4d, 0x9a, 0x4a, 0xfc, 0x1e, 0x86, 0x21, 0xab, 0x9b, 0x0e, 0x74,
	0xd3, 0x69, 0xa7, 0xa9, 0xd1, 0x4c, 0x97, 0xa9, 0x75, 0xfe, 0x20, 0x98, 0x76, 0x0c, 0xf1, 0x2b,
	0x18, 0x07, 0x3f, 0x08, 0x63, 0xb4, 0x08, 0x97, 0xfa, 0xfd, 0x71, 0xd2, 0x26, 0xb0, 0x0b, 0x4f,
	0x83, 0xda, 0xf6, 0x5a, 0xe4, 0x59, 0xce, 0xc2, 0xa5, 0x7e, 0xee, 0x28, 0x79, 0x98, 0xae, 0x7c,
	0x56, 0x44, 0x64, 0x54, 0x45, 0xe5, 0x9d, 0xd5, 0x9f, 0x23, 0x77, 0x9a, 0xb4, 0x89, 0x56, 0x0d,
	0x97, 0xd2, 0x3a, 0x9c, 0xf7, 0xdd, 0xa3, 0xa4, 0x4d, 0x60, 0x0b, 0x86, 0x31, 0xd9, 0x15, 0x9c,
	0xdc, 0x5a, 0x4f, 0xb4, 0xbb, 0x09, 0xf1, 0x5b, 0x18, 0xe9, 0x75, 0xa5, 0xbc, 0xb0, 0x06, 0x73,
	0xe4, 0x1e, 0xfb, 0xcf, 0x3b, 0x9c, 0x71, 0x23, 0x26, 0xf7, 0x65, 0xce, 0x2f, 0x04, 0xc7, 0x5d,
	0xfc, 0x47, 0x18, 0x6d, 0x18, 0x3d, 0x80, 0xbb, 0x8f, 0xf7, 0xf8, 0xcd, 0xb4, 0x9a, 0xad, 0xe5,
	0x37, 0xe9, 0x7d, 0x86, 0xc3, 0x0e, 0xc3, 0xeb, 0x0f, 0x30, 0x32, 0x63, 0xe2, 0x09, 0x0c, 0xd7,
	0xd1, 0x97, 0xe8, 0xfa, 0x6b, 0x34, 0xeb, 0xd5, 0x41, 0x18, 0x9c, 0xdd, 0xac, 0x66, 0x08, 0x4f,
	0x61, 0x7c, 0xb5, 0xbe, 0x5c, 0xd5, 0xe1, 0x41, 0xa5, 0xc5, 0xeb, 0xf3, 0xcb, 0xf0, 0xe6, 0xd3,
	0xac, 0xef, 0xff, 0x46, 0xf0, 0x32, 0x68, 0xa9, 0x03, 0xce, 0x6e, 0xcb, 0x5c, 0x7d, 0x24, 0xa9,
	0xe2, 0x62, 0x87, 0x83, 0x6a, 0x46, 0xc6, 0x68, 0xaa, 0xce, 0x9a, 0x63, 0xc4, 0x27, 0x5e, 0x7d,
	0x7a, 0x9e, 0x39, 0x3d, 0xef, 0xa2, 0x3a, 0x3d, 0xfb, 0x59, 0x67, 0x79, 0xcd, 0x7a, 0x9c, 0xde,
	0x1b, 0x84, 0x3f, 0xc3, 0x49, 0x2c, 0x78, 0x4a, 0xa5, 0x34, 0x26, 0x66, 0x79, 0xff, 0xec, 0xb1,
	0xff, 0xf3, 0x82, 0xd3, 0x73, 0xd1, 0x39, 0x7c, 0x1b, 0x99, 0xff, 0xb1, 0x19, 0x68, 0xfd, 0xdd,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x9b, 0x0d, 0x6d, 0x32, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CorruptibleConduitFactoryClient is the client API for CorruptibleConduitFactory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CorruptibleConduitFactoryClient interface {
	ConnectAttacker(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CorruptibleConduitFactory_ConnectAttackerClient, error)
	ProcessAttackerMessage(ctx context.Context, opts ...grpc.CallOption) (CorruptibleConduitFactory_ProcessAttackerMessageClient, error)
}

type corruptibleConduitFactoryClient struct {
	cc *grpc.ClientConn
}

func NewCorruptibleConduitFactoryClient(cc *grpc.ClientConn) CorruptibleConduitFactoryClient {
	return &corruptibleConduitFactoryClient{cc}
}

func (c *corruptibleConduitFactoryClient) ConnectAttacker(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CorruptibleConduitFactory_ConnectAttackerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CorruptibleConduitFactory_serviceDesc.Streams[0], "/corruptible.CorruptibleConduitFactory/ConnectAttacker", opts...)
	if err != nil {
		return nil, err
	}
	x := &corruptibleConduitFactoryConnectAttackerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CorruptibleConduitFactory_ConnectAttackerClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type corruptibleConduitFactoryConnectAttackerClient struct {
	grpc.ClientStream
}

func (x *corruptibleConduitFactoryConnectAttackerClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *corruptibleConduitFactoryClient) ProcessAttackerMessage(ctx context.Context, opts ...grpc.CallOption) (CorruptibleConduitFactory_ProcessAttackerMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CorruptibleConduitFactory_serviceDesc.Streams[1], "/corruptible.CorruptibleConduitFactory/ProcessAttackerMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &corruptibleConduitFactoryProcessAttackerMessageClient{stream}
	return x, nil
}

type CorruptibleConduitFactory_ProcessAttackerMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type corruptibleConduitFactoryProcessAttackerMessageClient struct {
	grpc.ClientStream
}

func (x *corruptibleConduitFactoryProcessAttackerMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *corruptibleConduitFactoryProcessAttackerMessageClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CorruptibleConduitFactoryServer is the server API for CorruptibleConduitFactory service.
type CorruptibleConduitFactoryServer interface {
	ConnectAttacker(*emptypb.Empty, CorruptibleConduitFactory_ConnectAttackerServer) error
	ProcessAttackerMessage(CorruptibleConduitFactory_ProcessAttackerMessageServer) error
}

// UnimplementedCorruptibleConduitFactoryServer can be embedded to have forward compatible implementations.
type UnimplementedCorruptibleConduitFactoryServer struct {
}

func (*UnimplementedCorruptibleConduitFactoryServer) ConnectAttacker(req *emptypb.Empty, srv CorruptibleConduitFactory_ConnectAttackerServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectAttacker not implemented")
}
func (*UnimplementedCorruptibleConduitFactoryServer) ProcessAttackerMessage(srv CorruptibleConduitFactory_ProcessAttackerMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessAttackerMessage not implemented")
}

func RegisterCorruptibleConduitFactoryServer(s *grpc.Server, srv CorruptibleConduitFactoryServer) {
	s.RegisterService(&_CorruptibleConduitFactory_serviceDesc, srv)
}

func _CorruptibleConduitFactory_ConnectAttacker_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CorruptibleConduitFactoryServer).ConnectAttacker(m, &corruptibleConduitFactoryConnectAttackerServer{stream})
}

type CorruptibleConduitFactory_ConnectAttackerServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type corruptibleConduitFactoryConnectAttackerServer struct {
	grpc.ServerStream
}

func (x *corruptibleConduitFactoryConnectAttackerServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _CorruptibleConduitFactory_ProcessAttackerMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CorruptibleConduitFactoryServer).ProcessAttackerMessage(&corruptibleConduitFactoryProcessAttackerMessageServer{stream})
}

type CorruptibleConduitFactory_ProcessAttackerMessageServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type corruptibleConduitFactoryProcessAttackerMessageServer struct {
	grpc.ServerStream
}

func (x *corruptibleConduitFactoryProcessAttackerMessageServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *corruptibleConduitFactoryProcessAttackerMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CorruptibleConduitFactory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "corruptible.CorruptibleConduitFactory",
	HandlerType: (*CorruptibleConduitFactoryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectAttacker",
			Handler:       _CorruptibleConduitFactory_ConnectAttacker_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProcessAttackerMessage",
			Handler:       _CorruptibleConduitFactory_ProcessAttackerMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "insecure/attacker.proto",
}
