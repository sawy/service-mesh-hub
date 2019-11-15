// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/v1/rbac.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// if RBAC policies have been specified, service isolation will not work as expected.
// Istio, the model for this feature, has an exclude-by-default behavior. If RBAC is enabled, services can only
// do what the RBAC policy allows them to do. This means that in the special case of
type RbacStatusCode int32

const (
	// Initial value, will be replaced as soon as operator evaluates the config
	RbacStatusCode_PENDING_VERIFICATION RbacStatusCode = 0
	// Config was applied successfully
	RbacStatusCode_OK RbacStatusCode = 1
	// If a mesh does not support the provided configuration, this error code is returned.
	RbacStatusCode_ERROR_RBAC_MODE_NOT_SUPPORTED_BY_MESH RbacStatusCode = 2
	// If other policies exist, we lose the ability to have total ON vs. OFF control of RBAC. Isolation is not supported
	// in this case.
	RbacStatusCode_ERROR_CANNOT_ISOLATE_RBAC_SINCE_POLICIES_EXIST RbacStatusCode = 3
	// If the config is not accepted for any other reason, this code is returned
	RbacStatusCode_ERROR_CONFIG_NOT_ACCEPTED RbacStatusCode = 4
)

var RbacStatusCode_name = map[int32]string{
	0: "PENDING_VERIFICATION",
	1: "OK",
	2: "ERROR_RBAC_MODE_NOT_SUPPORTED_BY_MESH",
	3: "ERROR_CANNOT_ISOLATE_RBAC_SINCE_POLICIES_EXIST",
	4: "ERROR_CONFIG_NOT_ACCEPTED",
}

var RbacStatusCode_value = map[string]int32{
	"PENDING_VERIFICATION":                  0,
	"OK":                                    1,
	"ERROR_RBAC_MODE_NOT_SUPPORTED_BY_MESH": 2,
	"ERROR_CANNOT_ISOLATE_RBAC_SINCE_POLICIES_EXIST": 3,
	"ERROR_CONFIG_NOT_ACCEPTED":                      4,
}

func (x RbacStatusCode) String() string {
	return proto.EnumName(RbacStatusCode_name, int32(x))
}

func (RbacStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22c76505334eb8ab, []int{0}
}

// Configure RBAC properties on the mesh
type RbacMode struct {
	// Mode describes the desired RBAC behavior an optionally takes mode-specific configuration
	// implementation note: using oneof instead of enum since future modes may accept config
	//
	// Types that are valid to be assigned to Mode:
	//	*RbacMode_Disable_
	//	*RbacMode_Enable_
	Mode isRbacMode_Mode `protobuf_oneof:"mode"`
	// Set by operator
	// - Initialized as pending.
	// - If isolation cannot be expressed, an error code corresponding to the reason is reported.
	// - If isolation can be expressed, an "OK" status code is reported.
	Status               *RbacStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RbacMode) Reset()         { *m = RbacMode{} }
func (m *RbacMode) String() string { return proto.CompactTextString(m) }
func (*RbacMode) ProtoMessage()    {}
func (*RbacMode) Descriptor() ([]byte, []int) {
	return fileDescriptor_22c76505334eb8ab, []int{0}
}
func (m *RbacMode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RbacMode.Unmarshal(m, b)
}
func (m *RbacMode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RbacMode.Marshal(b, m, deterministic)
}
func (m *RbacMode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RbacMode.Merge(m, src)
}
func (m *RbacMode) XXX_Size() int {
	return xxx_messageInfo_RbacMode.Size(m)
}
func (m *RbacMode) XXX_DiscardUnknown() {
	xxx_messageInfo_RbacMode.DiscardUnknown(m)
}

var xxx_messageInfo_RbacMode proto.InternalMessageInfo

type isRbacMode_Mode interface {
	isRbacMode_Mode()
	Equal(interface{}) bool
}

type RbacMode_Disable_ struct {
	Disable *RbacMode_Disable `protobuf:"bytes,1,opt,name=disable,proto3,oneof" json:"disable,omitempty"`
}
type RbacMode_Enable_ struct {
	Enable *RbacMode_Enable `protobuf:"bytes,2,opt,name=enable,proto3,oneof" json:"enable,omitempty"`
}

