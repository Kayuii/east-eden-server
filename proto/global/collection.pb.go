// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: collection.proto

package global

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

type C2S_CollectionActive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int32 `protobuf:"varint,1,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
}

func (x *C2S_CollectionActive) Reset() {
	*x = C2S_CollectionActive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CollectionActive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CollectionActive) ProtoMessage() {}

func (x *C2S_CollectionActive) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CollectionActive.ProtoReflect.Descriptor instead.
func (*C2S_CollectionActive) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_CollectionActive) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type C2S_CollectionStarup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int32 `protobuf:"varint,1,opt,name=TypeId,proto3" json:"TypeId,omitempty"` // 收集品typeid
}

func (x *C2S_CollectionStarup) Reset() {
	*x = C2S_CollectionStarup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CollectionStarup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CollectionStarup) ProtoMessage() {}

func (x *C2S_CollectionStarup) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CollectionStarup.ProtoReflect.Descriptor instead.
func (*C2S_CollectionStarup) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{1}
}

func (x *C2S_CollectionStarup) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type C2S_CollectionWakeup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int32 `protobuf:"varint,1,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
}

func (x *C2S_CollectionWakeup) Reset() {
	*x = C2S_CollectionWakeup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CollectionWakeup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CollectionWakeup) ProtoMessage() {}

func (x *C2S_CollectionWakeup) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CollectionWakeup.ProtoReflect.Descriptor instead.
func (*C2S_CollectionWakeup) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{2}
}

func (x *C2S_CollectionWakeup) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type S2C_CollectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Collection `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *S2C_CollectionInfo) Reset() {
	*x = S2C_CollectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_CollectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_CollectionInfo) ProtoMessage() {}

func (x *S2C_CollectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_CollectionInfo.ProtoReflect.Descriptor instead.
func (*S2C_CollectionInfo) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_CollectionInfo) GetInfo() *Collection {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_collection_proto protoreflect.FileDescriptor

var file_collection_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x32, 0x53, 0x5f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x32, 0x53, 0x5f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x75, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x32, 0x53, 0x5f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x6b, 0x65, 0x75, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x12, 0x53, 0x32, 0x43, 0x5f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0x32, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0xaa, 0x02, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collection_proto_rawDescOnce sync.Once
	file_collection_proto_rawDescData = file_collection_proto_rawDesc
)

func file_collection_proto_rawDescGZIP() []byte {
	file_collection_proto_rawDescOnce.Do(func() {
		file_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_collection_proto_rawDescData)
	})
	return file_collection_proto_rawDescData
}

var file_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_collection_proto_goTypes = []interface{}{
	(*C2S_CollectionActive)(nil), // 0: proto.C2S_CollectionActive
	(*C2S_CollectionStarup)(nil), // 1: proto.C2S_CollectionStarup
	(*C2S_CollectionWakeup)(nil), // 2: proto.C2S_CollectionWakeup
	(*S2C_CollectionInfo)(nil),   // 3: proto.S2C_CollectionInfo
	(*Collection)(nil),           // 4: proto.Collection
}
var file_collection_proto_depIdxs = []int32{
	4, // 0: proto.S2C_CollectionInfo.info:type_name -> proto.Collection
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_collection_proto_init() }
func file_collection_proto_init() {
	if File_collection_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CollectionActive); i {
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
		file_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CollectionStarup); i {
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
		file_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CollectionWakeup); i {
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
		file_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_CollectionInfo); i {
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
			RawDescriptor: file_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_collection_proto_goTypes,
		DependencyIndexes: file_collection_proto_depIdxs,
		MessageInfos:      file_collection_proto_msgTypes,
	}.Build()
	File_collection_proto = out.File
	file_collection_proto_rawDesc = nil
	file_collection_proto_goTypes = nil
	file_collection_proto_depIdxs = nil
}
