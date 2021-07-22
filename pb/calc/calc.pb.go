// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: calc.proto

package calc

import (
	context "context"
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

// CalcRequest 包含了兩個數字，將會傳送至計算服務並對兩個數字進行計算。
type CalcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberA int32 `protobuf:"varint,1,opt,name=number_a,json=numberA,proto3" json:"number_a,omitempty"`
	NumberB int32 `protobuf:"varint,2,opt,name=number_b,json=numberB,proto3" json:"number_b,omitempty"`
}

func (x *CalcRequest) Reset() {
	*x = CalcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRequest) ProtoMessage() {}

func (x *CalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRequest.ProtoReflect.Descriptor instead.
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return file_calc_proto_rawDescGZIP(), []int{0}
}

func (x *CalcRequest) GetNumberA() int32 {
	if x != nil {
		return x.NumberA
	}
	return 0
}

func (x *CalcRequest) GetNumberB() int32 {
	if x != nil {
		return x.NumberB
	}
	return 0
}

// CalcReply 是計算結果，將會回傳給客戶端。
type CalcReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalcReply) Reset() {
	*x = CalcReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcReply) ProtoMessage() {}

func (x *CalcReply) ProtoReflect() protoreflect.Message {
	mi := &file_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcReply.ProtoReflect.Descriptor instead.
func (*CalcReply) Descriptor() ([]byte, []int) {
	return file_calc_proto_rawDescGZIP(), []int{1}
}

func (x *CalcReply) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calc_proto protoreflect.FileDescriptor

var file_calc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x43, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x41, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x22, 0x23, 0x0a,
	0x09, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0x46, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x38, 0x0a, 0x04, 0x50, 0x6c, 0x75, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x3b, 0x63, 0x61, 0x6c, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calc_proto_rawDescOnce sync.Once
	file_calc_proto_rawDescData = file_calc_proto_rawDesc
)

func file_calc_proto_rawDescGZIP() []byte {
	file_calc_proto_rawDescOnce.Do(func() {
		file_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calc_proto_rawDescData)
	})
	return file_calc_proto_rawDescData
}

var file_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calc_proto_goTypes = []interface{}{
	(*CalcRequest)(nil), // 0: calculator.CalcRequest
	(*CalcReply)(nil),   // 1: calculator.CalcReply
}
var file_calc_proto_depIdxs = []int32{
	0, // 0: calculator.Calculator.Plus:input_type -> calculator.CalcRequest
	1, // 1: calculator.Calculator.Plus:output_type -> calculator.CalcReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calc_proto_init() }
func file_calc_proto_init() {
	if File_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcRequest); i {
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
		file_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcReply); i {
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
			RawDescriptor: file_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_proto_goTypes,
		DependencyIndexes: file_calc_proto_depIdxs,
		MessageInfos:      file_calc_proto_msgTypes,
	}.Build()
	File_calc_proto = out.File
	file_calc_proto_rawDesc = nil
	file_calc_proto_goTypes = nil
	file_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	// Plus 會接收 CalcRequest 資料作加總，最終會回傳 CalcReply。
	Plus(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcReply, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Plus(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcReply, error) {
	out := new(CalcReply)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Plus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	// Plus 會接收 CalcRequest 資料作加總，最終會回傳 CalcReply。
	Plus(context.Context, *CalcRequest) (*CalcReply, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Plus(context.Context, *CalcRequest) (*CalcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Plus not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Plus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Plus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Plus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Plus(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plus",
			Handler:    _Calculator_Plus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}