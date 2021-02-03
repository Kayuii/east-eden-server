// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: global/define.proto

package global

import (
	proto "github.com/golang/protobuf/proto"
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

////////////////////////////////////////////////
// Account
type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_define_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_global_define_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_global_define_proto_rawDescGZIP(), []int{0}
}

func (x *AccountInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

////////////////////////////////////////////////
// Att
type Att struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttType  int32 `protobuf:"varint,1,opt,name=AttType,proto3" json:"AttType,omitempty"`
	AttValue int64 `protobuf:"varint,2,opt,name=AttValue,proto3" json:"AttValue,omitempty"`
}

func (x *Att) Reset() {
	*x = Att{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_define_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Att) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Att) ProtoMessage() {}

func (x *Att) ProtoReflect() protoreflect.Message {
	mi := &file_global_define_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Att.ProtoReflect.Descriptor instead.
func (*Att) Descriptor() ([]byte, []int) {
	return file_global_define_proto_rawDescGZIP(), []int{1}
}

func (x *Att) GetAttType() int32 {
	if x != nil {
		return x.AttType
	}
	return 0
}

func (x *Att) GetAttValue() int64 {
	if x != nil {
		return x.AttValue
	}
	return 0
}

////////////////////////////////////////////////
// Hero
type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	TypeId    int32   `protobuf:"varint,2,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	Exp       int64   `protobuf:"varint,3,opt,name=Exp,proto3" json:"Exp,omitempty"`
	Level     int32   `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`
	EquipList []int64 `protobuf:"varint,5,rep,packed,name=EquipList,proto3" json:"EquipList,omitempty"` // 装备
	RuneList  []int64 `protobuf:"varint,6,rep,packed,name=RuneList,proto3" json:"RuneList,omitempty"`   // 魂石
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_define_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_global_define_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_global_define_proto_rawDescGZIP(), []int{2}
}

func (x *Hero) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Hero) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *Hero) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *Hero) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Hero) GetEquipList() []int64 {
	if x != nil {
		return x.EquipList
	}
	return nil
}

func (x *Hero) GetRuneList() []int64 {
	if x != nil {
		return x.RuneList
	}
	return nil
}

////////////////////////////////////////////////
// Item
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	TypeId     int32 `protobuf:"varint,2,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	Num        int32 `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`
	EquipObjId int64 `protobuf:"varint,4,opt,name=EquipObjId,proto3" json:"EquipObjId,omitempty"` // 装备者objid
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_define_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_global_define_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_global_define_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *Item) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Item) GetEquipObjId() int64 {
	if x != nil {
		return x.EquipObjId
	}
	return 0
}

////////////////////////////////////////////////
// player
type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId int64  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Exp       int64  `protobuf:"varint,4,opt,name=exp,proto3" json:"exp,omitempty"`
	Level     int32  `protobuf:"varint,5,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_define_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_global_define_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_global_define_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlayerInfo) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *PlayerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerInfo) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *PlayerInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

////////////////////////////////////////////////
// Rune
type RuneAtt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttType  int32 `protobuf:"varint,1,opt,name=AttType,proto3" json:"AttType,omitempty"`
	AttValue int64 `protobuf:"varint,2,opt,name=AttValue,proto3" json:"AttValue,omitempty"`
}

func (x *RuneAtt) Reset() {
	*x = RuneAtt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_define_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuneAtt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuneAtt) ProtoMessage() {}

func (x *RuneAtt) ProtoReflect() protoreflect.Message {
	mi := &file_global_define_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuneAtt.ProtoReflect.Descriptor instead.
func (*RuneAtt) Descriptor() ([]byte, []int) {
	return file_global_define_proto_rawDescGZIP(), []int{5}
}

func (x *RuneAtt) GetAttType() int32 {
	if x != nil {
		return x.AttType
	}
	return 0
}

func (x *RuneAtt) GetAttValue() int64 {
	if x != nil {
		return x.AttValue
	}
	return 0
}

var File_global_define_proto protoreflect.FileDescriptor

var file_global_define_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x22, 0x47, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x3b, 0x0a, 0x03, 0x41, 0x74, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x41, 0x74, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x41, 0x74, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x74, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x45, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x75,
	0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x52, 0x75,
	0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x3f, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x65, 0x41, 0x74, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x74, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x41,
	0x74, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x74, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x36, 0x5a, 0x2b, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0xaa, 0x02, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_global_define_proto_rawDescOnce sync.Once
	file_global_define_proto_rawDescData = file_global_define_proto_rawDesc
)

func file_global_define_proto_rawDescGZIP() []byte {
	file_global_define_proto_rawDescOnce.Do(func() {
		file_global_define_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_define_proto_rawDescData)
	})
	return file_global_define_proto_rawDescData
}

var file_global_define_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_global_define_proto_goTypes = []interface{}{
	(*AccountInfo)(nil), // 0: global.AccountInfo
	(*Att)(nil),         // 1: global.Att
	(*Hero)(nil),        // 2: global.Hero
	(*Item)(nil),        // 3: global.Item
	(*PlayerInfo)(nil),  // 4: global.PlayerInfo
	(*RuneAtt)(nil),     // 5: global.RuneAtt
}
var file_global_define_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_global_define_proto_init() }
func file_global_define_proto_init() {
	if File_global_define_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_global_define_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
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
		file_global_define_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Att); i {
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
		file_global_define_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero); i {
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
		file_global_define_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_global_define_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
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
		file_global_define_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuneAtt); i {
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
			RawDescriptor: file_global_define_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_global_define_proto_goTypes,
		DependencyIndexes: file_global_define_proto_depIdxs,
		MessageInfos:      file_global_define_proto_msgTypes,
	}.Build()
	File_global_define_proto = out.File
	file_global_define_proto_rawDesc = nil
	file_global_define_proto_goTypes = nil
	file_global_define_proto_depIdxs = nil
}
