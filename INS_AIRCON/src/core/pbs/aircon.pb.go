// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: aircon.proto

package InstanceAircon

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

type AirconFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trigger bool `protobuf:"varint,1,opt,name=trigger,proto3" json:"trigger,omitempty"` // 0: off / 1: on
	// modeCooling(냉방)
	// modeEco(절전)
	// modeDry(자동 건조)
	// modeSleep(열대야 취침)
	// modePowerCooling(쿨파워)
	// modeFan(송풍)
	// modePurify(청정)
	// modeSmartCare(스마트케어)
	Mode               string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
	AirflowDirect      bool   `protobuf:"varint,3,opt,name=airflowDirect,proto3" json:"airflowDirect,omitempty"` // 0: indirect airflow(간접 바람) / 1: direct airflow(직접 바람)
	FanSpeed           int64  `protobuf:"varint,4,opt,name=fanSpeed,proto3" json:"fanSpeed,omitempty"`
	BrightnessScreen   int64  `protobuf:"varint,5,opt,name=brightnessScreen,proto3" json:"brightnessScreen,omitempty"`
	ObjTemperature     int64  `protobuf:"varint,6,opt,name=objTemperature,proto3" json:"objTemperature,omitempty"`
	StartWakeupTimer   bool   `protobuf:"varint,7,opt,name=startWakeupTimer,proto3" json:"startWakeupTimer,omitempty"`
	StartShutdownTimer bool   `protobuf:"varint,8,opt,name=startShutdownTimer,proto3" json:"startShutdownTimer,omitempty"`
	StopWakeupTimer    bool   `protobuf:"varint,9,opt,name=stopWakeupTimer,proto3" json:"stopWakeupTimer,omitempty"`
	StopShutdownTimer  bool   `protobuf:"varint,10,opt,name=stopShutdownTimer,proto3" json:"stopShutdownTimer,omitempty"`
	WakeupTime         int64  `protobuf:"varint,11,opt,name=wakeupTime,proto3" json:"wakeupTime,omitempty"`
	ShutdownTime       int64  `protobuf:"varint,12,opt,name=shutdownTime,proto3" json:"shutdownTime,omitempty"`
}

func (x *AirconFrame) Reset() {
	*x = AirconFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aircon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirconFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirconFrame) ProtoMessage() {}

func (x *AirconFrame) ProtoReflect() protoreflect.Message {
	mi := &file_aircon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirconFrame.ProtoReflect.Descriptor instead.
func (*AirconFrame) Descriptor() ([]byte, []int) {
	return file_aircon_proto_rawDescGZIP(), []int{0}
}

func (x *AirconFrame) GetTrigger() bool {
	if x != nil {
		return x.Trigger
	}
	return false
}

func (x *AirconFrame) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *AirconFrame) GetAirflowDirect() bool {
	if x != nil {
		return x.AirflowDirect
	}
	return false
}

func (x *AirconFrame) GetFanSpeed() int64 {
	if x != nil {
		return x.FanSpeed
	}
	return 0
}

func (x *AirconFrame) GetBrightnessScreen() int64 {
	if x != nil {
		return x.BrightnessScreen
	}
	return 0
}

func (x *AirconFrame) GetObjTemperature() int64 {
	if x != nil {
		return x.ObjTemperature
	}
	return 0
}

func (x *AirconFrame) GetStartWakeupTimer() bool {
	if x != nil {
		return x.StartWakeupTimer
	}
	return false
}

func (x *AirconFrame) GetStartShutdownTimer() bool {
	if x != nil {
		return x.StartShutdownTimer
	}
	return false
}

func (x *AirconFrame) GetStopWakeupTimer() bool {
	if x != nil {
		return x.StopWakeupTimer
	}
	return false
}

func (x *AirconFrame) GetStopShutdownTimer() bool {
	if x != nil {
		return x.StopShutdownTimer
	}
	return false
}

func (x *AirconFrame) GetWakeupTime() int64 {
	if x != nil {
		return x.WakeupTime
	}
	return 0
}

func (x *AirconFrame) GetShutdownTime() int64 {
	if x != nil {
		return x.ShutdownTime
	}
	return 0
}

type AirconRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AirconRes) Reset() {
	*x = AirconRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aircon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirconRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirconRes) ProtoMessage() {}

func (x *AirconRes) ProtoReflect() protoreflect.Message {
	mi := &file_aircon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirconRes.ProtoReflect.Descriptor instead.
func (*AirconRes) Descriptor() ([]byte, []int) {
	return file_aircon_proto_rawDescGZIP(), []int{1}
}

func (x *AirconRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *AirconRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_aircon_proto protoreflect.FileDescriptor

var file_aircon_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x48, 0x4f, 0x4c, 0x4d, 0x45, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xc9,
	0x03, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x61, 0x69, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x69, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x2a,
	0x0a, 0x10, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x62,
	0x6a, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x6f, 0x62, 0x6a, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x57, 0x61, 0x6b, 0x65, 0x75,
	0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x57, 0x61, 0x6b, 0x65, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x70, 0x57, 0x61, 0x6b, 0x65, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x70, 0x57, 0x61, 0x6b,
	0x65, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x74, 0x6f, 0x70,
	0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x70, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x6b, 0x65, 0x75, 0x70,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x61, 0x6b, 0x65,
	0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f,
	0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x09, 0x61, 0x69,
	0x72, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x4f, 0x0a, 0x06, 0x41, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x12,
	0x45, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x2e, 0x48, 0x4f, 0x4c, 0x4d, 0x45, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x61, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x1a, 0x19, 0x2e, 0x48, 0x4f,
	0x4c, 0x4d, 0x45, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x69, 0x72,
	0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x19, 0x5a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x69, 0x72, 0x63, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aircon_proto_rawDescOnce sync.Once
	file_aircon_proto_rawDescData = file_aircon_proto_rawDesc
)

func file_aircon_proto_rawDescGZIP() []byte {
	file_aircon_proto_rawDescOnce.Do(func() {
		file_aircon_proto_rawDescData = protoimpl.X.CompressGZIP(file_aircon_proto_rawDescData)
	})
	return file_aircon_proto_rawDescData
}

var file_aircon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aircon_proto_goTypes = []interface{}{
	(*AirconFrame)(nil), // 0: HOLME_instance.airconFrame
	(*AirconRes)(nil),   // 1: HOLME_instance.airconRes
}
var file_aircon_proto_depIdxs = []int32{
	0, // 0: HOLME_instance.Aircon.HandleFrame:input_type -> HOLME_instance.airconFrame
	1, // 1: HOLME_instance.Aircon.HandleFrame:output_type -> HOLME_instance.airconRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aircon_proto_init() }
func file_aircon_proto_init() {
	if File_aircon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aircon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirconFrame); i {
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
		file_aircon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirconRes); i {
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
			RawDescriptor: file_aircon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aircon_proto_goTypes,
		DependencyIndexes: file_aircon_proto_depIdxs,
		MessageInfos:      file_aircon_proto_msgTypes,
	}.Build()
	File_aircon_proto = out.File
	file_aircon_proto_rawDesc = nil
	file_aircon_proto_goTypes = nil
	file_aircon_proto_depIdxs = nil
}