func (*RbacMode_Disable_) isRbacMode_Mode() {}
func (*RbacMode_Enable_) isRbacMode_Mode()  {}

func (m *RbacMode) GetMode() isRbacMode_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *RbacMode) GetDisable() *RbacMode_Disable {
	if x, ok := m.GetMode().(*RbacMode_Disable_); ok {
		return x.Disable
	}
	return nil
}

func (m *RbacMode) GetEnable() *RbacMode_Enable {
	if x, ok := m.GetMode().(*RbacMode_Enable_); ok {
		return x.Enable
	}
	return nil
}

func (m *RbacMode) GetStatus() *RbacStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RbacMode) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RbacMode_Disable_)(nil),
		(*RbacMode_Enable_)(nil),
	}
}

// Disable indicates that RBAC policies should not be enforced
// Compatibility: [only: istio]
type RbacMode_Disable struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RbacMode_Disable) Reset()         { *m = RbacMode_Disable{} }
func (m *RbacMode_Disable) String() string { return proto.CompactTextString(m) }
func (*RbacMode_Disable) ProtoMessage()    {}
func (*RbacMode_Disable) Descriptor() ([]byte, []int) {
	return fileDescriptor_22c76505334eb8ab, []int{0, 0}
}
func (m *RbacMode_Disable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RbacMode_Disable.Unmarshal(m, b)
}
func (m *RbacMode_Disable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RbacMode_Disable.Marshal(b, m, deterministic)
}
func (m *RbacMode_Disable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RbacMode_Disable.Merge(m, src)
}
func (m *RbacMode_Disable) XXX_Size() int {
	return xxx_messageInfo_RbacMode_Disable.Size(m)
}
func (m *RbacMode_Disable) XXX_DiscardUnknown() {
	xxx_messageInfo_RbacMode_Disable.DiscardUnknown(m)
}

var xxx_messageInfo_RbacMode_Disable proto.InternalMessageInfo

// Enable mode tells the mesh to evaluate any policies that are defined
// Compatibility: [only: istio]
type RbacMode_Enable struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RbacMode_Enable) Reset()         { *m = RbacMode_Enable{} }
func (m *RbacMode_Enable) String() string { return proto.CompactTextString(m) }
func (*RbacMode_Enable) ProtoMessage()    {}
func (*RbacMode_Enable) Descriptor() ([]byte, []int) {
	return fileDescriptor_22c76505334eb8ab, []int{0, 1}
}
func (m *RbacMode_Enable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RbacMode_Enable.Unmarshal(m, b)
}
func (m *RbacMode_Enable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RbacMode_Enable.Marshal(b, m, deterministic)
}
func (m *RbacMode_Enable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RbacMode_Enable.Merge(m, src)
}
func (m *RbacMode_Enable) XXX_Size() int {
	return xxx_messageInfo_RbacMode_Enable.Size(m)
}
func (m *RbacMode_Enable) XXX_DiscardUnknown() {
	xxx_messageInfo_RbacMode_Enable.DiscardUnknown(m)
}

var xxx_messageInfo_RbacMode_Enable proto.InternalMessageInfo

