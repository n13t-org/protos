// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n13t/gitlab/v4/project.proto

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

type Project struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	WebUrl               string   `protobuf:"bytes,4,opt,name=web_url,json=webUrl,proto3" json:"web_url,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	GitSshUrl            string   `protobuf:"bytes,6,opt,name=git_ssh_url,json=gitSshUrl,proto3" json:"git_ssh_url,omitempty"`
	GitHttpUrl           string   `protobuf:"bytes,7,opt,name=git_http_url,json=gitHttpUrl,proto3" json:"git_http_url,omitempty"`
	Namespace            string   `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace,omitempty"`
	VisibilityLevel      uint32   `protobuf:"varint,9,opt,name=visibility_level,json=visibilityLevel,proto3" json:"visibility_level,omitempty"`
	PathWithNamespace    string   `protobuf:"bytes,10,opt,name=path_with_namespace,json=pathWithNamespace,proto3" json:"path_with_namespace,omitempty"`
	DefaultBranch        string   `protobuf:"bytes,11,opt,name=default_branch,json=defaultBranch,proto3" json:"default_branch,omitempty"`
	Homepage             string   `protobuf:"bytes,12,opt,name=homepage,proto3" json:"homepage,omitempty"`
	Url                  string   `protobuf:"bytes,13,opt,name=url,proto3" json:"url,omitempty"`
	SshUrl               string   `protobuf:"bytes,14,opt,name=ssh_url,json=sshUrl,proto3" json:"ssh_url,omitempty"`
	HttpUrl              string   `protobuf:"bytes,15,opt,name=http_url,json=httpUrl,proto3" json:"http_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb2710ec9ccd9b0c, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetWebUrl() string {
	if m != nil {
		return m.WebUrl
	}
	return ""
}

func (m *Project) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *Project) GetGitSshUrl() string {
	if m != nil {
		return m.GitSshUrl
	}
	return ""
}

func (m *Project) GetGitHttpUrl() string {
	if m != nil {
		return m.GitHttpUrl
	}
	return ""
}

func (m *Project) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Project) GetVisibilityLevel() uint32 {
	if m != nil {
		return m.VisibilityLevel
	}
	return 0
}

func (m *Project) GetPathWithNamespace() string {
	if m != nil {
		return m.PathWithNamespace
	}
	return ""
}

func (m *Project) GetDefaultBranch() string {
	if m != nil {
		return m.DefaultBranch
	}
	return ""
}

func (m *Project) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func (m *Project) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Project) GetSshUrl() string {
	if m != nil {
		return m.SshUrl
	}
	return ""
}

func (m *Project) GetHttpUrl() string {
	if m != nil {
		return m.HttpUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Project)(nil), "n13t.gitlab.v4.Project")
}

func init() {
	proto.RegisterFile("n13t/gitlab/v4/project.proto", fileDescriptor_bb2710ec9ccd9b0c)
}

var fileDescriptor_bb2710ec9ccd9b0c = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcd, 0x8a, 0xdb, 0x30,
	0x14, 0x85, 0x71, 0xfe, 0x6c, 0xdf, 0x24, 0x4e, 0xab, 0x2e, 0x2a, 0x42, 0x16, 0xe9, 0x1f, 0x4d,
	0xbb, 0xb0, 0x29, 0xc9, 0x13, 0x64, 0xd5, 0x45, 0x29, 0x25, 0xa5, 0x74, 0x69, 0x64, 0x5b, 0xb5,
	0x54, 0x14, 0x5b, 0xc8, 0x37, 0x0e, 0xf3, 0x06, 0xf3, 0x0e, 0xf3, 0xa6, 0xb3, 0x1a, 0x24, 0x4f,
	0x92, 0xc1, 0xbb, 0xcb, 0x39, 0x9f, 0x2e, 0x9c, 0x73, 0x05, 0xab, 0xea, 0xdb, 0x16, 0x93, 0x52,
	0xa2, 0x62, 0x59, 0xd2, 0xee, 0x12, 0x6d, 0xea, 0xff, 0x3c, 0xc7, 0x58, 0x9b, 0x1a, 0x6b, 0x12,
	0x59, 0x37, 0xee, 0xdc, 0xb8, 0xdd, 0x2d, 0xdf, 0xb6, 0x4c, 0xc9, 0x82, 0x21, 0x4f, 0x2e, 0x43,
	0x07, 0xbe, 0x7f, 0x18, 0x81, 0xff, 0xab, 0x7b, 0x4a, 0x22, 0x18, 0xc8, 0x82, 0x7a, 0x6b, 0x6f,
	0x33, 0x3f, 0x0c, 0x64, 0x41, 0x08, 0x8c, 0x2a, 0x76, 0xe4, 0x74, 0xb0, 0xf6, 0x36, 0xe1, 0xc1,
	0xcd, 0x64, 0x0d, 0xd3, 0x82, 0x37, 0xb9, 0x91, 0x1a, 0x65, 0x5d, 0xd1, 0xa1, 0xb3, 0x5e, 0x4a,
	0xe4, 0x1d, 0xf8, 0x67, 0x9e, 0xa5, 0x27, 0xa3, 0xe8, 0xc8, 0xba, 0xfb, 0xe0, 0x71, 0x3f, 0x36,
	0xc3, 0x7b, 0xcf, 0x3b, 0x4c, 0xce, 0x3c, 0xfb, 0x63, 0x14, 0xf9, 0x0c, 0xc0, 0x5a, 0x86, 0xcc,
	0x38, 0x6a, 0xdc, 0xa3, 0xc2, 0xce, 0xb3, 0xe0, 0x06, 0xa6, 0xa5, 0xc4, 0xb4, 0x69, 0x84, 0x23,
	0x27, 0x7d, 0xb2, 0x94, 0xf8, 0xbb, 0x11, 0x96, 0xfc, 0x0a, 0x33, 0x4b, 0x0a, 0x44, 0xed, 0x50,
	0xbf, 0x87, 0x42, 0x29, 0xf1, 0x3b, 0xa2, 0xb6, 0xec, 0x0a, 0x42, 0x9b, 0xa5, 0xd1, 0x2c, 0xe7,
	0x34, 0x70, 0x09, 0x6e, 0x02, 0xf9, 0x02, 0xaf, 0x5a, 0xd9, 0xc8, 0x4c, 0x2a, 0x89, 0x77, 0xa9,
	0xe2, 0x2d, 0x57, 0x34, 0x74, 0x9d, 0x2c, 0x6e, 0xfa, 0x0f, 0x2b, 0x93, 0x18, 0xde, 0x68, 0x86,
	0x22, 0x3d, 0x4b, 0x14, 0xe9, 0x6d, 0x25, 0xb8, 0x95, 0xaf, 0xad, 0xf5, 0x57, 0xa2, 0xf8, 0x79,
	0x5d, 0xfd, 0x09, 0xa2, 0x82, 0xff, 0x63, 0x27, 0x85, 0x69, 0x66, 0x58, 0x95, 0x0b, 0x3a, 0x75,
	0xe8, 0xfc, 0x59, 0xdd, 0x3b, 0x91, 0x7c, 0x84, 0x40, 0xd4, 0x47, 0xae, 0x59, 0xc9, 0xe9, 0xac,
	0x97, 0xe3, 0xea, 0x90, 0x25, 0x0c, 0x6d, 0xd0, 0x79, 0x0f, 0xb0, 0xa2, 0xbd, 0xc1, 0xa5, 0xb3,
	0xa8, 0x7f, 0x83, 0xa6, 0x2b, 0xec, 0x03, 0x04, 0xd7, 0xb2, 0x16, 0x3d, 0xc6, 0x17, 0x5d, 0x53,
	0xd9, 0xc4, 0x7d, 0x92, 0xed, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xb3, 0x85, 0xa6, 0x6d,
	0x02, 0x00, 0x00,
}
