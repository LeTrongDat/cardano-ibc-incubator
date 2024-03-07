// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: ibc/applications/fee/v1/ack.proto

package types

import (
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// IncentivizedAcknowledgement is the acknowledgement format to be used by applications wrapped in the fee middleware
type IncentivizedAcknowledgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the underlying app acknowledgement bytes
	AppAcknowledgement []byte `protobuf:"bytes,1,opt,name=app_acknowledgement,json=appAcknowledgement,proto3" json:"app_acknowledgement,omitempty"`
	// the relayer address which submits the recv packet message
	ForwardRelayerAddress string `protobuf:"bytes,2,opt,name=forward_relayer_address,json=forwardRelayerAddress,proto3" json:"forward_relayer_address,omitempty"`
	// success flag of the base application callback
	UnderlyingAppSuccess bool `protobuf:"varint,3,opt,name=underlying_app_success,json=underlyingAppSuccess,proto3" json:"underlying_app_success,omitempty"`
}

func (x *IncentivizedAcknowledgement) Reset() {
	*x = IncentivizedAcknowledgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_fee_v1_ack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncentivizedAcknowledgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncentivizedAcknowledgement) ProtoMessage() {}

func (x *IncentivizedAcknowledgement) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_fee_v1_ack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncentivizedAcknowledgement.ProtoReflect.Descriptor instead.
func (*IncentivizedAcknowledgement) Descriptor() ([]byte, []int) {
	return file_ibc_applications_fee_v1_ack_proto_rawDescGZIP(), []int{0}
}

func (x *IncentivizedAcknowledgement) GetAppAcknowledgement() []byte {
	if x != nil {
		return x.AppAcknowledgement
	}
	return nil
}

func (x *IncentivizedAcknowledgement) GetForwardRelayerAddress() string {
	if x != nil {
		return x.ForwardRelayerAddress
	}
	return ""
}

func (x *IncentivizedAcknowledgement) GetUnderlyingAppSuccess() bool {
	if x != nil {
		return x.UnderlyingAppSuccess
	}
	return false
}

var File_ibc_applications_fee_v1_ack_proto protoreflect.FileDescriptor

var file_ibc_applications_fee_v1_ack_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x62, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x1b, 0x49, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x69,
	0x7a, 0x65, 0x64, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x13, 0x61, 0x70, 0x70, 0x5f, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x1e, 0xf2, 0xde, 0x1f, 0x1a, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x61, 0x70, 0x70, 0x5f, 0x61,
	0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x52,
	0x12, 0x61, 0x70, 0x70, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x17, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xf2, 0xde, 0x1f, 0x1e, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x52, 0x15, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x58, 0x0a, 0x16, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70,
	0x70, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x22, 0xf2, 0xde, 0x1f, 0x1e, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x75, 0x6e, 0x64, 0x65, 0x72,
	0x6c, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x6c, 0x22, 0x52, 0x14, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x41,
	0x70, 0x70, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x69,
	0x62, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x37, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x32, 0x39, 0x2d, 0x66, 0x65, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ibc_applications_fee_v1_ack_proto_rawDescOnce sync.Once
	file_ibc_applications_fee_v1_ack_proto_rawDescData = file_ibc_applications_fee_v1_ack_proto_rawDesc
)

func file_ibc_applications_fee_v1_ack_proto_rawDescGZIP() []byte {
	file_ibc_applications_fee_v1_ack_proto_rawDescOnce.Do(func() {
		file_ibc_applications_fee_v1_ack_proto_rawDescData = protoimpl.X.CompressGZIP(file_ibc_applications_fee_v1_ack_proto_rawDescData)
	})
	return file_ibc_applications_fee_v1_ack_proto_rawDescData
}

var file_ibc_applications_fee_v1_ack_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ibc_applications_fee_v1_ack_proto_goTypes = []interface{}{
	(*IncentivizedAcknowledgement)(nil), // 0: ibc.applications.fee.v1.IncentivizedAcknowledgement
}
var file_ibc_applications_fee_v1_ack_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ibc_applications_fee_v1_ack_proto_init() }
func file_ibc_applications_fee_v1_ack_proto_init() {
	if File_ibc_applications_fee_v1_ack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ibc_applications_fee_v1_ack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncentivizedAcknowledgement); i {
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
			RawDescriptor: file_ibc_applications_fee_v1_ack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ibc_applications_fee_v1_ack_proto_goTypes,
		DependencyIndexes: file_ibc_applications_fee_v1_ack_proto_depIdxs,
		MessageInfos:      file_ibc_applications_fee_v1_ack_proto_msgTypes,
	}.Build()
	File_ibc_applications_fee_v1_ack_proto = out.File
	file_ibc_applications_fee_v1_ack_proto_rawDesc = nil
	file_ibc_applications_fee_v1_ack_proto_goTypes = nil
	file_ibc_applications_fee_v1_ack_proto_depIdxs = nil
}
