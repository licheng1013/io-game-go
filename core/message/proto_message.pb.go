// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto_message.proto

package message

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

type ProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Merge     int64  `protobuf:"varint,1,opt,name=merge,proto3" json:"merge,omitempty"`
	Body      []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Heartbeat bool   `protobuf:"varint,3,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
	Code      int64  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Message   string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ProtoMessage) Reset() {
	*x = ProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMessage) ProtoMessage() {}

func (x *ProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMessage.ProtoReflect.Descriptor instead.
func (*ProtoMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoMessage) GetMerge() int64 {
	if x != nil {
		return x.Merge
	}
	return 0
}

func (x *ProtoMessage) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ProtoMessage) GetHeartbeat() bool {
	if x != nil {
		return x.Heartbeat
	}
	return false
}

func (x *ProtoMessage) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ProtoMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_message_proto protoreflect.FileDescriptor

var file_proto_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_message_proto_rawDescOnce sync.Once
	file_proto_message_proto_rawDescData = file_proto_message_proto_rawDesc
)

func file_proto_message_proto_rawDescGZIP() []byte {
	file_proto_message_proto_rawDescOnce.Do(func() {
		file_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_proto_rawDescData)
	})
	return file_proto_message_proto_rawDescData
}

var file_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_message_proto_goTypes = []interface{}{
	(*ProtoMessage)(nil), // 0: ProtoMessage
}
var file_proto_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_message_proto_init() }
func file_proto_message_proto_init() {
	if File_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoMessage); i {
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
			RawDescriptor: file_proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_proto_goTypes,
		DependencyIndexes: file_proto_message_proto_depIdxs,
		MessageInfos:      file_proto_message_proto_msgTypes,
	}.Build()
	File_proto_message_proto = out.File
	file_proto_message_proto_rawDesc = nil
	file_proto_message_proto_goTypes = nil
	file_proto_message_proto_depIdxs = nil
}
