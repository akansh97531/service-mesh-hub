// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service-mesh-hub/services/apiserver/api/v1/virtual_mesh.proto

package v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	types "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VirtualMesh struct {
	Spec                 *types.VirtualMeshSpec   `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Status               *types.VirtualMeshStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Ref                  *types1.ResourceRef      `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref,omitempty"`
	Labels               map[string]string        `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *VirtualMesh) Reset()         { *m = VirtualMesh{} }
func (m *VirtualMesh) String() string { return proto.CompactTextString(m) }
func (*VirtualMesh) ProtoMessage()    {}
func (*VirtualMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4a9dd7fa6557580, []int{0}
}
func (m *VirtualMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMesh.Unmarshal(m, b)
}
func (m *VirtualMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMesh.Marshal(b, m, deterministic)
}
func (m *VirtualMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMesh.Merge(m, src)
}
func (m *VirtualMesh) XXX_Size() int {
	return xxx_messageInfo_VirtualMesh.Size(m)
}
func (m *VirtualMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMesh.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMesh proto.InternalMessageInfo

func (m *VirtualMesh) GetSpec() *types.VirtualMeshSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *VirtualMesh) GetStatus() *types.VirtualMeshStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualMesh) GetRef() *types1.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *VirtualMesh) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ListVirtualMeshesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVirtualMeshesRequest) Reset()         { *m = ListVirtualMeshesRequest{} }
func (m *ListVirtualMeshesRequest) String() string { return proto.CompactTextString(m) }
func (*ListVirtualMeshesRequest) ProtoMessage()    {}
func (*ListVirtualMeshesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4a9dd7fa6557580, []int{1}
}
func (m *ListVirtualMeshesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVirtualMeshesRequest.Unmarshal(m, b)
}
func (m *ListVirtualMeshesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVirtualMeshesRequest.Marshal(b, m, deterministic)
}
func (m *ListVirtualMeshesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVirtualMeshesRequest.Merge(m, src)
}
func (m *ListVirtualMeshesRequest) XXX_Size() int {
	return xxx_messageInfo_ListVirtualMeshesRequest.Size(m)
}
func (m *ListVirtualMeshesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVirtualMeshesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListVirtualMeshesRequest proto.InternalMessageInfo

type ListVirtualMeshesResponse struct {
	VirtualMeshes        []*VirtualMesh `protobuf:"bytes,1,rep,name=virtual_meshes,json=virtualMeshes,proto3" json:"virtual_meshes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListVirtualMeshesResponse) Reset()         { *m = ListVirtualMeshesResponse{} }
func (m *ListVirtualMeshesResponse) String() string { return proto.CompactTextString(m) }
func (*ListVirtualMeshesResponse) ProtoMessage()    {}
func (*ListVirtualMeshesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4a9dd7fa6557580, []int{2}
}
func (m *ListVirtualMeshesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVirtualMeshesResponse.Unmarshal(m, b)
}
func (m *ListVirtualMeshesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVirtualMeshesResponse.Marshal(b, m, deterministic)
}
func (m *ListVirtualMeshesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVirtualMeshesResponse.Merge(m, src)
}
func (m *ListVirtualMeshesResponse) XXX_Size() int {
	return xxx_messageInfo_ListVirtualMeshesResponse.Size(m)
}
func (m *ListVirtualMeshesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVirtualMeshesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListVirtualMeshesResponse proto.InternalMessageInfo

func (m *ListVirtualMeshesResponse) GetVirtualMeshes() []*VirtualMesh {
	if m != nil {
		return m.VirtualMeshes
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualMesh)(nil), "rpc.solo.io.VirtualMesh")
	proto.RegisterMapType((map[string]string)(nil), "rpc.solo.io.VirtualMesh.LabelsEntry")
	proto.RegisterType((*ListVirtualMeshesRequest)(nil), "rpc.solo.io.ListVirtualMeshesRequest")
	proto.RegisterType((*ListVirtualMeshesResponse)(nil), "rpc.solo.io.ListVirtualMeshesResponse")
}

