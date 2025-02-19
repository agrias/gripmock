// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.2
// source: stub-subfolders.proto

package stub_subfolders

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message containing the user's name.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_simple_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ReturnCode int32  `protobuf:"varint,2,opt,name=return_code,json=returnCode,proto3" json:"return_code,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_simple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_simple_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Reply) GetReturnCode() int32 {
	if x != nil {
		return x.ReturnCode
	}
	return 0
}

var File_simple_proto protoreflect.FileDescriptor

var file_simple_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x36, 0x0a, 0x08, 0x47, 0x72, 0x69,
	0x70, 0x6d, 0x6f, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x12, 0x0f, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simple_proto_rawDescOnce sync.Once
	file_simple_proto_rawDescData = file_simple_proto_rawDesc
)

func file_simple_proto_rawDescGZIP() []byte {
	file_simple_proto_rawDescOnce.Do(func() {
		file_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_simple_proto_rawDescData)
	})
	return file_simple_proto_rawDescData
}

var file_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_simple_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: simple.Request
	(*Reply)(nil),   // 1: simple.Reply
}
var file_simple_proto_depIdxs = []int32{
	0, // 0: simple.Gripmock.SayHello:input_type -> simple.Request
	1, // 1: simple.Gripmock.SayHello:output_type -> simple.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_simple_proto_init() }
func file_simple_proto_init() {
	if File_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_simple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_simple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simple_proto_goTypes,
		DependencyIndexes: file_simple_proto_depIdxs,
		MessageInfos:      file_simple_proto_msgTypes,
	}.Build()
	File_simple_proto = out.File
	file_simple_proto_rawDesc = nil
	file_simple_proto_goTypes = nil
	file_simple_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GripmockClient is the client API for Gripmock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GripmockClient interface {
	// simple unary method
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type gripmockClient struct {
	cc grpc.ClientConnInterface
}

func NewGripmockClient(cc grpc.ClientConnInterface) GripmockClient {
	return &gripmockClient{cc}
}

func (c *gripmockClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/simple.Gripmock/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripmockServer is the server API for Gripmock service.
type GripmockServer interface {
	// simple unary method
	SayHello(context.Context, *Request) (*Reply, error)
}

// UnimplementedGripmockServer can be embedded to have forward compatible implementations.
type UnimplementedGripmockServer struct {
}

func (*UnimplementedGripmockServer) SayHello(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGripmockServer(s *grpc.Server, srv GripmockServer) {
	s.RegisterService(&_Gripmock_serviceDesc, srv)
}

func _Gripmock_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripmockServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.Gripmock/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripmockServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gripmock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simple.Gripmock",
	HandlerType: (*GripmockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Gripmock_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simple.proto",
}
