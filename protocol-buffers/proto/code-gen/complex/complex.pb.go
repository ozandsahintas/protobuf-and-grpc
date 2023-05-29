// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/code-gen/complex/complex.proto

package complex

import (
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

type Complex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Other  *Other   `protobuf:"bytes,1,opt,name=other,proto3" json:"other,omitempty"`
	Others []*Other `protobuf:"bytes,2,rep,name=others,proto3" json:"others,omitempty"`
}

func (x *Complex) Reset() {
	*x = Complex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_code_gen_complex_complex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Complex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Complex) ProtoMessage() {}

func (x *Complex) ProtoReflect() protoreflect.Message {
	mi := &file_proto_code_gen_complex_complex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Complex.ProtoReflect.Descriptor instead.
func (*Complex) Descriptor() ([]byte, []int) {
	return file_proto_code_gen_complex_complex_proto_rawDescGZIP(), []int{0}
}

func (x *Complex) GetOther() *Other {
	if x != nil {
		return x.Other
	}
	return nil
}

func (x *Complex) GetOthers() []*Other {
	if x != nil {
		return x.Others
	}
	return nil
}

type Other struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Other) Reset() {
	*x = Other{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_code_gen_complex_complex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Other) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Other) ProtoMessage() {}

func (x *Other) ProtoReflect() protoreflect.Message {
	mi := &file_proto_code_gen_complex_complex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Other.ProtoReflect.Descriptor instead.
func (*Other) Descriptor() ([]byte, []int) {
	return file_proto_code_gen_complex_complex_proto_rawDescGZIP(), []int{1}
}

func (x *Other) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Other) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_code_gen_complex_complex_proto protoreflect.FileDescriptor

var file_proto_code_gen_complex_complex_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x22, 0x67, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x12, 0x2c, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73,
	0x22, 0x2b, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x18, 0x5a,
	0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_code_gen_complex_complex_proto_rawDescOnce sync.Once
	file_proto_code_gen_complex_complex_proto_rawDescData = file_proto_code_gen_complex_complex_proto_rawDesc
)

func file_proto_code_gen_complex_complex_proto_rawDescGZIP() []byte {
	file_proto_code_gen_complex_complex_proto_rawDescOnce.Do(func() {
		file_proto_code_gen_complex_complex_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_code_gen_complex_complex_proto_rawDescData)
	})
	return file_proto_code_gen_complex_complex_proto_rawDescData
}

var file_proto_code_gen_complex_complex_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_code_gen_complex_complex_proto_goTypes = []interface{}{
	(*Complex)(nil), // 0: example.complex.Complex
	(*Other)(nil),   // 1: example.complex.Other
}
var file_proto_code_gen_complex_complex_proto_depIdxs = []int32{
	1, // 0: example.complex.Complex.other:type_name -> example.complex.Other
	1, // 1: example.complex.Complex.others:type_name -> example.complex.Other
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_code_gen_complex_complex_proto_init() }
func file_proto_code_gen_complex_complex_proto_init() {
	if File_proto_code_gen_complex_complex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_code_gen_complex_complex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Complex); i {
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
		file_proto_code_gen_complex_complex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Other); i {
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
			RawDescriptor: file_proto_code_gen_complex_complex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_code_gen_complex_complex_proto_goTypes,
		DependencyIndexes: file_proto_code_gen_complex_complex_proto_depIdxs,
		MessageInfos:      file_proto_code_gen_complex_complex_proto_msgTypes,
	}.Build()
	File_proto_code_gen_complex_complex_proto = out.File
	file_proto_code_gen_complex_complex_proto_rawDesc = nil
	file_proto_code_gen_complex_complex_proto_goTypes = nil
	file_proto_code_gen_complex_complex_proto_depIdxs = nil
}