func init() {
	proto.RegisterFile("service-mesh-hub/services/apiserver/api/v1/virtual_mesh.proto", fileDescriptor_e4a9dd7fa6557580)
}

var fileDescriptor_e4a9dd7fa6557580 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x49, 0x33, 0x54, 0x1a, 0x47, 0x8c, 0xc0, 0x62, 0x11, 0xb2, 0xaa, 0x22, 0x40, 0x23,
	0x34, 0x75, 0x34, 0x61, 0xc3, 0x3f, 0x1a, 0x7e, 0x76, 0xc3, 0xc6, 0x48, 0x2c, 0x10, 0x12, 0x4a,
	0xc2, 0x6d, 0x63, 0x35, 0x8d, 0x8d, 0xaf, 0x13, 0x54, 0x9e, 0x8e, 0x47, 0x43, 0x76, 0x82, 0x6a,
	0x68, 0x2b, 0xba, 0xbb, 0xb1, 0xcf, 0x77, 0xce, 0x89, 0x13, 0x93, 0x97, 0x08, 0xba, 0x17, 0x15,
	0xcc, 0xd7, 0x80, 0xf5, 0xbc, 0xee, 0xca, 0x6c, 0x5c, 0xc0, 0xac, 0x50, 0xc2, 0xce, 0xa0, 0xed,
	0x94, 0xf5, 0x97, 0x59, 0x2f, 0xb4, 0xe9, 0x8a, 0xe6, 0xab, 0x95, 0x32, 0xa5, 0xa5, 0x91, 0x34,
	0xd2, 0xaa, 0x62, 0x28, 0x1b, 0xc9, 0x84, 0x4c, 0x9e, 0xef, 0x78, 0x59, 0xb0, 0x05, 0xf3, 0x43,
	0xea, 0x95, 0x68, 0x97, 0x59, 0x7f, 0x59, 0x34, 0xaa, 0x2e, 0xf6, 0x39, 0x25, 0x17, 0x7b, 0xe1,
	0x4a, 0x6a, 0xd8, 0x62, 0x1a, 0x16, 0x83, 0x3a, 0xfd, 0x35, 0x21, 0xd1, 0xa7, 0xc1, 0xe4, 0x03,
	0x60, 0x4d, 0x5f, 0x91, 0x13, 0x54, 0x50, 0xc5, 0xc1, 0x2c, 0x38, 0x8f, 0xf2, 0x47, 0x6c, 0x1b,
	0xca, 0x7e, 0x82, 0xaa, 0x37, 0xfa, 0x4f, 0x49, 0xe6, 0x51, 0x1f, 0x15, 0x54, 0xdc, 0x71, 0xf4,
	0x1d, 0x99, 0xa2, 0x29, 0x4c, 0x87, 0xf1, 0xc4, 0x39, 0x5c, 0x1c, 0xe9, 0xe0, 0x18, 0x3e, 0xb2,
	0x34, 0x27, 0xa1, 0x86, 0x45, 0x1c, 0x3a, 0x8b, 0x19, 0xb3, 0xe5, 0xff, 0x85, 0x39, 0xa0, 0xec,
	0x74, 0x05, 0x1c, 0x16, 0xdc, 0x8a, 0xe9, 0x0b, 0x32, 0x6d, 0x8a, 0x12, 0x1a, 0x8c, 0x4f, 0x66,
	0xe1, 0x79, 0x94, 0xdf, 0x67, 0xde, 0x91, 0xfa, 0x59, 0xec, 0xda, 0xc9, 0xde, 0xb7, 0x46, 0x6f,
	0xf8, 0xc8, 0x24, 0x4f, 0x49, 0xe4, 0x2d, 0xd3, 0xdb, 0x24, 0x5c, 0xc1, 0xc6, 0x9d, 0xc2, 0x29,
	0xb7, 0x23, 0xbd, 0x4b, 0x6e, 0xf6, 0x45, 0xd3, 0x81, 0x7b, 0xaf, 0x53, 0x3e, 0x3c, 0x3c, 0x9b,
	0x3c, 0x09, 0xd2, 0x84, 0xc4, 0xd7, 0x02, 0x8d, 0x97, 0x00, 0xc8, 0xe1, 0x7b, 0x07, 0x68, 0xd2,
	0x2f, 0xe4, 0xde, 0x9e, 0x3d, 0x54, 0xb2, 0x45, 0xa0, 0xaf, 0xc9, 0x99, 0xff, 0xfd, 0x00, 0xe3,
	0xc0, 0x35, 0x8f, 0x0f, 0x35, 0xe7, 0xb7, 0x7a, 0xdf, 0x28, 0xef, 0xc9, 0x99, 0xb7, 0x7b, 0xa5,
	0x04, 0xfd, 0x46, 0xee, 0xec, 0xe4, 0xd1, 0x07, 0x7f, 0xf9, 0x1d, 0xea, 0x9a, 0x3c, 0xfc, 0x9f,
	0x6c, 0xa8, 0x9d, 0xde, 0x78, 0xf3, 0xf6, 0xf3, 0xd5, 0x52, 0x98, 0xba, 0x2b, 0x59, 0x25, 0xd7,
	0x99, 0x25, 0xe6, 0x42, 0x66, 0xc7, 0x5c, 0x00, 0xb5, 0x5a, 0x8e, 0x97, 0xa0, 0x9c, 0xba, 0x1f,
	0xf0, 0xf1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0x0e, 0x13, 0xdd, 0x39, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualMeshApiClient is the client API for VirtualMeshApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMeshApiClient interface {
	ListVirtualMeshes(ctx context.Context, in *ListVirtualMeshesRequest, opts ...grpc.CallOption) (*ListVirtualMeshesResponse, error)
}

type virtualMeshApiClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMeshApiClient(cc *grpc.ClientConn) VirtualMeshApiClient {
	return &virtualMeshApiClient{cc}
}

func (c *virtualMeshApiClient) ListVirtualMeshes(ctx context.Context, in *ListVirtualMeshesRequest, opts ...grpc.CallOption) (*ListVirtualMeshesResponse, error) {
	out := new(ListVirtualMeshesResponse)
	err := c.cc.Invoke(ctx, "/rpc.solo.io.VirtualMeshApi/ListVirtualMeshes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMeshApiServer is the server API for VirtualMeshApi service.
type VirtualMeshApiServer interface {
	ListVirtualMeshes(context.Context, *ListVirtualMeshesRequest) (*ListVirtualMeshesResponse, error)
}

// UnimplementedVirtualMeshApiServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMeshApiServer struct {
}

func (*UnimplementedVirtualMeshApiServer) ListVirtualMeshes(ctx context.Context, req *ListVirtualMeshesRequest) (*ListVirtualMeshesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVirtualMeshes not implemented")
}

func RegisterVirtualMeshApiServer(s *grpc.Server, srv VirtualMeshApiServer) {
	s.RegisterService(&_VirtualMeshApi_serviceDesc, srv)
}

func _VirtualMeshApi_ListVirtualMeshes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVirtualMeshesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMeshApiServer).ListVirtualMeshes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.solo.io.VirtualMeshApi/ListVirtualMeshes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMeshApiServer).ListVirtualMeshes(ctx, req.(*ListVirtualMeshesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMeshApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.solo.io.VirtualMeshApi",
	HandlerType: (*VirtualMeshApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVirtualMeshes",
			Handler:    _VirtualMeshApi_ListVirtualMeshes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-mesh-hub/services/apiserver/api/v1/virtual_mesh.proto",
}
