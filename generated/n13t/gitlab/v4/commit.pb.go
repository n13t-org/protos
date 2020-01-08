// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n13t/gitlab/v4/commit.proto

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

type Commit struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Author               *Author  `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Added                []string `protobuf:"bytes,6,rep,name=added,proto3" json:"added,omitempty"`
	Modified             []string `protobuf:"bytes,7,rep,name=modified,proto3" json:"modified,omitempty"`
	Removed              []string `protobuf:"bytes,8,rep,name=removed,proto3" json:"removed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_5025854e35f4382a, []int{0}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Commit) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Commit) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Commit) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Commit) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Commit) GetAdded() []string {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *Commit) GetModified() []string {
	if m != nil {
		return m.Modified
	}
	return nil
}

func (m *Commit) GetRemoved() []string {
	if m != nil {
		return m.Removed
	}
	return nil
}

func init() {
	proto.RegisterType((*Commit)(nil), "n13t.gitlab.v4.Commit")
}

func init() { proto.RegisterFile("n13t/gitlab/v4/commit.proto", fileDescriptor_5025854e35f4382a) }

var fileDescriptor_5025854e35f4382a = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0xe5, 0x84, 0xa4, 0xc9, 0x20, 0x75, 0x61, 0x21, 0x18, 0x05, 0x16, 0x11, 0xab, 0xac,
	0x1c, 0x41, 0x7b, 0x01, 0xca, 0x0d, 0x72, 0x03, 0x97, 0x31, 0xc5, 0x52, 0x8c, 0x2b, 0xc7, 0xcd,
	0x19, 0x38, 0x2a, 0x67, 0x60, 0x85, 0x6c, 0x13, 0x10, 0xdd, 0xf9, 0xbf, 0x37, 0x1a, 0xcf, 0x87,
	0xdb, 0xf7, 0x87, 0x8d, 0xef, 0x0f, 0xda, 0x8f, 0x72, 0xdf, 0xcf, 0xdb, 0xfe, 0xc5, 0x1a, 0xa3,
	0xbd, 0x38, 0x3a, 0xeb, 0x2d, 0x5f, 0x07, 0x29, 0x92, 0x14, 0xf3, 0xb6, 0xb9, 0x99, 0xe5, 0xa8,
	0x49, 0x7a, 0xd5, 0x2f, 0x8f, 0x34, 0xd8, 0x9c, 0x6f, 0x91, 0x27, 0xff, 0x66, 0x5d, 0x92, 0xf7,
	0x9f, 0x0c, 0xca, 0xe7, 0xb8, 0x96, 0xaf, 0x21, 0xd3, 0x84, 0xac, 0x65, 0x5d, 0x3d, 0x64, 0x9a,
	0x38, 0xc2, 0xca, 0xa8, 0x69, 0x92, 0x07, 0x85, 0x59, 0x84, 0x4b, 0xe4, 0x77, 0x50, 0x7b, 0x6d,
	0xd4, 0xe4, 0xa5, 0x39, 0x62, 0x1e, 0xdd, 0x1f, 0xe0, 0x0d, 0xe4, 0x27, 0x37, 0xe2, 0x45, 0xe0,
	0xbb, 0xea, 0x6b, 0x57, 0xb8, 0xfc, 0x83, 0xb1, 0x21, 0x40, 0x2e, 0xa0, 0x4c, 0xdf, 0x63, 0xd1,
	0xb2, 0xee, 0xf2, 0xf1, 0x5a, 0xfc, 0x6f, 0x21, 0x9e, 0xa2, 0x1d, 0x7e, 0xa6, 0xf8, 0x15, 0x14,
	0x92, 0x48, 0x11, 0x96, 0x6d, 0xde, 0xd5, 0x43, 0x0a, 0xbc, 0x81, 0xca, 0x58, 0xd2, 0xaf, 0x5a,
	0x11, 0xae, 0xa2, 0xf8, 0xcd, 0xe1, 0x6a, 0xa7, 0x8c, 0x9d, 0x15, 0x61, 0x15, 0xd5, 0x12, 0xf7,
	0x65, 0x6c, 0xbc, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x64, 0xd1, 0xb0, 0x57, 0x56, 0x01, 0x00,
	0x00,
}
