// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rg.proto

package rgpb

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

type ResourceGroupLimit struct {
	NodeNum              int32    `protobuf:"varint,1,opt,name=nodeNum,proto3" json:"nodeNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceGroupLimit) Reset()         { *m = ResourceGroupLimit{} }
func (m *ResourceGroupLimit) String() string { return proto.CompactTextString(m) }
func (*ResourceGroupLimit) ProtoMessage()    {}
func (*ResourceGroupLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_103cbf007c63c19b, []int{0}
}

func (m *ResourceGroupLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceGroupLimit.Unmarshal(m, b)
}
func (m *ResourceGroupLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceGroupLimit.Marshal(b, m, deterministic)
}
func (m *ResourceGroupLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceGroupLimit.Merge(m, src)
}
func (m *ResourceGroupLimit) XXX_Size() int {
	return xxx_messageInfo_ResourceGroupLimit.Size(m)
}
func (m *ResourceGroupLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceGroupLimit.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceGroupLimit proto.InternalMessageInfo

func (m *ResourceGroupLimit) GetNodeNum() int32 {
	if m != nil {
		return m.NodeNum
	}
	return 0
}

type ResourceGroupTransfer struct {
	ResourceGroup        string   `protobuf:"bytes,1,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceGroupTransfer) Reset()         { *m = ResourceGroupTransfer{} }
func (m *ResourceGroupTransfer) String() string { return proto.CompactTextString(m) }
func (*ResourceGroupTransfer) ProtoMessage()    {}
func (*ResourceGroupTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_103cbf007c63c19b, []int{1}
}

func (m *ResourceGroupTransfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceGroupTransfer.Unmarshal(m, b)
}
func (m *ResourceGroupTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceGroupTransfer.Marshal(b, m, deterministic)
}
func (m *ResourceGroupTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceGroupTransfer.Merge(m, src)
}
func (m *ResourceGroupTransfer) XXX_Size() int {
	return xxx_messageInfo_ResourceGroupTransfer.Size(m)
}
func (m *ResourceGroupTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceGroupTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceGroupTransfer proto.InternalMessageInfo

func (m *ResourceGroupTransfer) GetResourceGroup() string {
	if m != nil {
		return m.ResourceGroup
	}
	return ""
}

type ResourceGroupConfig struct {
	Requests             *ResourceGroupLimit      `protobuf:"bytes,1,opt,name=requests,proto3" json:"requests,omitempty"`
	Limits               *ResourceGroupLimit      `protobuf:"bytes,2,opt,name=limits,proto3" json:"limits,omitempty"`
	From                 []*ResourceGroupTransfer `protobuf:"bytes,3,rep,name=from,proto3" json:"from,omitempty"`
	To                   []*ResourceGroupTransfer `protobuf:"bytes,4,rep,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ResourceGroupConfig) Reset()         { *m = ResourceGroupConfig{} }
func (m *ResourceGroupConfig) String() string { return proto.CompactTextString(m) }
func (*ResourceGroupConfig) ProtoMessage()    {}
func (*ResourceGroupConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_103cbf007c63c19b, []int{2}
}

func (m *ResourceGroupConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceGroupConfig.Unmarshal(m, b)
}
func (m *ResourceGroupConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceGroupConfig.Marshal(b, m, deterministic)
}
func (m *ResourceGroupConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceGroupConfig.Merge(m, src)
}
func (m *ResourceGroupConfig) XXX_Size() int {
	return xxx_messageInfo_ResourceGroupConfig.Size(m)
}
func (m *ResourceGroupConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceGroupConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceGroupConfig proto.InternalMessageInfo

func (m *ResourceGroupConfig) GetRequests() *ResourceGroupLimit {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *ResourceGroupConfig) GetLimits() *ResourceGroupLimit {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *ResourceGroupConfig) GetFrom() []*ResourceGroupTransfer {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ResourceGroupConfig) GetTo() []*ResourceGroupTransfer {
	if m != nil {
		return m.To
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceGroupLimit)(nil), "milvus.protov2.rg.ResourceGroupLimit")
	proto.RegisterType((*ResourceGroupTransfer)(nil), "milvus.protov2.rg.ResourceGroupTransfer")
	proto.RegisterType((*ResourceGroupConfig)(nil), "milvus.protov2.rg.ResourceGroupConfig")
}

func init() { proto.RegisterFile("rg.proto", fileDescriptor_103cbf007c63c19b) }

var fileDescriptor_103cbf007c63c19b = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0x37, 0xe7, 0xcc, 0x70, 0x60, 0x44, 0xe8, 0x71, 0x14, 0x06, 0xbd, 0x2c, 0x95,
	0x7a, 0xf1, 0xa0, 0x82, 0xdb, 0x61, 0x17, 0x95, 0x51, 0x3c, 0x79, 0x91, 0xb5, 0xa6, 0x31, 0xd0,
	0xf6, 0xc5, 0x97, 0xa4, 0x5f, 0xc3, 0xaf, 0xe0, 0xd9, 0x4f, 0x29, 0x6d, 0x57, 0xb1, 0xec, 0x20,
	0xbb, 0xe5, 0xe5, 0xfd, 0x7f, 0xbf, 0xbc, 0xf0, 0xc8, 0x18, 0x05, 0x53, 0x08, 0x06, 0xe8, 0x59,
	0x21, 0xf3, 0xca, 0xea, 0xb6, 0xaa, 0x22, 0x86, 0xc2, 0x67, 0x84, 0xc6, 0x5c, 0x83, 0xc5, 0x94,
	0xaf, 0x11, 0xac, 0x7a, 0x90, 0x85, 0x34, 0xd4, 0x23, 0xc7, 0x25, 0xbc, 0xf1, 0x27, 0x5b, 0x78,
	0xce, 0xcc, 0x09, 0x8e, 0xe2, 0xae, 0xf4, 0xef, 0xc8, 0x45, 0x2f, 0xff, 0x8c, 0xdb, 0x52, 0x67,
	0x1c, 0xe9, 0x9c, 0x4c, 0x71, 0xd7, 0x78, 0x15, 0x75, 0xa7, 0x21, 0x4f, 0xe2, 0x53, 0xfc, 0x1b,
	0xf7, 0x3f, 0x5d, 0x72, 0xde, 0x13, 0xac, 0xa0, 0xcc, 0xa4, 0xa0, 0xf7, 0x64, 0x8c, 0xfc, 0xc3,
	0x72, 0x6d, 0x74, 0x03, 0x4e, 0xa2, 0x39, 0xdb, 0x9b, 0x96, 0xed, 0x8f, 0x1a, 0xff, 0x62, 0xf4,
	0x96, 0x8c, 0xf2, 0xfa, 0x4a, 0x7b, 0xee, 0x21, 0x82, 0x1d, 0x44, 0x6f, 0xc8, 0x30, 0x43, 0x28,
	0xbc, 0xc1, 0x6c, 0x10, 0x4c, 0xa2, 0xe0, 0x3f, 0xb8, 0xfb, 0x78, 0xdc, 0x50, 0xf4, 0x9a, 0xb8,
	0x06, 0xbc, 0xe1, 0x81, 0xac, 0x6b, 0x60, 0xa9, 0xc8, 0x54, 0x42, 0x47, 0x08, 0x54, 0xe9, 0xb2,
	0xbf, 0x91, 0x4d, 0x2d, 0xd9, 0x38, 0x2f, 0x97, 0x42, 0x9a, 0x77, 0x9b, 0xb0, 0x14, 0x8a, 0xb0,
	0x4d, 0x2f, 0x24, 0x74, 0xa7, 0xe6, 0xa5, 0x50, 0xc0, 0x62, 0xab, 0x64, 0x58, 0x45, 0x21, 0x0a,
	0x95, 0x7c, 0x39, 0xce, 0xb7, 0x4b, 0x1f, 0x5b, 0xf1, 0x2a, 0x97, 0xbc, 0x34, 0x6c, 0x8d, 0x2a,
	0x4d, 0x46, 0x4d, 0xfa, 0xea, 0x27, 0x00, 0x00, 0xff, 0xff, 0x42, 0xf7, 0xb8, 0x5f, 0x19, 0x02,
	0x00, 0x00,
}
