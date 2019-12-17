// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1alpha1.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Protocol Buffers messages
type User struct {
	Enabled    bool                `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Attributes map[string]*any.Any `protobuf:"bytes,100,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// OpenId Standard Claims Members (https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims)
	Sub                  string               `protobuf:"bytes,10,opt,name=sub,proto3" json:"sub,omitempty"`
	Name                 string               `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	GivenName            string               `protobuf:"bytes,12,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           string               `protobuf:"bytes,14,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	MiddleName           string               `protobuf:"bytes,15,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	Nickname             string               `protobuf:"bytes,16,opt,name=nickname,proto3" json:"nickname,omitempty"`
	PreferredUsername    string               `protobuf:"bytes,17,opt,name=preferred_username,json=preferredUsername,proto3" json:"preferred_username,omitempty"`
	Profile              string               `protobuf:"bytes,20,opt,name=profile,proto3" json:"profile,omitempty"`
	Picture              string               `protobuf:"bytes,21,opt,name=picture,proto3" json:"picture,omitempty"`
	Website              string               `protobuf:"bytes,22,opt,name=website,proto3" json:"website,omitempty"`
	Email                string               `protobuf:"bytes,30,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified        bool                 `protobuf:"varint,31,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Gender               string               `protobuf:"bytes,32,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthdate            string               `protobuf:"bytes,33,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Zoneinfo             string               `protobuf:"bytes,34,opt,name=zoneinfo,proto3" json:"zoneinfo,omitempty"`
	Locale               string               `protobuf:"bytes,35,opt,name=locale,proto3" json:"locale,omitempty"`
	PhoneNumber          string               `protobuf:"bytes,40,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	PhoneNumberVerified  bool                 `protobuf:"varint,41,opt,name=phone_number_verified,json=phoneNumberVerified,proto3" json:"phone_number_verified,omitempty"`
	Address              *User_Address        `protobuf:"bytes,42,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,102,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff3cb1123904c43b, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *User) GetAttributes() map[string]*any.Any {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *User) GetSub() string {
	if m != nil {
		return m.Sub
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetGivenName() string {
	if m != nil {
		return m.GivenName
	}
	return ""
}

func (m *User) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *User) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetPreferredUsername() string {
	if m != nil {
		return m.PreferredUsername
	}
	return ""
}

func (m *User) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *User) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *User) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
}

func (m *User) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *User) GetBirthdate() string {
	if m != nil {
		return m.Birthdate
	}
	return ""
}

func (m *User) GetZoneinfo() string {
	if m != nil {
		return m.Zoneinfo
	}
	return ""
}

func (m *User) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetPhoneNumberVerified() bool {
	if m != nil {
		return m.PhoneNumberVerified
	}
	return false
}

func (m *User) GetAddress() *User_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type User_Address struct {
	Formatted            string   `protobuf:"bytes,1,opt,name=formatted,proto3" json:"formatted,omitempty"`
	StreetAddress        string   `protobuf:"bytes,2,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	Locality             string   `protobuf:"bytes,3,opt,name=locality,proto3" json:"locality,omitempty"`
	Region               string   `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	PostalCode           string   `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Country              string   `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User_Address) Reset()         { *m = User_Address{} }
func (m *User_Address) String() string { return proto.CompactTextString(m) }
func (*User_Address) ProtoMessage()    {}
func (*User_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff3cb1123904c43b, []int{0, 1}
}

func (m *User_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User_Address.Unmarshal(m, b)
}
func (m *User_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User_Address.Marshal(b, m, deterministic)
}
func (m *User_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User_Address.Merge(m, src)
}
func (m *User_Address) XXX_Size() int {
	return xxx_messageInfo_User_Address.Size(m)
}
func (m *User_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_User_Address.DiscardUnknown(m)
}

var xxx_messageInfo_User_Address proto.InternalMessageInfo

func (m *User_Address) GetFormatted() string {
	if m != nil {
		return m.Formatted
	}
	return ""
}

func (m *User_Address) GetStreetAddress() string {
	if m != nil {
		return m.StreetAddress
	}
	return ""
}

func (m *User_Address) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *User_Address) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *User_Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *User_Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

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
	return fileDescriptor_ff3cb1123904c43b, []int{1}
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

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff3cb1123904c43b, []int{2}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "nghinhut.idm.v1alpha1.User")
	proto.RegisterMapType((map[string]*any.Any)(nil), "nghinhut.idm.v1alpha1.User.AttributesEntry")
	proto.RegisterType((*User_Address)(nil), "nghinhut.idm.v1alpha1.User.Address")
	proto.RegisterType((*ListUsersRequest)(nil), "nghinhut.idm.v1alpha1.ListUsersRequest")
	proto.RegisterType((*SearchRequest)(nil), "nghinhut.idm.v1alpha1.SearchRequest")
}

