// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n13t/gitlab/v4/webhook-event.proto

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
type WebhookEvent struct {
	ObjectKind           string   `protobuf:"bytes,1,opt,name=object_kind,json=objectKind,proto3" json:"object_kind,omitempty"`
	Author               *Author  `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebhookEvent) Reset()         { *m = WebhookEvent{} }
func (m *WebhookEvent) String() string { return proto.CompactTextString(m) }
func (*WebhookEvent) ProtoMessage()    {}
func (*WebhookEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0697965f6806d59c, []int{0}
}

func (m *WebhookEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookEvent.Unmarshal(m, b)
}
func (m *WebhookEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookEvent.Marshal(b, m, deterministic)
}
func (m *WebhookEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookEvent.Merge(m, src)
}
func (m *WebhookEvent) XXX_Size() int {
	return xxx_messageInfo_WebhookEvent.Size(m)
}
func (m *WebhookEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookEvent proto.InternalMessageInfo

func (m *WebhookEvent) GetObjectKind() string {
	if m != nil {
		return m.ObjectKind
	}
	return ""
}

func (m *WebhookEvent) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func init() {
	proto.RegisterType((*WebhookEvent)(nil), "n13t.gitlab.v4.WebhookEvent")
}

func init() { proto.RegisterFile("n13t/gitlab/v4/webhook-event.proto", fileDescriptor_0697965f6806d59c) }

var fileDescriptor_0697965f6806d59c = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x15, 0x86, 0x4a, 0xb8, 0x88, 0x21, 0x03, 0x54, 0x65, 0xa9, 0x2a, 0x86, 0x30, 0xe0,
	0x88, 0xb6, 0x3f, 0x40, 0x51, 0x85, 0x04, 0x5b, 0x17, 0x36, 0x22, 0x9b, 0x3c, 0x25, 0x8f, 0x04,
	0xdb, 0xd8, 0xcf, 0xee, 0xe2, 0x3f, 0xe2, 0x0f, 0x99, 0x50, 0x62, 0x18, 0xe8, 0x76, 0xa5, 0x7b,
	0x74, 0x74, 0x2f, 0x5b, 0xaa, 0xbb, 0x35, 0x95, 0x0d, 0x52, 0x2f, 0x64, 0x19, 0x36, 0xe5, 0x01,
	0x64, 0xab, 0x75, 0x77, 0x0b, 0x01, 0x14, 0x71, 0x63, 0x35, 0xe9, 0xfc, 0x7c, 0x60, 0x78, 0x62,
	0x78, 0xd8, 0xcc, 0x2f, 0x83, 0xe8, 0xb1, 0x16, 0x04, 0xe5, 0x5f, 0x48, 0xe0, 0xfc, 0xea, 0x48,
	0x26, 0x3c, 0xb5, 0xda, 0xa6, 0x72, 0xf9, 0x95, 0xb1, 0xb3, 0x97, 0x64, 0xdf, 0x0d, 0xf2, 0xbc,
	0x63, 0x53, 0x2d, 0xdf, 0xe1, 0x8d, 0xaa, 0x0e, 0x55, 0x3d, 0xcb, 0x16, 0x59, 0x71, 0xba, 0x7d,
	0xfa, 0xde, 0x3e, 0xda, 0xdd, 0xea, 0xe1, 0xb5, 0x30, 0xde, 0xb5, 0x91, 0x44, 0x53, 0x8d, 0x01,
	0x9d, 0xf3, 0x10, 0x95, 0x26, 0x88, 0x1f, 0x60, 0x1b, 0xa8, 0x2c, 0x7c, 0x7a, 0x70, 0x14, 0x0f,
	0xd8, 0x61, 0x65, 0x44, 0x03, 0xd1, 0xa0, 0x81, 0x1e, 0x15, 0x44, 0xe9, 0xb1, 0xaf, 0x6f, 0xae,
	0xf7, 0x2c, 0xe9, 0x9f, 0x51, 0xd5, 0x39, 0x67, 0x93, 0xb4, 0x66, 0x76, 0xb2, 0xc8, 0x8a, 0xe9,
	0xea, 0x82, 0xff, 0x3f, 0xc5, 0xef, 0xc7, 0x76, 0xff, 0x4b, 0xc9, 0xc9, 0x38, 0x7a, 0xfd, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x61, 0x6f, 0x25, 0x53, 0x20, 0x01, 0x00, 0x00,
}
