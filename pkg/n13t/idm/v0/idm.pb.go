// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n13t/idm/v0/idm.proto

package n13t_idm_v0

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

type ListUsersRequest struct {
	Search               string   `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{0}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

// The sum request contains two parameters.
type SumRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

// The sum response contains the result of the calculation.
type SumReply struct {
	V                    int64    `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumReply) Reset()         { *m = SumReply{} }
func (m *SumReply) String() string { return proto.CompactTextString(m) }
func (*SumReply) ProtoMessage()    {}
func (*SumReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{2}
}

func (m *SumReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumReply.Unmarshal(m, b)
}
func (m *SumReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumReply.Marshal(b, m, deterministic)
}
func (m *SumReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumReply.Merge(m, src)
}
func (m *SumReply) XXX_Size() int {
	return xxx_messageInfo_SumReply.Size(m)
}
func (m *SumReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SumReply.DiscardUnknown(m)
}

var xxx_messageInfo_SumReply proto.InternalMessageInfo

func (m *SumReply) GetV() int64 {
	if m != nil {
		return m.V
	}
	return 0
}

func (m *SumReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// The Concat request contains two parameters.
type ConcatRequest struct {
	A                    string   `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    string   `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcatRequest) Reset()         { *m = ConcatRequest{} }
func (m *ConcatRequest) String() string { return proto.CompactTextString(m) }
func (*ConcatRequest) ProtoMessage()    {}
func (*ConcatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{3}
}

func (m *ConcatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcatRequest.Unmarshal(m, b)
}
func (m *ConcatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcatRequest.Marshal(b, m, deterministic)
}
func (m *ConcatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcatRequest.Merge(m, src)
}
func (m *ConcatRequest) XXX_Size() int {
	return xxx_messageInfo_ConcatRequest.Size(m)
}
func (m *ConcatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConcatRequest proto.InternalMessageInfo

func (m *ConcatRequest) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *ConcatRequest) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

// The Concat response contains the result of the concatenation.
type ConcatReply struct {
	V                    string   `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcatReply) Reset()         { *m = ConcatReply{} }
func (m *ConcatReply) String() string { return proto.CompactTextString(m) }
func (*ConcatReply) ProtoMessage()    {}
func (*ConcatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{4}
}

func (m *ConcatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcatReply.Unmarshal(m, b)
}
func (m *ConcatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcatReply.Marshal(b, m, deterministic)
}
func (m *ConcatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcatReply.Merge(m, src)
}
func (m *ConcatReply) XXX_Size() int {
	return xxx_messageInfo_ConcatReply.Size(m)
}
func (m *ConcatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcatReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConcatReply proto.InternalMessageInfo

func (m *ConcatReply) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func (m *ConcatReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type IsValidRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsValidRequest) Reset()         { *m = IsValidRequest{} }
func (m *IsValidRequest) String() string { return proto.CompactTextString(m) }
func (*IsValidRequest) ProtoMessage()    {}
func (*IsValidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{5}
}

func (m *IsValidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsValidRequest.Unmarshal(m, b)
}
func (m *IsValidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsValidRequest.Marshal(b, m, deterministic)
}
func (m *IsValidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsValidRequest.Merge(m, src)
}
func (m *IsValidRequest) XXX_Size() int {
	return xxx_messageInfo_IsValidRequest.Size(m)
}
func (m *IsValidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsValidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsValidRequest proto.InternalMessageInfo

func (m *IsValidRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IsValidRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type IsValidResponse struct {
	V                    bool     `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsValidResponse) Reset()         { *m = IsValidResponse{} }
func (m *IsValidResponse) String() string { return proto.CompactTextString(m) }
func (*IsValidResponse) ProtoMessage()    {}
func (*IsValidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e677661a4128a4, []int{6}
}

func (m *IsValidResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsValidResponse.Unmarshal(m, b)
}
func (m *IsValidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsValidResponse.Marshal(b, m, deterministic)
}
func (m *IsValidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsValidResponse.Merge(m, src)
}
func (m *IsValidResponse) XXX_Size() int {
	return xxx_messageInfo_IsValidResponse.Size(m)
}
func (m *IsValidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsValidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsValidResponse proto.InternalMessageInfo

func (m *IsValidResponse) GetV() bool {
	if m != nil {
		return m.V
	}
	return false
}

func (m *IsValidResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*ListUsersRequest)(nil), "n13t.idm.v0.ListUsersRequest")
	proto.RegisterType((*SumRequest)(nil), "n13t.idm.v0.SumRequest")
	proto.RegisterType((*SumReply)(nil), "n13t.idm.v0.SumReply")
	proto.RegisterType((*ConcatRequest)(nil), "n13t.idm.v0.ConcatRequest")
	proto.RegisterType((*ConcatReply)(nil), "n13t.idm.v0.ConcatReply")
	proto.RegisterType((*IsValidRequest)(nil), "n13t.idm.v0.IsValidRequest")
	proto.RegisterType((*IsValidResponse)(nil), "n13t.idm.v0.IsValidResponse")
}

