// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pcol.proto

package pcol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RecvChat struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecvChat) Reset()         { *m = RecvChat{} }
func (m *RecvChat) String() string { return proto.CompactTextString(m) }
func (*RecvChat) ProtoMessage()    {}
func (*RecvChat) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d5e4618b2e5f5d1, []int{0}
}

func (m *RecvChat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecvChat.Unmarshal(m, b)
}
func (m *RecvChat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecvChat.Marshal(b, m, deterministic)
}
func (m *RecvChat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecvChat.Merge(m, src)
}
func (m *RecvChat) XXX_Size() int {
	return xxx_messageInfo_RecvChat.Size(m)
}
func (m *RecvChat) XXX_DiscardUnknown() {
	xxx_messageInfo_RecvChat.DiscardUnknown(m)
}

var xxx_messageInfo_RecvChat proto.InternalMessageInfo

func (m *RecvChat) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RecvChat) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RecvBroadcast struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecvBroadcast) Reset()         { *m = RecvBroadcast{} }
func (m *RecvBroadcast) String() string { return proto.CompactTextString(m) }
func (*RecvBroadcast) ProtoMessage()    {}
func (*RecvBroadcast) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d5e4618b2e5f5d1, []int{1}
}

func (m *RecvBroadcast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecvBroadcast.Unmarshal(m, b)
}
func (m *RecvBroadcast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecvBroadcast.Marshal(b, m, deterministic)
}
func (m *RecvBroadcast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecvBroadcast.Merge(m, src)
}
func (m *RecvBroadcast) XXX_Size() int {
	return xxx_messageInfo_RecvBroadcast.Size(m)
}
func (m *RecvBroadcast) XXX_DiscardUnknown() {
	xxx_messageInfo_RecvBroadcast.DiscardUnknown(m)
}

var xxx_messageInfo_RecvBroadcast proto.InternalMessageInfo

func (m *RecvBroadcast) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *RecvBroadcast) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RegisterClient struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              int32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterClient) Reset()         { *m = RegisterClient{} }
func (m *RegisterClient) String() string { return proto.CompactTextString(m) }
func (*RegisterClient) ProtoMessage()    {}
func (*RegisterClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d5e4618b2e5f5d1, []int{2}
}

func (m *RegisterClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterClient.Unmarshal(m, b)
}
func (m *RegisterClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterClient.Marshal(b, m, deterministic)
}
func (m *RegisterClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterClient.Merge(m, src)
}
func (m *RegisterClient) XXX_Size() int {
	return xxx_messageInfo_RegisterClient.Size(m)
}
func (m *RegisterClient) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterClient.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterClient proto.InternalMessageInfo

func (m *RegisterClient) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterClient) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type SendChat struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendChat) Reset()         { *m = SendChat{} }
func (m *SendChat) String() string { return proto.CompactTextString(m) }
func (*SendChat) ProtoMessage()    {}
func (*SendChat) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d5e4618b2e5f5d1, []int{3}
}

func (m *SendChat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendChat.Unmarshal(m, b)
}
func (m *SendChat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendChat.Marshal(b, m, deterministic)
}
func (m *SendChat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendChat.Merge(m, src)
}
func (m *SendChat) XXX_Size() int {
	return xxx_messageInfo_SendChat.Size(m)
}
func (m *SendChat) XXX_DiscardUnknown() {
	xxx_messageInfo_SendChat.DiscardUnknown(m)
}

var xxx_messageInfo_SendChat proto.InternalMessageInfo

func (m *SendChat) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RecvChat)(nil), "pcol.RecvChat")
	proto.RegisterType((*RecvBroadcast)(nil), "pcol.RecvBroadcast")
	proto.RegisterType((*RegisterClient)(nil), "pcol.RegisterClient")
	proto.RegisterType((*SendChat)(nil), "pcol.SendChat")
}

func init() { proto.RegisterFile("pcol.proto", fileDescriptor_1d5e4618b2e5f5d1) }

var fileDescriptor_1d5e4618b2e5f5d1 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xb1, 0xce, 0x82, 0x30,
	0x14, 0x85, 0xc3, 0x1f, 0x7e, 0xc4, 0x9b, 0xe8, 0xd0, 0xc1, 0x10, 0x27, 0x43, 0x1c, 0x9c, 0x5c,
	0xdc, 0x8d, 0xca, 0x1b, 0xd4, 0x27, 0xa8, 0x70, 0x82, 0x24, 0xd0, 0x92, 0xde, 0xca, 0xf3, 0x1b,
	0x1a, 0x6a, 0x70, 0x71, 0xeb, 0x97, 0xe6, 0x7e, 0xe7, 0x1c, 0xa2, 0xbe, 0x34, 0xed, 0xb1, 0xb7,
	0xc6, 0x19, 0x11, 0x8f, 0xef, 0xfc, 0x42, 0xa9, 0x44, 0x39, 0x14, 0x4f, 0xe5, 0xc4, 0x96, 0xd2,
	0x17, 0xc3, 0x6a, 0xd5, 0x21, 0x8b, 0x76, 0xd1, 0x61, 0x29, 0x3f, 0x2c, 0x32, 0x5a, 0x74, 0x60,
	0x56, 0x35, 0xb2, 0x3f, 0xff, 0x15, 0x30, 0xbf, 0xd2, 0x6a, 0x34, 0xdc, 0xac, 0x51, 0x55, 0xa9,
	0xd8, 0x89, 0x0d, 0x25, 0x0c, 0x5d, 0xc1, 0x4e, 0x92, 0x89, 0x7e, 0x28, 0xce, 0xb4, 0x96, 0xa8,
	0x1b, 0x76, 0xb0, 0x45, 0xdb, 0x40, 0x3b, 0x21, 0x28, 0x9e, 0xd5, 0x88, 0x43, 0x85, 0x01, 0x96,
	0x1b, 0xa3, 0xfd, 0xfd, 0xbf, 0x0c, 0x98, 0xef, 0x29, 0xbd, 0x43, 0x57, 0x7e, 0xc4, 0x2c, 0x25,
	0xfa, 0x4a, 0x79, 0x24, 0x7e, 0xf7, 0xe9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x68, 0x12, 0xea, 0x8d,
	0x05, 0x01, 0x00, 0x00,
}