func init() { proto.RegisterFile("v1alpha1.proto", fileDescriptor_ff3cb1123904c43b) }

var fileDescriptor_ff3cb1123904c43b = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0xef, 0x26, 0xd9, 0x64, 0xf7, 0x6d, 0xf3, 0xa7, 0x43, 0x12, 0x0d, 0x2e, 0x90, 0xed, 0x56,
	0xc0, 0x12, 0x84, 0x4b, 0xc3, 0x05, 0x90, 0x38, 0x84, 0xd0, 0x03, 0x02, 0xaa, 0xca, 0x21, 0x1c,
	0xe0, 0x60, 0xcd, 0xee, 0xbc, 0xf5, 0x8e, 0x62, 0x8f, 0xdd, 0xf1, 0x38, 0xc8, 0x7c, 0x37, 0x24,
	0xbe, 0x04, 0xdf, 0x07, 0xcd, 0x1b, 0xcf, 0x36, 0xda, 0x42, 0x38, 0xe4, 0x36, 0xbf, 0x3f, 0xef,
	0x79, 0xde, 0x9f, 0x31, 0xec, 0xdd, 0x3c, 0x17, 0x79, 0xb5, 0x14, 0xcf, 0xe3, 0xca, 0x94, 0xb6,
	0x64, 0x47, 0x3a, 0x5b, 0x2a, 0xbd, 0x6c, 0x6c, 0xac, 0x64, 0x11, 0x07, 0x31, 0x7a, 0x37, 0x2b,
	0xcb, 0x2c, 0xc7, 0x67, 0x64, 0x9a, 0x35, 0x8b, 0x67, 0x42, 0xb7, 0x3e, 0x22, 0x7a, 0xbc, 0x2e,
	0x61, 0x51, 0xd9, 0x20, 0x9e, 0xac, 0x8b, 0x56, 0x15, 0x58, 0x5b, 0x51, 0x54, 0xde, 0x30, 0xf9,
	0x7b, 0x00, 0x5b, 0x57, 0x35, 0x1a, 0xc6, 0x61, 0x07, 0xb5, 0x98, 0xe5, 0x28, 0xf9, 0x70, 0xdc,
	0x9b, 0x0e, 0x92, 0x00, 0xd9, 0x0f, 0x00, 0xc2, 0x5a, 0xa3, 0x66, 0x8d, 0xc5, 0x9a, 0xcb, 0xf1,
	0xe6, 0x74, 0x74, 0xf6, 0x69, 0xfc, 0xaf, 0xf7, 0x8c, 0x5d, 0xaa, 0xf8, 0x7c, 0xe5, 0x7e, 0xa1,
	0xad, 0x69, 0x93, 0x5b, 0xe1, 0xec, 0x00, 0x36, 0xeb, 0x66, 0xc6, 0x61, 0xdc, 0x9b, 0x0e, 0x13,
	0x77, 0x64, 0x0c, 0xb6, 0xb4, 0x28, 0x90, 0x8f, 0x88, 0xa2, 0x33, 0x7b, 0x1f, 0x20, 0x53, 0x37,
	0xa8, 0x53, 0x52, 0x1e, 0x92, 0x32, 0x24, 0xe6, 0xa5, 0x93, 0x4f, 0x60, 0xb4, 0x10, 0x85, 0xca,
	0x5b, 0xaf, 0xef, 0x91, 0x0e, 0x9e, 0x0a, 0x86, 0x42, 0x49, 0x99, 0xa3, 0x37, 0xec, 0x7b, 0x83,
	0xa7, 0xc8, 0x10, 0xc1, 0x40, 0xab, 0xf9, 0x35, 0xa9, 0x07, 0xa4, 0xae, 0x30, 0xfb, 0x0c, 0x58,
	0x65, 0x70, 0x81, 0xc6, 0xa0, 0x4c, 0x9b, 0x1a, 0x0d, 0xb9, 0x1e, 0x91, 0xeb, 0xd1, 0x4a, 0xb9,
	0xea, 0x04, 0xd7, 0xb8, 0xca, 0x94, 0x0b, 0x95, 0x23, 0x3f, 0x24, 0x4f, 0x80, 0xa4, 0xa8, 0xb9,
	0x6d, 0x0c, 0xf2, 0xa3, 0x4e, 0xf1, 0xd0, 0x29, 0xbf, 0xe3, 0xac, 0x56, 0x16, 0xf9, 0xb1, 0x57,
	0x3a, 0xc8, 0x0e, 0xa1, 0x8f, 0x85, 0x50, 0x39, 0xff, 0x80, 0x78, 0x0f, 0xd8, 0x87, 0xb0, 0x47,
	0x87, 0xf4, 0x06, 0x8d, 0x5a, 0x28, 0x94, 0xfc, 0x84, 0x66, 0xb4, 0x4b, 0xec, 0x2f, 0x1d, 0xc9,
	0x8e, 0x61, 0x3b, 0x43, 0x2d, 0xd1, 0xf0, 0x31, 0x45, 0x77, 0x88, 0xbd, 0x07, 0xc3, 0x99, 0x32,
	0x76, 0x29, 0x85, 0x45, 0xfe, 0xc4, 0x77, 0x73, 0x45, 0xb8, 0x5e, 0xfc, 0x51, 0x6a, 0x54, 0x7a,
	0x51, 0xf2, 0x89, 0xef, 0x45, 0xc0, 0x2e, 0x63, 0x5e, 0xce, 0x45, 0x8e, 0xfc, 0xa9, 0xcf, 0xe8,
	0x11, 0x7b, 0x02, 0x0f, 0xab, 0x65, 0xa9, 0x31, 0xd5, 0x4d, 0x31, 0x43, 0xc3, 0xa7, 0xa4, 0x8e,
	0x88, 0x7b, 0x49, 0x14, 0x3b, 0x83, 0xa3, 0xdb, 0x96, 0x37, 0x57, 0xff, 0x84, 0xae, 0xfe, 0xce,
	0x2d, 0xef, 0xaa, 0x80, 0x6f, 0x60, 0x47, 0x48, 0x69, 0xb0, 0xae, 0xf9, 0xe9, 0xb8, 0x37, 0x1d,
	0x9d, 0x3d, 0xbd, 0x73, 0xcf, 0xbc, 0x35, 0x09, 0x31, 0xec, 0x2b, 0x80, 0xb9, 0x41, 0x61, 0x51,
	0xa6, 0xc2, 0x72, 0xa4, 0x0c, 0x51, 0xec, 0x9f, 0x40, 0x1c, 0x9e, 0x40, 0xfc, 0x73, 0x78, 0x02,
	0xc9, 0xb0, 0x73, 0x9f, 0x5b, 0x17, 0xda, 0x54, 0x32, 0x84, 0x2e, 0xfe, 0x3f, 0xb4, 0x73, 0x9f,
	0xdb, 0xe8, 0x12, 0xf6, 0xd7, 0x36, 0xde, 0x6d, 0xf9, 0x35, 0xb6, 0xbc, 0xe7, 0xb7, 0xfc, 0x1a,
	0x5b, 0x76, 0x0a, 0xfd, 0x1b, 0x91, 0x37, 0xc8, 0x37, 0x28, 0xf5, 0xe1, 0x5b, 0xa9, 0xcf, 0x75,
	0x9b, 0x78, 0xcb, 0xd7, 0x1b, 0x5f, 0xf6, 0xa2, 0x3f, 0x7b, 0xb0, 0xd3, 0xd5, 0xe7, 0xc6, 0xb7,
	0x28, 0x4d, 0x21, 0xac, 0x45, 0xd9, 0xe5, 0x7c, 0x43, 0xb8, 0xdd, 0xa8, 0xad, 0x41, 0xb4, 0x69,
	0x68, 0xdd, 0x06, 0x59, 0x76, 0x3d, 0x1b, 0x92, 0x44, 0x30, 0xa0, 0xd9, 0x29, 0xdb, 0xf2, 0x4d,
	0x3f, 0xe5, 0x80, 0xdd, 0x94, 0x0d, 0x66, 0xaa, 0xd4, 0x7c, 0xcb, 0x4f, 0xd9, 0x23, 0xf7, 0x8c,
	0xaa, 0xb2, 0xb6, 0x22, 0x4f, 0xe7, 0xa5, 0x44, 0xde, 0xf7, 0xcf, 0xc8, 0x53, 0x17, 0xa5, 0xa4,
	0x3d, 0x9e, 0x97, 0x8d, 0x2b, 0x99, 0x6f, 0xfb, 0x3d, 0xee, 0xe0, 0xe4, 0x14, 0x0e, 0x7e, 0x54,
	0xb5, 0x75, 0x73, 0xaa, 0x13, 0x7c, 0xdd, 0x60, 0x6d, 0xdd, 0x67, 0x6a, 0x14, 0x66, 0xbe, 0xec,
	0x8a, 0xe8, 0xd0, 0x44, 0xc3, 0xee, 0x25, 0x9d, 0x82, 0xf1, 0x10, 0xfa, 0xaf, 0x1b, 0x34, 0xa1,
	0x81, 0x1e, 0xd0, 0x6d, 0x44, 0xb6, 0x5a, 0x39, 0x57, 0x65, 0x3f, 0x01, 0x47, 0x75, 0x1b, 0xf7,
	0x11, 0xec, 0x1b, 0xac, 0x9b, 0xdc, 0xa6, 0x15, 0x9a, 0xd4, 0x09, 0x54, 0x69, 0x3f, 0xd9, 0xf5,
	0xf4, 0x2b, 0x34, 0xaf, 0x44, 0x86, 0x67, 0x7f, 0x6d, 0x00, 0xfb, 0x5e, 0xa2, 0xb6, 0xca, 0xb6,
	0x3f, 0x09, 0x2d, 0x32, 0x2c, 0x50, 0x5b, 0x76, 0x01, 0x70, 0x41, 0xfb, 0x40, 0xff, 0xc3, 0xc7,
	0x77, 0x6c, 0x5e, 0x74, 0xfc, 0xd6, 0xf8, 0x5e, 0xb8, 0x9f, 0xee, 0xe4, 0x01, 0xfb, 0x0d, 0xf6,
	0x57, 0x75, 0x5f, 0x5a, 0x83, 0xa2, 0x60, 0x1f, 0xff, 0x47, 0xa6, 0xf5, 0xfe, 0x44, 0x77, 0x7d,
	0x72, 0xf2, 0xe0, 0xf3, 0x9e, 0xbb, 0xe1, 0x15, 0xad, 0xdd, 0x7d, 0x6e, 0x78, 0x01, 0xf0, 0x1d,
	0xe6, 0x78, 0xaf, 0x24, 0xdf, 0xc2, 0xaf, 0x83, 0x60, 0x9d, 0x6d, 0x93, 0xfa, 0xc5, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd0, 0xa7, 0xf7, 0x5b, 0xcb, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityManagementClient is the client API for IdentityManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityManagementClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	ListUsersStream(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (IdentityManagement_ListUsersStreamClient, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
}

type identityManagementClient struct {
	cc *grpc.ClientConn
}

func NewIdentityManagementClient(cc *grpc.ClientConn) IdentityManagementClient {
	return &identityManagementClient{cc}
}

func (c *identityManagementClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nghinhut.idm.v1alpha1.IdentityManagement/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) ListUsersStream(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (IdentityManagement_ListUsersStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IdentityManagement_serviceDesc.Streams[0], "/nghinhut.idm.v1alpha1.IdentityManagement/ListUsersStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &identityManagementListUsersStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IdentityManagement_ListUsersStreamClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type identityManagementListUsersStreamClient struct {
	grpc.ClientStream
}

func (x *identityManagementListUsersStreamClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *identityManagementClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nghinhut.idm.v1alpha1.IdentityManagement/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nghinhut.idm.v1alpha1.IdentityManagement/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityManagementServer is the server API for IdentityManagement service.
type IdentityManagementServer interface {
	CreateUser(context.Context, *User) (*empty.Empty, error)
	ListUsersStream(*ListUsersRequest, IdentityManagement_ListUsersStreamServer) error
	UpdateUser(context.Context, *User) (*empty.Empty, error)
	DeleteUser(context.Context, *User) (*empty.Empty, error)
}

// UnimplementedIdentityManagementServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityManagementServer struct {
}

func (*UnimplementedIdentityManagementServer) CreateUser(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedIdentityManagementServer) ListUsersStream(req *ListUsersRequest, srv IdentityManagement_ListUsersStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsersStream not implemented")
}
func (*UnimplementedIdentityManagementServer) UpdateUser(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedIdentityManagementServer) DeleteUser(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterIdentityManagementServer(s *grpc.Server, srv IdentityManagementServer) {
	s.RegisterService(&_IdentityManagement_serviceDesc, srv)
}

func _IdentityManagement_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nghinhut.idm.v1alpha1.IdentityManagement/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_ListUsersStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IdentityManagementServer).ListUsersStream(m, &identityManagementListUsersStreamServer{stream})
}

type IdentityManagement_ListUsersStreamServer interface {
	Send(*User) error
	grpc.ServerStream
}

type identityManagementListUsersStreamServer struct {
	grpc.ServerStream
}

func (x *identityManagementListUsersStreamServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _IdentityManagement_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nghinhut.idm.v1alpha1.IdentityManagement/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nghinhut.idm.v1alpha1.IdentityManagement/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nghinhut.idm.v1alpha1.IdentityManagement",
	HandlerType: (*IdentityManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _IdentityManagement_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _IdentityManagement_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _IdentityManagement_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsersStream",
			Handler:       _IdentityManagement_ListUsersStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1alpha1.proto",
}
