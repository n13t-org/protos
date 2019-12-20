// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package n13t_idm_v0

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//*
// Represents the user
//
// Compatible with OIDC (OpenID Connect)
type User struct {
	Enabled bool `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
	//    google.protobuf.Timestamp created_at = 101;     // ???
	Attributes map[string]*any.Any `protobuf:"bytes,100,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// OpenId Standard Claims Members (https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims)
	Sub               string `protobuf:"bytes,10,opt,name=sub,proto3" json:"sub,omitempty"`
	Name              string `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	GivenName         string `protobuf:"bytes,12,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName        string `protobuf:"bytes,14,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	MiddleName        string `protobuf:"bytes,15,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	Nickname          string `protobuf:"bytes,16,opt,name=nickname,proto3" json:"nickname,omitempty"`
	PreferredUsername string `protobuf:"bytes,17,opt,name=preferred_username,json=preferredUsername,proto3" json:"preferred_username,omitempty"`
	Profile           string `protobuf:"bytes,20,opt,name=profile,proto3" json:"profile,omitempty"`
	Picture           string `protobuf:"bytes,21,opt,name=picture,proto3" json:"picture,omitempty"`
	Website           string `protobuf:"bytes,22,opt,name=website,proto3" json:"website,omitempty"`
	Email             string `protobuf:"bytes,30,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified     bool   `protobuf:"varint,31,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Gender            string `protobuf:"bytes,32,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthdate         string `protobuf:"bytes,33,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Zoneinfo          string `protobuf:"bytes,34,opt,name=zoneinfo,proto3" json:"zoneinfo,omitempty"`
	Locale            string `protobuf:"bytes,35,opt,name=locale,proto3" json:"locale,omitempty"`
	//*
	//  Ex: +1 (425) 555-1212 or +56 (2) 687 2400   (E.164)
	//   or +1 (604) 555-1234;ext=5678              (RFC3966)
	//  see E.164/RFC3966
	PhoneNumber          string               `protobuf:"bytes,40,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	PhoneNumberVerified  bool                 `protobuf:"varint,41,opt,name=phone_number_verified,json=phoneNumberVerified,proto3" json:"phone_number_verified,omitempty"`
	Address              *User_Address        `protobuf:"bytes,42,opt,name=address,proto3" json:"address,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,102,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
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
	return fileDescriptor_116e343673f7ffaf, []int{0, 1}
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

func init() {
	proto.RegisterType((*User)(nil), "n13t.idm.v0.User")
	proto.RegisterMapType((map[string]*any.Any)(nil), "n13t.idm.v0.User.AttributesEntry")
	proto.RegisterType((*User_Address)(nil), "n13t.idm.v0.User.Address")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x5b, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0x77, 0x6d, 0x4f, 0x77, 0x35, 0xdb, 0xe4, 0x85, 0xcb, 0xba, 0x21, 0xa4, 0x32, 0x89,
	0x0c, 0xb6, 0x17, 0xe0, 0xad, 0x42, 0xbc, 0xee, 0xa1, 0x30, 0x5e, 0x23, 0xa7, 0x3e, 0xe9, 0xac,
	0x25, 0x76, 0xe4, 0x38, 0x45, 0xe1, 0x07, 0xf0, 0xaf, 0xf8, 0x6f, 0xc8, 0xc7, 0x49, 0x57, 0x0d,
	0xde, 0xfc, 0x5d, 0x7c, 0xe4, 0x73, 0xce, 0x67, 0x80, 0xba, 0x42, 0x1b, 0x97, 0xd6, 0x38, 0xc3,
	0x86, 0xfa, 0xc3, 0x8d, 0x8b, 0x95, 0x2c, 0xe2, 0xc5, 0xfb, 0xe8, 0x74, 0x6e, 0xcc, 0x3c, 0xc7,
	0x2b, 0x92, 0xd2, 0x3a, 0xbb, 0x12, 0xba, 0x09, 0xbe, 0xe8, 0xf9, 0x53, 0x09, 0x8b, 0xd2, 0x75,
	0xe2, 0xd9, 0x53, 0xd1, 0xa9, 0x02, 0x2b, 0x27, 0x8a, 0x32, 0x18, 0x2e, 0x7e, 0xf7, 0x61, 0xe3,
	0xae, 0x42, 0xcb, 0x38, 0x6c, 0xa3, 0x16, 0x69, 0x8e, 0x92, 0x0f, 0x46, 0xbd, 0x71, 0x7f, 0xda,
	0x41, 0x36, 0x01, 0x10, 0xce, 0x59, 0x95, 0xd6, 0x0e, 0x2b, 0x2e, 0x47, 0xeb, 0xe3, 0xe1, 0xf5,
	0x79, 0xbc, 0xf2, 0xba, 0xd8, 0x17, 0x88, 0x27, 0x4b, 0xcf, 0x57, 0xed, 0x6c, 0x33, 0x5d, 0xb9,
	0xc4, 0x0e, 0x60, 0xbd, 0xaa, 0x53, 0x0e, 0xa3, 0xde, 0x78, 0x30, 0xf5, 0x47, 0xc6, 0x60, 0x43,
	0x8b, 0x02, 0xf9, 0x90, 0x28, 0x3a, 0xb3, 0x97, 0x00, 0x73, 0xb5, 0x40, 0x9d, 0x90, 0xb2, 0x43,
	0xca, 0x80, 0x98, 0x5b, 0x2f, 0x9f, 0xc1, 0x30, 0x13, 0x85, 0xca, 0x9b, 0xa0, 0xef, 0x91, 0x0e,
	0x81, 0xea, 0x0c, 0x85, 0x92, 0x32, 0xc7, 0x60, 0xd8, 0x0f, 0x86, 0x40, 0x91, 0x21, 0x82, 0xbe,
	0x56, 0xb3, 0x07, 0x52, 0x0f, 0x48, 0x5d, 0x62, 0xf6, 0x0e, 0x58, 0x69, 0x31, 0x43, 0x6b, 0x51,
	0x26, 0x7e, 0x0d, 0xe4, 0x3a, 0x24, 0xd7, 0xe1, 0x52, 0xb9, 0x6b, 0x05, 0x3f, 0xae, 0xd2, 0x9a,
	0x4c, 0xe5, 0xc8, 0x8f, 0xc8, 0xd3, 0x41, 0x52, 0xd4, 0xcc, 0xd5, 0x16, 0xf9, 0x71, 0xab, 0x04,
	0xe8, 0x95, 0x9f, 0x98, 0x56, 0xca, 0x21, 0x3f, 0x09, 0x4a, 0x0b, 0xd9, 0x11, 0x6c, 0x62, 0x21,
	0x54, 0xce, 0x5f, 0x11, 0x1f, 0x00, 0x7b, 0x03, 0x7b, 0x74, 0x48, 0x16, 0x68, 0x55, 0xa6, 0x50,
	0xf2, 0x33, 0xda, 0xcc, 0x2e, 0xb1, 0x3f, 0x5a, 0x92, 0x9d, 0xc0, 0xd6, 0x1c, 0xb5, 0x44, 0xcb,
	0x47, 0x74, 0xbb, 0x45, 0xec, 0x05, 0x0c, 0x52, 0x65, 0xdd, 0xbd, 0x14, 0x0e, 0xf9, 0x79, 0x98,
	0xe6, 0x92, 0xf0, 0xb3, 0xf8, 0x65, 0x34, 0x2a, 0x9d, 0x19, 0x7e, 0x11, 0x66, 0xd1, 0x61, 0x5f,
	0x31, 0x37, 0x33, 0x91, 0x23, 0x7f, 0x1d, 0x2a, 0x06, 0xc4, 0xce, 0x61, 0xa7, 0xbc, 0x37, 0x1a,
	0x13, 0x5d, 0x17, 0x29, 0x5a, 0x3e, 0x26, 0x75, 0x48, 0xdc, 0x2d, 0x51, 0xec, 0x1a, 0x8e, 0x57,
	0x2d, 0x8f, 0x4f, 0x7f, 0x4b, 0x4f, 0x7f, 0xb6, 0xe2, 0x5d, 0x36, 0x70, 0x03, 0xdb, 0x42, 0x4a,
	0x8b, 0x55, 0xc5, 0x2f, 0x47, 0xbd, 0xf1, 0xf0, 0xfa, 0xf4, 0x3f, 0xe9, 0x0a, 0x86, 0x69, 0xe7,
	0x64, 0x9f, 0x00, 0xea, 0xd2, 0x77, 0x22, 0x13, 0xe1, 0x78, 0x46, 0xf7, 0xa2, 0x38, 0xc4, 0x3d,
	0xee, 0xe2, 0x1e, 0x7f, 0xef, 0xe2, 0x3e, 0x1d, 0xb4, 0xee, 0x89, 0x8b, 0xbe, 0xc1, 0xfe, 0x93,
	0xb0, 0xfa, 0x80, 0x3e, 0x60, 0xc3, 0x7b, 0x21, 0xa0, 0x0f, 0xd8, 0xb0, 0x4b, 0xd8, 0x5c, 0x88,
	0xbc, 0x46, 0xbe, 0x46, 0xa5, 0x8f, 0xfe, 0x29, 0x3d, 0xd1, 0xcd, 0x34, 0x58, 0x3e, 0xaf, 0x7d,
	0xec, 0x45, 0x7f, 0x7a, 0xb0, 0xdd, 0x3e, 0xd2, 0x4f, 0x3e, 0x33, 0xb6, 0x10, 0xce, 0xa1, 0x6c,
	0x6b, 0x3e, 0x12, 0x7e, 0xad, 0x95, 0xb3, 0x88, 0x2e, 0xe9, 0xba, 0x5e, 0x23, 0xcb, 0x6e, 0x60,
	0xbb, 0x22, 0x11, 0xf4, 0x69, 0xec, 0xca, 0x35, 0x7c, 0x3d, 0x2c, 0xa8, 0xc3, 0x7e, 0x41, 0x16,
	0xe7, 0xca, 0x68, 0xbe, 0x11, 0x16, 0x14, 0x90, 0xff, 0x01, 0xa5, 0xa9, 0x9c, 0xc8, 0x93, 0x99,
	0x91, 0xc8, 0x37, 0xc3, 0x0f, 0x08, 0xd4, 0x17, 0x23, 0x29, 0x82, 0x33, 0x53, 0xfb, 0x96, 0xf9,
	0x56, 0x88, 0x60, 0x0b, 0xd3, 0x2d, 0x6a, 0xec, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0,
	0x2e, 0xcf, 0x7c, 0x83, 0x04, 0x00, 0x00,
}
