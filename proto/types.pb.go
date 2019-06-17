// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Message to store bundle/config.json bytes
type ExtraData struct {
	JsonSpec             []byte     `protobuf:"bytes,1,opt,name=JsonSpec,proto3" json:"JsonSpec,omitempty"`
	RuncOptions          *types.Any `protobuf:"bytes,2,opt,name=RuncOptions,proto3" json:"RuncOptions,omitempty"`
	StdinPort            uint32     `protobuf:"varint,3,opt,name=StdinPort,proto3" json:"StdinPort,omitempty"`
	StdoutPort           uint32     `protobuf:"varint,4,opt,name=StdoutPort,proto3" json:"StdoutPort,omitempty"`
	StderrPort           uint32     `protobuf:"varint,5,opt,name=StderrPort,proto3" json:"StderrPort,omitempty"`
	DriveID              string     `protobuf:"bytes,6,opt,name=DriveID,proto3" json:"DriveID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ExtraData) Reset()         { *m = ExtraData{} }
func (m *ExtraData) String() string { return proto.CompactTextString(m) }
func (*ExtraData) ProtoMessage()    {}
func (*ExtraData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}
func (m *ExtraData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraData.Unmarshal(m, b)
}
func (m *ExtraData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraData.Marshal(b, m, deterministic)
}
func (m *ExtraData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraData.Merge(m, src)
}
func (m *ExtraData) XXX_Size() int {
	return xxx_messageInfo_ExtraData.Size(m)
}
func (m *ExtraData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraData.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraData proto.InternalMessageInfo

func (m *ExtraData) GetJsonSpec() []byte {
	if m != nil {
		return m.JsonSpec
	}
	return nil
}

func (m *ExtraData) GetRuncOptions() *types.Any {
	if m != nil {
		return m.RuncOptions
	}
	return nil
}

func (m *ExtraData) GetStdinPort() uint32 {
	if m != nil {
		return m.StdinPort
	}
	return 0
}

func (m *ExtraData) GetStdoutPort() uint32 {
	if m != nil {
		return m.StdoutPort
	}
	return 0
}

func (m *ExtraData) GetStderrPort() uint32 {
	if m != nil {
		return m.StderrPort
	}
	return 0
}

func (m *ExtraData) GetDriveID() string {
	if m != nil {
		return m.DriveID
	}
	return ""
}

// Message to specify network config for a Firecracker VM
type FirecrackerNetworkInterface struct {
	MacAddress           string                  `protobuf:"bytes,1,opt,name=MacAddress,proto3" json:"MacAddress,omitempty"`
	HostDevName          string                  `protobuf:"bytes,2,opt,name=HostDevName,proto3" json:"HostDevName,omitempty"`
	AllowMMDS            bool                    `protobuf:"varint,3,opt,name=AllowMMDS,proto3" json:"AllowMMDS,omitempty"`
	InRateLimiter        *FirecrackerRateLimiter `protobuf:"bytes,4,opt,name=InRateLimiter,proto3" json:"InRateLimiter,omitempty"`
	OutRateLimiter       *FirecrackerRateLimiter `protobuf:"bytes,5,opt,name=OutRateLimiter,proto3" json:"OutRateLimiter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FirecrackerNetworkInterface) Reset()         { *m = FirecrackerNetworkInterface{} }
func (m *FirecrackerNetworkInterface) String() string { return proto.CompactTextString(m) }
func (*FirecrackerNetworkInterface) ProtoMessage()    {}
func (*FirecrackerNetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}
func (m *FirecrackerNetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirecrackerNetworkInterface.Unmarshal(m, b)
}
func (m *FirecrackerNetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirecrackerNetworkInterface.Marshal(b, m, deterministic)
}
func (m *FirecrackerNetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirecrackerNetworkInterface.Merge(m, src)
}
func (m *FirecrackerNetworkInterface) XXX_Size() int {
	return xxx_messageInfo_FirecrackerNetworkInterface.Size(m)
}
func (m *FirecrackerNetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_FirecrackerNetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_FirecrackerNetworkInterface proto.InternalMessageInfo

func (m *FirecrackerNetworkInterface) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *FirecrackerNetworkInterface) GetHostDevName() string {
	if m != nil {
		return m.HostDevName
	}
	return ""
}

func (m *FirecrackerNetworkInterface) GetAllowMMDS() bool {
	if m != nil {
		return m.AllowMMDS
	}
	return false
}

func (m *FirecrackerNetworkInterface) GetInRateLimiter() *FirecrackerRateLimiter {
	if m != nil {
		return m.InRateLimiter
	}
	return nil
}

func (m *FirecrackerNetworkInterface) GetOutRateLimiter() *FirecrackerRateLimiter {
	if m != nil {
		return m.OutRateLimiter
	}
	return nil
}

// Message to set the machine config for a Firecracker VM
type FirecrackerMachineConfiguration struct {
	CPUTemplate string `protobuf:"bytes,1,opt,name=CPUTemplate,proto3" json:"CPUTemplate,omitempty"`
	HtEnabled   bool   `protobuf:"varint,2,opt,name=HtEnabled,proto3" json:"HtEnabled,omitempty"`
	// Specifies the memory size of VM
	// This lets us create a Firecracker VM of up to 4096 TiB, which
	// for a mircroVM should be large enough
	MemSizeMib           uint32   `protobuf:"varint,3,opt,name=MemSizeMib,proto3" json:"MemSizeMib,omitempty"`
	VcpuCount            uint32   `protobuf:"varint,4,opt,name=VcpuCount,proto3" json:"VcpuCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FirecrackerMachineConfiguration) Reset()         { *m = FirecrackerMachineConfiguration{} }
func (m *FirecrackerMachineConfiguration) String() string { return proto.CompactTextString(m) }
func (*FirecrackerMachineConfiguration) ProtoMessage()    {}
func (*FirecrackerMachineConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}
func (m *FirecrackerMachineConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirecrackerMachineConfiguration.Unmarshal(m, b)
}
func (m *FirecrackerMachineConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirecrackerMachineConfiguration.Marshal(b, m, deterministic)
}
func (m *FirecrackerMachineConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirecrackerMachineConfiguration.Merge(m, src)
}
func (m *FirecrackerMachineConfiguration) XXX_Size() int {
	return xxx_messageInfo_FirecrackerMachineConfiguration.Size(m)
}
func (m *FirecrackerMachineConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_FirecrackerMachineConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_FirecrackerMachineConfiguration proto.InternalMessageInfo

func (m *FirecrackerMachineConfiguration) GetCPUTemplate() string {
	if m != nil {
		return m.CPUTemplate
	}
	return ""
}

func (m *FirecrackerMachineConfiguration) GetHtEnabled() bool {
	if m != nil {
		return m.HtEnabled
	}
	return false
}

func (m *FirecrackerMachineConfiguration) GetMemSizeMib() uint32 {
	if m != nil {
		return m.MemSizeMib
	}
	return 0
}

func (m *FirecrackerMachineConfiguration) GetVcpuCount() uint32 {
	if m != nil {
		return m.VcpuCount
	}
	return 0
}

// Message to specify the block device config for a Firecracker VM
type FirecrackerDrive struct {
	IsReadOnly           bool                    `protobuf:"varint,1,opt,name=IsReadOnly,proto3" json:"IsReadOnly,omitempty"`
	IsRootDevice         bool                    `protobuf:"varint,2,opt,name=IsRootDevice,proto3" json:"IsRootDevice,omitempty"`
	Partuuid             string                  `protobuf:"bytes,3,opt,name=Partuuid,proto3" json:"Partuuid,omitempty"`
	PathOnHost           string                  `protobuf:"bytes,4,opt,name=PathOnHost,proto3" json:"PathOnHost,omitempty"`
	RateLimiter          *FirecrackerRateLimiter `protobuf:"bytes,5,opt,name=RateLimiter,proto3" json:"RateLimiter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FirecrackerDrive) Reset()         { *m = FirecrackerDrive{} }
func (m *FirecrackerDrive) String() string { return proto.CompactTextString(m) }
func (*FirecrackerDrive) ProtoMessage()    {}
func (*FirecrackerDrive) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}
func (m *FirecrackerDrive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirecrackerDrive.Unmarshal(m, b)
}
func (m *FirecrackerDrive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirecrackerDrive.Marshal(b, m, deterministic)
}
func (m *FirecrackerDrive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirecrackerDrive.Merge(m, src)
}
func (m *FirecrackerDrive) XXX_Size() int {
	return xxx_messageInfo_FirecrackerDrive.Size(m)
}
func (m *FirecrackerDrive) XXX_DiscardUnknown() {
	xxx_messageInfo_FirecrackerDrive.DiscardUnknown(m)
}

var xxx_messageInfo_FirecrackerDrive proto.InternalMessageInfo

func (m *FirecrackerDrive) GetIsReadOnly() bool {
	if m != nil {
		return m.IsReadOnly
	}
	return false
}

func (m *FirecrackerDrive) GetIsRootDevice() bool {
	if m != nil {
		return m.IsRootDevice
	}
	return false
}

func (m *FirecrackerDrive) GetPartuuid() string {
	if m != nil {
		return m.Partuuid
	}
	return ""
}

func (m *FirecrackerDrive) GetPathOnHost() string {
	if m != nil {
		return m.PathOnHost
	}
	return ""
}

func (m *FirecrackerDrive) GetRateLimiter() *FirecrackerRateLimiter {
	if m != nil {
		return m.RateLimiter
	}
	return nil
}

// Message to specify an IO rate limiter with bytes/s and ops/s limits
type FirecrackerRateLimiter struct {
	Bandwidth            *FirecrackerTokenBucket `protobuf:"bytes,1,opt,name=Bandwidth,proto3" json:"Bandwidth,omitempty"`
	Ops                  *FirecrackerTokenBucket `protobuf:"bytes,2,opt,name=Ops,proto3" json:"Ops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FirecrackerRateLimiter) Reset()         { *m = FirecrackerRateLimiter{} }
func (m *FirecrackerRateLimiter) String() string { return proto.CompactTextString(m) }
func (*FirecrackerRateLimiter) ProtoMessage()    {}
func (*FirecrackerRateLimiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}
func (m *FirecrackerRateLimiter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirecrackerRateLimiter.Unmarshal(m, b)
}
func (m *FirecrackerRateLimiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirecrackerRateLimiter.Marshal(b, m, deterministic)
}
func (m *FirecrackerRateLimiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirecrackerRateLimiter.Merge(m, src)
}
func (m *FirecrackerRateLimiter) XXX_Size() int {
	return xxx_messageInfo_FirecrackerRateLimiter.Size(m)
}
func (m *FirecrackerRateLimiter) XXX_DiscardUnknown() {
	xxx_messageInfo_FirecrackerRateLimiter.DiscardUnknown(m)
}

var xxx_messageInfo_FirecrackerRateLimiter proto.InternalMessageInfo

func (m *FirecrackerRateLimiter) GetBandwidth() *FirecrackerTokenBucket {
	if m != nil {
		return m.Bandwidth
	}
	return nil
}

func (m *FirecrackerRateLimiter) GetOps() *FirecrackerTokenBucket {
	if m != nil {
		return m.Ops
	}
	return nil
}

// Message to specify a token buicket used to rate limit disk and network IO for a Firecracker VM
type FirecrackerTokenBucket struct {
	OneTimeBurst         int64    `protobuf:"varint,1,opt,name=OneTimeBurst,proto3" json:"OneTimeBurst,omitempty"`
	RefillTime           int64    `protobuf:"varint,2,opt,name=RefillTime,proto3" json:"RefillTime,omitempty"`
	Capacity             int64    `protobuf:"varint,3,opt,name=Capacity,proto3" json:"Capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FirecrackerTokenBucket) Reset()         { *m = FirecrackerTokenBucket{} }
func (m *FirecrackerTokenBucket) String() string { return proto.CompactTextString(m) }
func (*FirecrackerTokenBucket) ProtoMessage()    {}
func (*FirecrackerTokenBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{5}
}
func (m *FirecrackerTokenBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirecrackerTokenBucket.Unmarshal(m, b)
}
func (m *FirecrackerTokenBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirecrackerTokenBucket.Marshal(b, m, deterministic)
}
func (m *FirecrackerTokenBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirecrackerTokenBucket.Merge(m, src)
}
func (m *FirecrackerTokenBucket) XXX_Size() int {
	return xxx_messageInfo_FirecrackerTokenBucket.Size(m)
}
func (m *FirecrackerTokenBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_FirecrackerTokenBucket.DiscardUnknown(m)
}

var xxx_messageInfo_FirecrackerTokenBucket proto.InternalMessageInfo

func (m *FirecrackerTokenBucket) GetOneTimeBurst() int64 {
	if m != nil {
		return m.OneTimeBurst
	}
	return 0
}

func (m *FirecrackerTokenBucket) GetRefillTime() int64 {
	if m != nil {
		return m.RefillTime
	}
	return 0
}

func (m *FirecrackerTokenBucket) GetCapacity() int64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func init() {
	proto.RegisterType((*ExtraData)(nil), "ExtraData")
	proto.RegisterType((*FirecrackerNetworkInterface)(nil), "FirecrackerNetworkInterface")
	proto.RegisterType((*FirecrackerMachineConfiguration)(nil), "FirecrackerMachineConfiguration")
	proto.RegisterType((*FirecrackerDrive)(nil), "FirecrackerDrive")
	proto.RegisterType((*FirecrackerRateLimiter)(nil), "FirecrackerRateLimiter")
	proto.RegisterType((*FirecrackerTokenBucket)(nil), "FirecrackerTokenBucket")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xc7, 0x59, 0x63, 0xdb, 0xec, 0x6c, 0x2b, 0x32, 0x88, 0xc6, 0x2a, 0x1a, 0xf6, 0x2a, 0xde,
	0xa4, 0x50, 0x51, 0xf0, 0x42, 0xa4, 0x49, 0x2a, 0x8d, 0x98, 0x6e, 0x98, 0x54, 0x2f, 0xbc, 0x9b,
	0xec, 0x9e, 0x24, 0x43, 0x36, 0x33, 0xcb, 0xec, 0x99, 0xa6, 0xe9, 0xcb, 0xf8, 0x44, 0x5e, 0xfb,
	0x34, 0x82, 0xcc, 0xe4, 0x6b, 0x52, 0x54, 0xf0, 0x2a, 0x9c, 0xdf, 0xf9, 0x98, 0xf3, 0xff, 0x9f,
	0x2c, 0x89, 0x70, 0x51, 0x40, 0xd9, 0x2c, 0xb4, 0x42, 0x75, 0xfc, 0x74, 0xac, 0xd4, 0x38, 0x87,
	0x13, 0x17, 0x0d, 0xcd, 0xe8, 0x84, 0xcb, 0xc5, 0x32, 0x15, 0xff, 0x0c, 0x48, 0x78, 0x7e, 0x83,
	0x9a, 0x77, 0x38, 0x72, 0x7a, 0x4c, 0xaa, 0x9f, 0x4a, 0x25, 0x07, 0x05, 0xa4, 0xb5, 0xa0, 0x1e,
	0x34, 0x0e, 0xd9, 0x26, 0xa6, 0x6f, 0x49, 0xc4, 0x8c, 0x4c, 0x93, 0x02, 0x85, 0x92, 0x65, 0xed,
	0x5e, 0x3d, 0x68, 0x44, 0xa7, 0x8f, 0x9a, 0xcb, 0xd1, 0xcd, 0xf5, 0xe8, 0xe6, 0x99, 0x5c, 0x30,
	0xbf, 0x90, 0x3e, 0x27, 0xe1, 0x00, 0x33, 0x21, 0xfb, 0x4a, 0x63, 0xad, 0x52, 0x0f, 0x1a, 0x47,
	0x6c, 0x0b, 0xe8, 0x0b, 0x42, 0x06, 0x98, 0x29, 0x83, 0x2e, 0x7d, 0xdf, 0xa5, 0x3d, 0xb2, 0xca,
	0x83, 0xd6, 0x2e, 0xbf, 0xb7, 0xc9, 0xaf, 0x08, 0xad, 0x91, 0x83, 0x8e, 0x16, 0xd7, 0xd0, 0xed,
	0xd4, 0xf6, 0xeb, 0x41, 0x23, 0x64, 0xeb, 0x30, 0xfe, 0x15, 0x90, 0x67, 0x1f, 0x85, 0x86, 0x54,
	0xf3, 0x74, 0x0a, 0xfa, 0x12, 0x70, 0xae, 0xf4, 0xb4, 0x2b, 0x11, 0xf4, 0x88, 0xa7, 0x60, 0x27,
	0xf7, 0x78, 0x7a, 0x96, 0x65, 0x1a, 0xca, 0xd2, 0xa9, 0x0d, 0x99, 0x47, 0x68, 0x9d, 0x44, 0x17,
	0xaa, 0xc4, 0x0e, 0x5c, 0x5f, 0xf2, 0x19, 0x38, 0xbd, 0x21, 0xf3, 0x91, 0x55, 0x76, 0x96, 0xe7,
	0x6a, 0xde, 0xeb, 0x75, 0x06, 0x4e, 0x59, 0x95, 0x6d, 0x01, 0x7d, 0x4f, 0x8e, 0xba, 0x92, 0x71,
	0x84, 0xcf, 0x62, 0x26, 0x10, 0xb4, 0x13, 0x17, 0x9d, 0x3e, 0x69, 0x7a, 0x4b, 0x79, 0x69, 0xb6,
	0x5b, 0x4d, 0x3f, 0x90, 0x07, 0x89, 0x41, 0xbf, 0x7f, 0xef, 0xdf, 0xfd, 0x77, 0xca, 0xe3, 0xef,
	0x01, 0x79, 0xe9, 0x95, 0xf6, 0x78, 0x3a, 0x11, 0x12, 0xda, 0x4a, 0x8e, 0xc4, 0xd8, 0x68, 0x6e,
	0x8f, 0x63, 0x35, 0xb6, 0xfb, 0x5f, 0xae, 0x60, 0x56, 0xe4, 0x1c, 0x61, 0x65, 0x82, 0x8f, 0xac,
	0xc6, 0x0b, 0x3c, 0x97, 0x7c, 0x98, 0x43, 0xe6, 0x3c, 0xa8, 0xb2, 0x2d, 0x70, 0x1e, 0xc2, 0x6c,
	0x20, 0x6e, 0xa1, 0x27, 0x86, 0xab, 0xe3, 0x7a, 0xc4, 0x76, 0x7f, 0x4d, 0x0b, 0xd3, 0x56, 0x46,
	0xae, 0x8f, 0xbb, 0x05, 0xf1, 0x8f, 0x80, 0x3c, 0xf4, 0x36, 0x74, 0x87, 0xb3, 0x23, 0xbb, 0x25,
	0x03, 0x9e, 0x25, 0x32, 0x5f, 0xb8, 0x8d, 0xaa, 0xcc, 0x23, 0x34, 0x26, 0x87, 0xdd, 0x92, 0x29,
	0x65, 0xaf, 0x20, 0x52, 0x58, 0xed, 0xb4, 0xc3, 0xec, 0xdf, 0xb8, 0xcf, 0x35, 0x1a, 0x23, 0x32,
	0xb7, 0x54, 0xc8, 0x36, 0xb1, 0x9d, 0xdf, 0xe7, 0x38, 0x49, 0xa4, 0xbd, 0xa4, 0xdb, 0x29, 0x64,
	0x1e, 0xa1, 0xef, 0x48, 0xf4, 0x1f, 0xa6, 0xfb, 0xb5, 0xf1, 0x2d, 0x79, 0xfc, 0xe7, 0x32, 0xfa,
	0x86, 0x84, 0x2d, 0x2e, 0xb3, 0xb9, 0xc8, 0x70, 0xe2, 0x34, 0xdd, 0x19, 0x79, 0xa5, 0xa6, 0x20,
	0x5b, 0x26, 0x9d, 0x02, 0xb2, 0x6d, 0x25, 0x7d, 0x45, 0x2a, 0x49, 0xb1, 0xfe, 0xd4, 0xfe, 0xda,
	0x60, 0x6b, 0xe2, 0x9b, 0x9d, 0xb7, 0xbd, 0xb4, 0x35, 0x2c, 0x91, 0x70, 0x25, 0x66, 0xd0, 0x32,
	0xba, 0x44, 0xf7, 0x7c, 0x85, 0xed, 0x30, 0x6b, 0x0a, 0x83, 0x91, 0xc8, 0x73, 0x8b, 0xdc, 0x7b,
	0x15, 0xe6, 0x11, 0x6b, 0x68, 0x9b, 0x17, 0x3c, 0x15, 0xb8, 0x70, 0x86, 0x56, 0xd8, 0x26, 0x6e,
	0x1d, 0x7c, 0xdb, 0x5b, 0x7e, 0xfc, 0xfb, 0xee, 0xe7, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x21, 0x20, 0xfb, 0x58, 0x7b, 0x04, 0x00, 0x00,
}