type RbacStatus struct {
	// Status code summarizing Rbac Config acceptance state
	Code RbacStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=zephyr.solo.io.RbacStatusCode" json:"code,omitempty"`
	// As needed according to the status code, this message will surface any relevant configuration details or issues.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RbacStatus) Reset()         { *m = RbacStatus{} }
func (m *RbacStatus) String() string { return proto.CompactTextString(m) }
func (*RbacStatus) ProtoMessage()    {}
func (*RbacStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_22c76505334eb8ab, []int{1}
}
func (m *RbacStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RbacStatus.Unmarshal(m, b)
}
func (m *RbacStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RbacStatus.Marshal(b, m, deterministic)
}
func (m *RbacStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RbacStatus.Merge(m, src)
}
func (m *RbacStatus) XXX_Size() int {
	return xxx_messageInfo_RbacStatus.Size(m)
}
func (m *RbacStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RbacStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RbacStatus proto.InternalMessageInfo

func (m *RbacStatus) GetCode() RbacStatusCode {
	if m != nil {
		return m.Code
	}
	return RbacStatusCode_PENDING_VERIFICATION
}

func (m *RbacStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("zephyr.solo.io.RbacStatusCode", RbacStatusCode_name, RbacStatusCode_value)
	proto.RegisterType((*RbacMode)(nil), "zephyr.solo.io.RbacMode")
	proto.RegisterType((*RbacMode_Disable)(nil), "zephyr.solo.io.RbacMode.Disable")
	proto.RegisterType((*RbacMode_Enable)(nil), "zephyr.solo.io.RbacMode.Enable")
	proto.RegisterType((*RbacStatus)(nil), "zephyr.solo.io.RbacStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/v1/rbac.proto", fileDescriptor_22c76505334eb8ab)
}

var fileDescriptor_22c76505334eb8ab = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0x9b, 0x35, 0x64, 0xdb, 0x23, 0x2c, 0x61, 0xe8, 0xc5, 0x1a, 0x70, 0x2d, 0x05, 0xc1,
	0x0f, 0x9a, 0xb0, 0xf1, 0x46, 0xc1, 0x9b, 0x7c, 0x4c, 0xdb, 0xc1, 0x6e, 0x12, 0x26, 0x51, 0xb4,
	0x37, 0x21, 0x1f, 0x63, 0x36, 0x76, 0x77, 0x27, 0x24, 0x59, 0x45, 0xaf, 0x7c, 0x1c, 0x2f, 0x7d,
	0x1e, 0x2f, 0x7d, 0x03, 0xdf, 0x40, 0x76, 0x92, 0x20, 0x85, 0x2e, 0x5e, 0x65, 0x4e, 0xe6, 0xf7,
	0xfb, 0xcf, 0x39, 0x70, 0xc0, 0x2c, 0xca, 0x76, 0xb9, 0x4d, 0xf5, 0x8c, 0xaf, 0x8d, 0x86, 0xaf,
	0xf8, 0x59, 0xc9, 0x8d, 0x35, 0x6b, 0x96, 0x67, 0x55, 0xcd, 0x3f, 0xb1, 0xac, 0x6d, 0x8c, 0xa4,
	0x2a, 0x8d, 0xcf, 0x73, 0xa3, 0x4e, 0x93, 0x4c, 0xaf, 0x6a, 0xde, 0x72, 0x34, 0xf9, 0xc6, 0xaa,
	0xe5, 0xd7, 0x5a, 0xdf, 0xf1, 0x7a, 0xc9, 0xb5, 0xf9, 0x1d, 0x19, 0xe2, 0x7b, 0x53, 0xb6, 0x83,
	0x3e, 0xd4, 0x5d, 0x84, 0x76, 0x5c, 0xf0, 0x82, 0x8b, 0xa3, 0xb1, 0x3b, 0xf5, 0x7f, 0x67, 0x05,
	0xe7, 0xc5, 0x8a, 0x19, 0xa2, 0x4a, 0xb7, 0x1f, 0x8d, 0x7c, 0x5b, 0x27, 0x6d, 0xc9, 0x37, 0xfb,
	0xee, 0xbf, 0xd4, 0x49, 0x55, 0xb1, 0xba, 0xe9, 0xee, 0x4f, 0x7f, 0x49, 0x70, 0x48, 0xd3, 0x24,
	0x5b, 0xf0, 0x9c, 0xa1, 0xd7, 0x30, 0xce, 0xcb, 0x26, 0x49, 0x57, 0x6c, 0x2a, 0x9d, 0x48, 0x4f,
	0xee, 0x9b, 0x27, 0xfa, 0xed, 0xbe, 0xf5, 0x01, 0xd5, 0xdd, 0x8e, 0xbb, 0x3c, 0xa0, 0x83, 0x82,
	0x5e, 0x81, 0xc2, 0x36, 0x42, 0x1e, 0x09, 0xf9, 0xd1, 0x5e, 0x19, 0x6f, 0x7a, 0xb7, 0x17, 0xd0,
	0x4b, 0x50, 0x9a, 0x36, 0x69, 0xb7, 0xcd, 0x54, 0x16, 0xaa, 0x76, 0x97, 0x1a, 0x0a, 0xc2, 0x96,
	0xbf, 0xff, 0x91, 0x25, 0xda, 0xf3, 0xda, 0x11, 0x8c, 0xfb, 0x56, 0xb4, 0x43, 0x50, 0xba, 0x60,
	0x5b, 0x01, 0x79, 0xcd, 0x73, 0x76, 0x7a, 0x0d, 0xf0, 0x4f, 0x44, 0x26, 0xc8, 0x19, 0xcf, 0xbb,
	0xd1, 0x26, 0xe6, 0x6c, 0xff, 0x13, 0x0e, 0xcf, 0x19, 0x15, 0x2c, 0x9a, 0xc2, 0x78, 0xcd, 0x9a,
	0x26, 0x29, 0xba, 0xa1, 0x8e, 0xe8, 0x50, 0x3e, 0xfb, 0x29, 0xc1, 0xe4, 0xb6, 0x82, 0xa6, 0x70,
	0x1c, 0x60, 0xcf, 0x25, 0xde, 0x45, 0xfc, 0x0e, 0x53, 0x72, 0x4e, 0x1c, 0x2b, 0x22, 0xbe, 0xa7,
	0x1e, 0x20, 0x05, 0x46, 0xfe, 0x1b, 0x55, 0x42, 0x4f, 0xe1, 0x31, 0xa6, 0xd4, 0xa7, 0x31, 0xb5,
	0x2d, 0x27, 0x5e, 0xf8, 0x2e, 0x8e, 0x3d, 0x3f, 0x8a, 0xc3, 0xb7, 0x41, 0xe0, 0xd3, 0x08, 0xbb,
	0xb1, 0xfd, 0x21, 0x5e, 0xe0, 0xf0, 0x52, 0x1d, 0x21, 0x13, 0xf4, 0x0e, 0x75, 0x2c, 0x6f, 0xc7,
	0x90, 0xd0, 0xbf, 0xb2, 0x22, 0xdc, 0x99, 0x21, 0xf1, 0x1c, 0x1c, 0x07, 0xfe, 0x15, 0x71, 0x08,
	0x0e, 0x63, 0xfc, 0x9e, 0x84, 0x91, 0x7a, 0x0f, 0x3d, 0x84, 0x07, 0xbd, 0xe3, 0x7b, 0xe7, 0xe4,
	0x42, 0x64, 0x5b, 0x8e, 0x83, 0x83, 0x08, 0xbb, 0xaa, 0x6c, 0xcf, 0x7f, 0xfc, 0x9e, 0x49, 0xd7,
	0xcf, 0xff, 0xbb, 0xbe, 0xd5, 0x4d, 0xd1, 0xef, 0x60, 0xaa, 0x88, 0x2d, 0x79, 0xf1, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xf3, 0xd8, 0x04, 0xd2, 0xf4, 0x02, 0x00, 0x00,
}

func (this *RbacMode) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RbacMode)
	if !ok {
		that2, ok := that.(RbacMode)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Mode == nil {
		if this.Mode != nil {
			return false
		}
	} else if this.Mode == nil {
		return false
	} else if !this.Mode.Equal(that1.Mode) {
		return false
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RbacMode_Disable_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RbacMode_Disable_)
	if !ok {
		that2, ok := that.(RbacMode_Disable_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Disable.Equal(that1.Disable) {
		return false
	}
	return true
}
func (this *RbacMode_Enable_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RbacMode_Enable_)
	if !ok {
		that2, ok := that.(RbacMode_Enable_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Enable.Equal(that1.Enable) {
		return false
	}
	return true
}
func (this *RbacMode_Disable) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RbacMode_Disable)
	if !ok {
		that2, ok := that.(RbacMode_Disable)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RbacMode_Enable) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RbacMode_Enable)
	if !ok {
		that2, ok := that.(RbacMode_Enable)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RbacStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RbacStatus)
	if !ok {
		that2, ok := that.(RbacStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}