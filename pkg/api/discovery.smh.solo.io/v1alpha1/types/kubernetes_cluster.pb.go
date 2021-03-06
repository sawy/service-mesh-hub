// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/kubernetes_cluster.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
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

//
//Representation of a Kubernetes cluster that has been registered in Service Mesh Hub.
type KubernetesClusterSpec struct {
	// pointer to secret which contains the kubeconfig with information to connect to the remote cluster.
	SecretRef *types.ResourceRef `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// context to use within the kubeconfig pointed to by the above reference
	Context string `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	// version of kubernetes
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// cloud provider, empty if unknown
	Cloud string `protobuf:"bytes,4,opt,name=cloud,proto3" json:"cloud,omitempty"`
	// namespace to use when writing Service Mesh Hub resources to this cluster
	WriteNamespace       string   `protobuf:"bytes,5,opt,name=write_namespace,json=writeNamespace,proto3" json:"write_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesClusterSpec) Reset()         { *m = KubernetesClusterSpec{} }
func (m *KubernetesClusterSpec) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec) ProtoMessage()    {}
func (*KubernetesClusterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_cebb6fc6a7744031, []int{0}
}
func (m *KubernetesClusterSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec.Merge(m, src)
}
func (m *KubernetesClusterSpec) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec.Size(m)
}
func (m *KubernetesClusterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec proto.InternalMessageInfo

func (m *KubernetesClusterSpec) GetSecretRef() *types.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *KubernetesClusterSpec) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *KubernetesClusterSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *KubernetesClusterSpec) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *KubernetesClusterSpec) GetWriteNamespace() string {
	if m != nil {
		return m.WriteNamespace
	}
	return ""
}

func init() {
	proto.RegisterType((*KubernetesClusterSpec)(nil), "discovery.smh.solo.io.KubernetesClusterSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/kubernetes_cluster.proto", fileDescriptor_cebb6fc6a7744031)
}

var fileDescriptor_cebb6fc6a7744031 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x18, 0xc4, 0xc9, 0xfb, 0x5a, 0xa5, 0x2b, 0x28, 0x04, 0x0b, 0x41, 0x10, 0x8a, 0x17, 0x7b, 0xb0,
	0x1b, 0xaa, 0x57, 0x4f, 0x7a, 0x14, 0x44, 0xd2, 0x9b, 0x97, 0x92, 0x6c, 0x27, 0xcd, 0xd2, 0x24,
	0xcf, 0xb2, 0xcf, 0x6e, 0xb4, 0xdf, 0xd0, 0x8f, 0x25, 0x49, 0xec, 0x1f, 0xc1, 0x83, 0xc7, 0x99,
	0xd9, 0xfd, 0x0d, 0xcf, 0x88, 0xf9, 0x4a, 0xbb, 0xc2, 0x67, 0x52, 0x51, 0x15, 0x33, 0x95, 0x34,
	0xd5, 0x14, 0x33, 0x6c, 0xa3, 0x15, 0xa6, 0x15, 0xb8, 0x98, 0x16, 0x3e, 0x8b, 0x53, 0xa3, 0xe3,
	0xa5, 0x66, 0x45, 0x0d, 0xec, 0x26, 0x6e, 0x66, 0x69, 0x69, 0x8a, 0x74, 0x16, 0xaf, 0x7d, 0x06,
	0x5b, 0xc3, 0x81, 0x17, 0xaa, 0xf4, 0xec, 0x60, 0xa5, 0xb1, 0xe4, 0x28, 0x1c, 0xed, 0x1e, 0x4b,
	0xae, 0x0a, 0xd9, 0x72, 0xa5, 0xa6, 0xcb, 0xdb, 0x5f, 0xc1, 0x8a, 0x2c, 0xf6, 0x4c, 0x8b, 0xbc,
	0x87, 0x5c, 0x7f, 0x06, 0x62, 0xf4, 0xbc, 0x6b, 0x78, 0xea, 0x0b, 0xe6, 0x06, 0x2a, 0x7c, 0x10,
	0x82, 0xa1, 0x2c, 0xdc, 0xc2, 0x22, 0x8f, 0x82, 0x71, 0x30, 0x39, 0xbd, 0xbb, 0x92, 0x2d, 0xe7,
	0xb0, 0x4e, 0x26, 0x60, 0xf2, 0x56, 0x21, 0x41, 0x9e, 0x0c, 0xfb, 0x0f, 0x09, 0xf2, 0x30, 0x12,
	0x27, 0x8a, 0x6a, 0x87, 0x0f, 0x17, 0xfd, 0x1b, 0x07, 0x93, 0x61, 0xb2, 0x95, 0x6d, 0xd2, 0xc0,
	0xb2, 0xa6, 0x3a, 0xfa, 0xdf, 0x27, 0xdf, 0x32, 0xbc, 0x10, 0x03, 0x55, 0x92, 0x5f, 0x46, 0x47,
	0x9d, 0xdf, 0x8b, 0xf0, 0x46, 0x9c, 0xbf, 0x5b, 0xed, 0xb0, 0xa8, 0xd3, 0x0a, 0x6c, 0x52, 0x85,
	0x68, 0xd0, 0xe5, 0x67, 0x9d, 0xfd, 0xb2, 0x75, 0x1f, 0x93, 0xb7, 0xd7, 0xbf, 0xcc, 0x6c, 0xd6,
	0xab, 0x9f, 0x53, 0x1f, 0x9e, 0xb3, 0x9f, 0xc8, 0x6d, 0x0c, 0x38, 0x3b, 0xee, 0x56, 0xba, 0xff,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x5d, 0xe1, 0x01, 0xc1, 0x01, 0x00, 0x00,
}
