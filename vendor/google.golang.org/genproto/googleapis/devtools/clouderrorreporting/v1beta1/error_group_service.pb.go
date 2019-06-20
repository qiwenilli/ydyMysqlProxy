// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/clouderrorreporting/v1beta1/error_group_service.proto

package clouderrorreporting

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// A request to return an individual group.
type GetGroupRequest struct {
	// [Required] The group resource name. Written as
	// <code>projects/<var>projectID</var>/groups/<var>group_name</var></code>.
	// Call
	// <a href="/error-reporting/reference/rest/v1beta1/projects.groupStats/list">
	// <code>groupStats.list</code></a> to return a list of groups belonging to
	// this project.
	//
	// Example: <code>projects/my-project-123/groups/my-group</code>
	GroupName            string   `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupRequest) Reset()         { *m = GetGroupRequest{} }
func (m *GetGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupRequest) ProtoMessage()    {}
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cada5d349d61e784, []int{0}
}

func (m *GetGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupRequest.Unmarshal(m, b)
}
func (m *GetGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupRequest.Merge(m, src)
}
func (m *GetGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupRequest.Size(m)
}
func (m *GetGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupRequest proto.InternalMessageInfo

func (m *GetGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

// A request to replace the existing data for the given group.
type UpdateGroupRequest struct {
	// [Required] The group which replaces the resource on the server.
	Group                *ErrorGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateGroupRequest) Reset()         { *m = UpdateGroupRequest{} }
func (m *UpdateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupRequest) ProtoMessage()    {}
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cada5d349d61e784, []int{1}
}

func (m *UpdateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupRequest.Unmarshal(m, b)
}
func (m *UpdateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupRequest.Marshal(b, m, deterministic)
}
func (m *UpdateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupRequest.Merge(m, src)
}
func (m *UpdateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupRequest.Size(m)
}
func (m *UpdateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupRequest proto.InternalMessageInfo

func (m *UpdateGroupRequest) GetGroup() *ErrorGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGroupRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.GetGroupRequest")
	proto.RegisterType((*UpdateGroupRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.UpdateGroupRequest")
}

func init() {
	proto.RegisterFile("google/devtools/clouderrorreporting/v1beta1/error_group_service.proto", fileDescriptor_cada5d349d61e784)
}

var fileDescriptor_cada5d349d61e784 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0x4a, 0x23, 0x41,
	0x14, 0x86, 0xe9, 0x0c, 0x33, 0x4c, 0x2a, 0x8b, 0x61, 0x6a, 0x31, 0x0c, 0xcd, 0x0c, 0x48, 0xdc,
	0x68, 0x02, 0x55, 0x76, 0x5c, 0x18, 0xbc, 0x20, 0x44, 0x42, 0x56, 0x4a, 0x88, 0x98, 0x85, 0x04,
	0x43, 0xa5, 0x53, 0x14, 0x2d, 0xdd, 0x75, 0xda, 0xea, 0x4a, 0x36, 0xe2, 0xc6, 0x07, 0x70, 0xe3,
	0x5b, 0xb8, 0xf6, 0x05, 0xdc, 0xba, 0xf5, 0x15, 0x7c, 0x07, 0xb7, 0xd2, 0x55, 0xb9, 0x98, 0x8b,
	0x60, 0x67, 0x7b, 0x2e, 0xff, 0xff, 0xd5, 0x5f, 0x07, 0xd5, 0x05, 0x80, 0x08, 0x39, 0xed, 0xf3,
	0xa1, 0x06, 0x08, 0x13, 0xea, 0x87, 0x30, 0xe8, 0x73, 0xa5, 0x40, 0x29, 0x1e, 0x83, 0xd2, 0x81,
	0x14, 0x74, 0xe8, 0xf5, 0xb8, 0x66, 0x1e, 0x35, 0xe5, 0xae, 0x50, 0x30, 0x88, 0xbb, 0x09, 0x57,
	0xc3, 0xc0, 0xe7, 0x24, 0x56, 0xa0, 0x01, 0x97, 0xad, 0x0c, 0x19, 0xcb, 0x90, 0x25, 0x32, 0x64,
	0x24, 0xe3, 0xfe, 0x1b, 0x79, 0xb2, 0x38, 0xa0, 0x4c, 0x4a, 0xd0, 0x4c, 0x07, 0x20, 0x13, 0x2b,
	0xe5, 0x56, 0xb3, 0x10, 0xf9, 0x10, 0x45, 0x20, 0xed, 0x66, 0x71, 0x0b, 0xfd, 0x6a, 0x70, 0xdd,
	0x48, 0xf1, 0x5a, 0xfc, 0x6a, 0xc0, 0x13, 0x8d, 0xff, 0x23, 0x64, 0x71, 0x25, 0x8b, 0xf8, 0x5f,
	0x67, 0xcd, 0xd9, 0xc8, 0xb7, 0xf2, 0xa6, 0x72, 0xc2, 0x22, 0x5e, 0xf4, 0x11, 0x3e, 0x8b, 0xfb,
	0x4c, 0xf3, 0x99, 0xa5, 0x63, 0xf4, 0xdd, 0x8c, 0x98, 0xf9, 0x42, 0x65, 0x87, 0x64, 0x78, 0x1c,
	0xa9, 0xa7, 0x65, 0x2b, 0x67, 0x55, 0x2a, 0x77, 0xdf, 0xd0, 0xef, 0x69, 0xf5, 0xd4, 0xe6, 0x86,
	0x1f, 0x1d, 0xf4, 0x73, 0x4c, 0x8b, 0xf7, 0x33, 0x59, 0xcc, 0x3d, 0xd2, 0x5d, 0x15, 0xb0, 0xe8,
	0xdd, 0xbe, 0xbc, 0xde, 0xe7, 0xca, 0x78, 0x73, 0x92, 0xe7, 0xf5, 0x34, 0xad, 0x83, 0x58, 0xc1,
	0x25, 0xf7, 0x75, 0x42, 0x4b, 0xd4, 0x54, 0x13, 0x5a, 0xba, 0xc1, 0x4f, 0x0e, 0x2a, 0x7c, 0x88,
	0x0c, 0x1f, 0x66, 0xf2, 0x5e, 0x0c, 0x7b, 0x75, 0xf8, 0xaa, 0x81, 0xaf, 0xb8, 0xf3, 0xf0, 0xe4,
	0x53, 0xf8, 0x5d, 0xfb, 0x21, 0xb5, 0x37, 0x07, 0xa5, 0x87, 0x93, 0xc5, 0xb8, 0xf6, 0x67, 0xe1,
	0x07, 0x9b, 0xe9, 0xcd, 0x35, 0x9d, 0xf3, 0x8b, 0x91, 0x8c, 0x80, 0x90, 0x49, 0x41, 0x40, 0x09,
	0x2a, 0xb8, 0x34, 0x17, 0x49, 0x6d, 0x8b, 0xc5, 0x41, 0xf2, 0xa5, 0x73, 0xde, 0x5b, 0xd2, 0x7b,
	0xc8, 0xad, 0x37, 0xac, 0xc1, 0x51, 0xda, 0xb4, 0x09, 0xb4, 0x26, 0x7c, 0x6d, 0xaf, 0x96, 0x6e,
	0x3e, 0x8f, 0xa7, 0x3a, 0x66, 0xaa, 0x33, 0x3b, 0xd5, 0x69, 0x5b, 0xfd, 0xde, 0x0f, 0x83, 0xb5,
	0xfd, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x0a, 0xfa, 0x93, 0xf6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ErrorGroupServiceClient is the client API for ErrorGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ErrorGroupServiceClient interface {
	// Get the specified group.
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error)
	// Replace the data for the specified group.
	// Fails if the group does not exist.
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error)
}

type errorGroupServiceClient struct {
	cc *grpc.ClientConn
}

func NewErrorGroupServiceClient(cc *grpc.ClientConn) ErrorGroupServiceClient {
	return &errorGroupServiceClient{cc}
}

func (c *errorGroupServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error) {
	out := new(ErrorGroup)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorGroupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error) {
	out := new(ErrorGroup)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErrorGroupServiceServer is the server API for ErrorGroupService service.
type ErrorGroupServiceServer interface {
	// Get the specified group.
	GetGroup(context.Context, *GetGroupRequest) (*ErrorGroup, error)
	// Replace the data for the specified group.
	// Fails if the group does not exist.
	UpdateGroup(context.Context, *UpdateGroupRequest) (*ErrorGroup, error)
}

func RegisterErrorGroupServiceServer(s *grpc.Server, srv ErrorGroupServiceServer) {
	s.RegisterService(&_ErrorGroupService_serviceDesc, srv)
}

func _ErrorGroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorGroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorGroupServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ErrorGroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorGroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorGroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ErrorGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouderrorreporting.v1beta1.ErrorGroupService",
	HandlerType: (*ErrorGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroup",
			Handler:    _ErrorGroupService_GetGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _ErrorGroupService_UpdateGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouderrorreporting/v1beta1/error_group_service.proto",
}
