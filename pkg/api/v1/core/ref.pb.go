// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/v1/core/ref.proto

package core

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

type ClusterResourceRef struct {
	Resource             *core.ResourceRef `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Cluster              string            `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterResourceRef) Reset()         { *m = ClusterResourceRef{} }
func (m *ClusterResourceRef) String() string { return proto.CompactTextString(m) }
func (*ClusterResourceRef) ProtoMessage()    {}
func (*ClusterResourceRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8132c94bbd6b2eb2, []int{0}
}
func (m *ClusterResourceRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResourceRef.Unmarshal(m, b)
}
func (m *ClusterResourceRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResourceRef.Marshal(b, m, deterministic)
}
func (m *ClusterResourceRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResourceRef.Merge(m, src)
}
func (m *ClusterResourceRef) XXX_Size() int {
	return xxx_messageInfo_ClusterResourceRef.Size(m)
}
func (m *ClusterResourceRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResourceRef.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResourceRef proto.InternalMessageInfo

func (m *ClusterResourceRef) GetResource() *core.ResourceRef {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ClusterResourceRef) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func init() {
	proto.RegisterType((*ClusterResourceRef)(nil), "core.zephyr.solo.io.ClusterResourceRef")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/v1/core/ref.proto", fileDescriptor_8132c94bbd6b2eb2)
}

var fileDescriptor_8132c94bbd6b2eb2 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0xcf,
	0x4d, 0x2d, 0xce, 0xd0, 0x2d, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x4f, 0x2c, 0xc8,
	0xd4, 0x2f, 0x33, 0xd4, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x4a, 0x4d, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x06, 0xf1, 0xf5, 0xaa, 0x52, 0x0b, 0x32, 0x2a, 0x8b, 0xf4, 0x40, 0x1a,
	0xf5, 0x32, 0xf3, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xf2, 0xfa, 0x20, 0x16, 0x44, 0xa9,
	0x94, 0x0e, 0x16, 0x2b, 0xc0, 0x74, 0x76, 0x66, 0x09, 0xcc, 0x74, 0xb8, 0xc1, 0x4a, 0xa9, 0x5c,
	0x42, 0xce, 0x39, 0xa5, 0xc5, 0x25, 0xa9, 0x45, 0x41, 0xa9, 0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9,
	0x41, 0xa9, 0x69, 0x42, 0xa6, 0x5c, 0x1c, 0x45, 0x50, 0xae, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7,
	0x91, 0xa4, 0x1e, 0xd8, 0x05, 0x50, 0xab, 0xf5, 0x90, 0x14, 0x07, 0xc1, 0x95, 0x0a, 0x49, 0x70,
	0xb1, 0x27, 0x43, 0x0c, 0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x9d, 0xcc, 0x56,
	0x3c, 0x92, 0x63, 0x8c, 0x32, 0x20, 0xe8, 0xfb, 0x82, 0xec, 0x74, 0xe4, 0x10, 0x48, 0x62, 0x03,
	0xbb, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xfc, 0x6c, 0x33, 0x38, 0x01, 0x00, 0x00,
}

func (this *ClusterResourceRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterResourceRef)
	if !ok {
		that2, ok := that.(ClusterResourceRef)
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
	if !this.Resource.Equal(that1.Resource) {
		return false
	}
	if this.Cluster != that1.Cluster {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
