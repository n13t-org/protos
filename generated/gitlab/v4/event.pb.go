// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package n13t_gitlab_v4

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// See https://docs.gitlab.com/ee/api/events.html
type Event struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	TargetType           string   `protobuf:"bytes,2,opt,name=target_type,json=targetType,proto3" json:"target_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Event) GetTargetType() string {
	if m != nil {
		return m.TargetType
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "n13t.gitlab.v4.Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x55, 0xd4, 0x36, 0x24, 0x15, 0x0c, 0x59, 0xa8, 0x98, 0x10, 0x13, 0x53, 0x10, 0x2a,
	0x3b, 0x22, 0x82, 0x17, 0x38, 0xb1, 0x57, 0xae, 0xfd, 0x37, 0xb8, 0x4a, 0x6c, 0x73, 0xbe, 0x44,
	0x64, 0xe1, 0x91, 0x79, 0x80, 0x4e, 0xc8, 0x89, 0x98, 0xee, 0xd7, 0xe9, 0xbb, 0xef, 0xee, 0x8a,
	0x0d, 0x06, 0x38, 0xa9, 0x02, 0x7b, 0xf1, 0xe5, 0xb5, 0x7b, 0xda, 0x49, 0xd5, 0x58, 0x69, 0xd5,
	0xa1, 0x1a, 0x9e, 0x6f, 0x6f, 0x06, 0xd5, 0x5a, 0xa3, 0x04, 0x8f, 0xff, 0x61, 0x06, 0xef, 0x7f,
	0x17, 0xc5, 0xea, 0x3d, 0x0d, 0x96, 0x3f, 0xc5, 0x5a, 0x69, 0xb1, 0xde, 0x6d, 0x17, 0x77, 0x8b,
	0x87, 0xbc, 0x3e, 0x9e, 0x6b, 0xcd, 0x8a, 0x32, 0xcd, 0x50, 0x02, 0x43, 0x59, 0x1f, 0xcc, 0x14,
	0xd6, 0xba, 0xf5, 0x11, 0x86, 0x2e, 0x19, 0x3e, 0xc0, 0xa5, 0x4e, 0xe8, 0xe3, 0x27, 0x0c, 0xe5,
	0xda, 0x77, 0x1d, 0xdc, 0x04, 0x75, 0xe0, 0x26, 0xd5, 0x93, 0xb7, 0x09, 0x59, 0xb6, 0x38, 0x0a,
	0xe5, 0x06, 0x51, 0xd8, 0x8f, 0x49, 0x87, 0xef, 0x60, 0x39, 0x11, 0xf3, 0xd6, 0x12, 0xc5, 0x46,
	0x14, 0x37, 0x90, 0xbd, 0x8c, 0x01, 0xdb, 0x8b, 0xe9, 0x88, 0xb7, 0x73, 0xfd, 0xca, 0x2f, 0xb4,
	0xb2, 0x31, 0xf6, 0xa0, 0xbc, 0xb3, 0x2d, 0xa2, 0x78, 0x07, 0xba, 0x9a, 0xfc, 0x7b, 0xc6, 0x57,
	0x8f, 0x28, 0xb4, 0x74, 0x5e, 0x40, 0x59, 0x60, 0x7f, 0x82, 0x16, 0xca, 0xa2, 0xb3, 0x21, 0x40,
	0x68, 0xd9, 0x47, 0x30, 0x15, 0xb3, 0xf8, 0x63, 0x0c, 0x38, 0xac, 0xa7, 0xbf, 0x77, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x77, 0x48, 0x76, 0x66, 0x2f, 0x01, 0x00, 0x00,
}
