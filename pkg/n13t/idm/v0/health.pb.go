// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n13t/idm/v0/health.proto

package n13t_idm_v0

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d95ddfda6757aa4, []int{1, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d95ddfda6757aa4, []int{0}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=n13t.idm.v0.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d95ddfda6757aa4, []int{1}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("n13t.idm.v0.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterType((*HealthCheckRequest)(nil), "n13t.idm.v0.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "n13t.idm.v0.HealthCheckResponse")
}

func init() {
	proto.RegisterFile("n13t/idm/v0/health.proto", fileDescriptor_8d95ddfda6757aa4)
}

var fileDescriptor_8d95ddfda6757aa4 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x33, 0x34, 0x2e,
	0xd1, 0xcf, 0x4c, 0xc9, 0xd5, 0x2f, 0x33, 0xd0, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0xc9, 0xe8, 0x65, 0xa6, 0xe4, 0xea, 0x95, 0x19, 0x28,
	0xe9, 0x71, 0x09, 0x79, 0x80, 0x25, 0x9d, 0x33, 0x52, 0x93, 0xb3, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0xa5, 0x99, 0x8c, 0x5c, 0xc2, 0x28, 0x1a, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x85, 0xdc, 0xb8, 0xd8, 0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0xc1, 0x1a, 0xf8,
	0x8c, 0xf4, 0xf4, 0x90, 0x6c, 0xd1, 0xc3, 0xa2, 0x43, 0x2f, 0x18, 0x64, 0x62, 0x5e, 0x7a, 0x30,
	0x58, 0x57, 0x10, 0x54, 0xb7, 0x92, 0x15, 0x17, 0x2f, 0x8a, 0x84, 0x10, 0x37, 0x17, 0x7b, 0xa8,
	0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x03, 0x88, 0x13, 0xec, 0x1a, 0x14, 0xe6, 0xe9, 0xe7,
	0x2e, 0xc0, 0x28, 0xc4, 0xcf, 0xc5, 0xed, 0xe7, 0x1f, 0x12, 0x0f, 0x13, 0x60, 0x4a, 0x62, 0x03,
	0xfb, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x79, 0xea, 0x4e, 0xd1, 0xfb, 0x00, 0x00, 0x00,
}