func init() {
	proto.RegisterFile("n13t/idm/v0/idm.proto", fileDescriptor_c4e677661a4128a4)
}

var fileDescriptor_c4e677661a4128a4 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdf, 0x6b, 0xea, 0x30,
	0x14, 0xc7, 0xa9, 0x85, 0x5e, 0x7b, 0xbc, 0xd7, 0x2b, 0x01, 0x5d, 0xe9, 0xc6, 0x36, 0xfa, 0x24,
	0x8e, 0x55, 0x9d, 0xec, 0x6d, 0x0c, 0x86, 0x7b, 0x11, 0xe6, 0x4b, 0xc5, 0xbd, 0x47, 0x1b, 0xb6,
	0x82, 0xfd, 0xb1, 0x24, 0xed, 0xf0, 0xdf, 0xda, 0x5f, 0x38, 0xd2, 0xa6, 0xda, 0xcc, 0xec, 0x49,
	0x4f, 0xbe, 0x9f, 0xf3, 0x09, 0x27, 0x47, 0xa1, 0x9f, 0x4c, 0x67, 0x7c, 0x1c, 0x85, 0xf1, 0xb8,
	0x98, 0x88, 0x0f, 0x3f, 0xa3, 0x29, 0x4f, 0x51, 0x47, 0x1c, 0xfb, 0xa2, 0x2e, 0x26, 0xee, 0xa0,
	0xc9, 0xe4, 0x8c, 0xd0, 0x0a, 0x72, 0x9d, 0xe6, 0x39, 0x23, 0x98, 0x6e, 0xdf, 0xab, 0xc4, 0x1b,
	0x41, 0xef, 0x25, 0x62, 0x7c, 0xcd, 0x08, 0x65, 0x01, 0xf9, 0xc8, 0x09, 0xe3, 0x68, 0x00, 0x56,
	0xc5, 0x38, 0xc6, 0xb5, 0x31, 0xb4, 0x03, 0x59, 0x79, 0x43, 0x80, 0x55, 0x1e, 0xd7, 0xd4, 0x5f,
	0x30, 0x70, 0x09, 0x98, 0x81, 0x81, 0x45, 0xb5, 0x71, 0x5a, 0x55, 0xb5, 0xf1, 0x46, 0xd0, 0x2e,
	0xc9, 0x6c, 0xb7, 0x17, 0x49, 0x51, 0x73, 0x05, 0xea, 0x81, 0x49, 0x28, 0x2d, 0x49, 0x3b, 0x10,
	0x5f, 0xbd, 0x1b, 0xf8, 0x37, 0x4f, 0x93, 0x2d, 0xe6, 0x27, 0x62, 0x5b, 0x11, 0xdb, 0x42, 0x7c,
	0x0b, 0x9d, 0x1a, 0x56, 0xdc, 0xb6, 0xde, 0xfd, 0x00, 0xdd, 0x05, 0x7b, 0xc5, 0xbb, 0x28, 0xac,
	0xe5, 0x5d, 0x68, 0x45, 0xa1, 0x6c, 0x69, 0x45, 0x21, 0x72, 0xa1, 0x9d, 0x61, 0xc6, 0x3e, 0x53,
	0x1a, 0xca, 0xc6, 0x43, 0xed, 0x4d, 0xe1, 0xff, 0xa1, 0x9b, 0x65, 0x69, 0xc2, 0xc8, 0xf1, 0xc2,
	0xb6, 0xf6, 0xc2, 0xbb, 0x2f, 0x13, 0xd0, 0x22, 0x24, 0x09, 0x8f, 0xf8, 0x7e, 0x89, 0x13, 0xfc,
	0x46, 0x62, 0x92, 0x70, 0x74, 0x0f, 0xe6, 0x2a, 0x8f, 0xd1, 0x99, 0xdf, 0x58, 0x96, 0x7f, 0x7c,
	0x4b, 0xb7, 0x7f, 0x1a, 0x88, 0xf1, 0x1e, 0xc1, 0xaa, 0xa6, 0x45, 0xae, 0x02, 0x28, 0xef, 0xe5,
	0x3a, 0xda, 0x4c, 0xf4, 0x2f, 0x01, 0xe6, 0x94, 0x60, 0x4e, 0xc4, 0x7a, 0xd1, 0xa5, 0xca, 0x1d,
	0x82, 0xda, 0x73, 0xf5, 0x6b, 0x2e, 0x87, 0x5f, 0x02, 0xac, 0xb3, 0x50, 0xaf, 0x3b, 0x06, 0x7a,
	0x5d, 0x33, 0x97, 0xba, 0x27, 0xb0, 0x56, 0xe5, 0x0f, 0xeb, 0xc7, 0x74, 0xd5, 0x61, 0xad, 0x39,
	0xd7, 0x66, 0x52, 0xf1, 0x0c, 0x7f, 0xe4, 0x86, 0x90, 0xca, 0xa9, 0x5b, 0x77, 0x2f, 0xf4, 0x61,
	0x65, 0xd9, 0x58, 0xe5, 0x5f, 0x61, 0xf6, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x38, 0x64, 0x72, 0x1a,
	0x62, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IdentityManagementClient is the client API for IdentityManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityManagementClient interface {
	// Sums two integers.
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
	// Concatenates two strings
	Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	//    rpc DeleteUser(User) returns (google.protobuf.Empty) {
	//        option (google.api.http) = {
	//			delete: "/v0/users/{sub}"
	//		};
	//    }
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// for keycloak to check password
	IsValid(ctx context.Context, in *IsValidRequest, opts ...grpc.CallOption) (*IsValidResponse, error)
}

type identityManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityManagementClient(cc grpc.ClientConnInterface) IdentityManagementClient {
	return &identityManagementClient{cc}
}

func (c *identityManagementClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error) {
	out := new(ConcatReply)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/Concat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) IsValid(ctx context.Context, in *IsValidRequest, opts ...grpc.CallOption) (*IsValidResponse, error) {
	out := new(IsValidResponse)
	err := c.cc.Invoke(ctx, "/n13t.idm.v0.IdentityManagement/IsValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityManagementServer is the server API for IdentityManagement service.
type IdentityManagementServer interface {
	// Sums two integers.
	Sum(context.Context, *SumRequest) (*SumReply, error)
	// Concatenates two strings
	Concat(context.Context, *ConcatRequest) (*ConcatReply, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	//    rpc DeleteUser(User) returns (google.protobuf.Empty) {
	//        option (google.api.http) = {
	//			delete: "/v0/users/{sub}"
	//		};
	//    }
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// for keycloak to check password
	IsValid(context.Context, *IsValidRequest) (*IsValidResponse, error)
}

// UnimplementedIdentityManagementServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityManagementServer struct {
}

func (*UnimplementedIdentityManagementServer) Sum(ctx context.Context, req *SumRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedIdentityManagementServer) Concat(ctx context.Context, req *ConcatRequest) (*ConcatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Concat not implemented")
}
func (*UnimplementedIdentityManagementServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedIdentityManagementServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedIdentityManagementServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedIdentityManagementServer) IsValid(ctx context.Context, req *IsValidRequest) (*IsValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValid not implemented")
}

func RegisterIdentityManagementServer(s *grpc.Server, srv IdentityManagementServer) {
	s.RegisterService(&_IdentityManagement_serviceDesc, srv)
}

func _IdentityManagement_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConcatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).Concat(ctx, req.(*ConcatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_IsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).IsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n13t.idm.v0.IdentityManagement/IsValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).IsValid(ctx, req.(*IsValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n13t.idm.v0.IdentityManagement",
	HandlerType: (*IdentityManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _IdentityManagement_Sum_Handler,
		},
		{
			MethodName: "Concat",
			Handler:    _IdentityManagement_Concat_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _IdentityManagement_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _IdentityManagement_UpdateUser_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _IdentityManagement_Search_Handler,
		},
		{
			MethodName: "IsValid",
			Handler:    _IdentityManagement_IsValid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "n13t/idm/v0/idm.proto",
}
