// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: waterDispenser.proto

package InstanceWaterDispenser

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

type WaterDispenserFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TriggerReminder bool `protobuf:"varint,1,opt,name=triggerReminder,proto3" json:"triggerReminder,omitempty"` // 0: before medication reminder / 1: after/during medication reminder
	TriggerWater    bool `protobuf:"varint,2,opt,name=triggerWater,proto3" json:"triggerWater,omitempty"`       // 0: water not dispensed 1: water dispensed
}

func (x *WaterDispenserFrame) Reset() {
	*x = WaterDispenserFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waterDispenser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaterDispenserFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaterDispenserFrame) ProtoMessage() {}

func (x *WaterDispenserFrame) ProtoReflect() protoreflect.Message {
	mi := &file_waterDispenser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaterDispenserFrame.ProtoReflect.Descriptor instead.
func (*WaterDispenserFrame) Descriptor() ([]byte, []int) {
	return file_waterDispenser_proto_rawDescGZIP(), []int{0}
}

func (x *WaterDispenserFrame) GetTriggerReminder() bool {
	if x != nil {
		return x.TriggerReminder
	}
	return false
}

func (x *WaterDispenserFrame) GetTriggerWater() bool {
	if x != nil {
		return x.TriggerWater
	}
	return false
}

type WaterDispenserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *WaterDispenserRes) Reset() {
	*x = WaterDispenserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waterDispenser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaterDispenserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaterDispenserRes) ProtoMessage() {}

func (x *WaterDispenserRes) ProtoReflect() protoreflect.Message {
	mi := &file_waterDispenser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaterDispenserRes.ProtoReflect.Descriptor instead.
func (*WaterDispenserRes) Descriptor() ([]byte, []int) {
	return file_waterDispenser_proto_rawDescGZIP(), []int{1}
}

func (x *WaterDispenserRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *WaterDispenserRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_waterDispenser_proto protoreflect.FileDescriptor

var file_waterDispenser_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x61, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x48, 0x4f, 0x4c, 0x4d, 0x45, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x63, 0x0a, 0x13, 0x77, 0x61, 0x74, 0x65, 0x72, 0x44,
	0x69, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x57, 0x61, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x57, 0x61, 0x74, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x11, 0x77,
	0x61, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x67,
	0x0a, 0x0e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72,
	0x12, 0x55, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x2e, 0x48, 0x4f, 0x4c, 0x4d, 0x45, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x1a, 0x21, 0x2e, 0x48, 0x4f, 0x4c, 0x4d, 0x45, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x72, 0x63, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x61, 0x74, 0x65,
	0x72, 0x44, 0x69, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_waterDispenser_proto_rawDescOnce sync.Once
	file_waterDispenser_proto_rawDescData = file_waterDispenser_proto_rawDesc
)

func file_waterDispenser_proto_rawDescGZIP() []byte {
	file_waterDispenser_proto_rawDescOnce.Do(func() {
		file_waterDispenser_proto_rawDescData = protoimpl.X.CompressGZIP(file_waterDispenser_proto_rawDescData)
	})
	return file_waterDispenser_proto_rawDescData
}

var file_waterDispenser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_waterDispenser_proto_goTypes = []interface{}{
	(*WaterDispenserFrame)(nil), // 0: HOLME_instance.waterDispenserFrame
	(*WaterDispenserRes)(nil),   // 1: HOLME_instance.waterDispenserRes
}
var file_waterDispenser_proto_depIdxs = []int32{
	0, // 0: HOLME_instance.WaterDispenser.HandleFrame:input_type -> HOLME_instance.waterDispenserFrame
	1, // 1: HOLME_instance.WaterDispenser.HandleFrame:output_type -> HOLME_instance.waterDispenserRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_waterDispenser_proto_init() }
func file_waterDispenser_proto_init() {
	if File_waterDispenser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waterDispenser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaterDispenserFrame); i {
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
		file_waterDispenser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaterDispenserRes); i {
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
			RawDescriptor: file_waterDispenser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_waterDispenser_proto_goTypes,
		DependencyIndexes: file_waterDispenser_proto_depIdxs,
		MessageInfos:      file_waterDispenser_proto_msgTypes,
	}.Build()
	File_waterDispenser_proto = out.File
	file_waterDispenser_proto_rawDesc = nil
	file_waterDispenser_proto_goTypes = nil
	file_waterDispenser_proto_depIdxs = nil
}
