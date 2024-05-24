// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: proto/exchange_rate.proto

package proto

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

type GetRatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRatesRequest) Reset() {
	*x = GetRatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesRequest) ProtoMessage() {}

func (x *GetRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatesRequest.ProtoReflect.Descriptor instead.
func (*GetRatesRequest) Descriptor() ([]byte, []int) {
	return file_proto_exchange_rate_proto_rawDescGZIP(), []int{0}
}

type GetRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ask       float64 `protobuf:"fixed64,1,opt,name=ask,proto3" json:"ask,omitempty"`
	Bid       float64 `protobuf:"fixed64,2,opt,name=bid,proto3" json:"bid,omitempty"`
	Timestamp string  `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetRatesResponse) Reset() {
	*x = GetRatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_rate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesResponse) ProtoMessage() {}

func (x *GetRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_rate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatesResponse.ProtoReflect.Descriptor instead.
func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return file_proto_exchange_rate_proto_rawDescGZIP(), []int{1}
}

func (x *GetRatesResponse) GetAsk() float64 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *GetRatesResponse) GetBid() float64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *GetRatesResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type HealthcheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthcheckRequest) Reset() {
	*x = HealthcheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_rate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthcheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthcheckRequest) ProtoMessage() {}

func (x *HealthcheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_rate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthcheckRequest.ProtoReflect.Descriptor instead.
func (*HealthcheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_exchange_rate_proto_rawDescGZIP(), []int{2}
}

type HealthcheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HealthcheckResponse) Reset() {
	*x = HealthcheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exchange_rate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthcheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthcheckResponse) ProtoMessage() {}

func (x *HealthcheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exchange_rate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthcheckResponse.ProtoReflect.Descriptor instead.
func (*HealthcheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_exchange_rate_proto_rawDescGZIP(), []int{3}
}

func (x *HealthcheckResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_exchange_rate_proto protoreflect.FileDescriptor

var file_proto_exchange_rate_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x61, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x14, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x13, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xb8, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x21, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2d,
	0x72, 0x61, 0x74, 0x65, 0x2d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_exchange_rate_proto_rawDescOnce sync.Once
	file_proto_exchange_rate_proto_rawDescData = file_proto_exchange_rate_proto_rawDesc
)

func file_proto_exchange_rate_proto_rawDescGZIP() []byte {
	file_proto_exchange_rate_proto_rawDescOnce.Do(func() {
		file_proto_exchange_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_exchange_rate_proto_rawDescData)
	})
	return file_proto_exchange_rate_proto_rawDescData
}

var file_proto_exchange_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_exchange_rate_proto_goTypes = []interface{}{
	(*GetRatesRequest)(nil),     // 0: exchange_rate.GetRatesRequest
	(*GetRatesResponse)(nil),    // 1: exchange_rate.GetRatesResponse
	(*HealthcheckRequest)(nil),  // 2: exchange_rate.HealthcheckRequest
	(*HealthcheckResponse)(nil), // 3: exchange_rate.HealthcheckResponse
}
var file_proto_exchange_rate_proto_depIdxs = []int32{
	0, // 0: exchange_rate.ExchangeRateService.GetRates:input_type -> exchange_rate.GetRatesRequest
	2, // 1: exchange_rate.ExchangeRateService.Healthcheck:input_type -> exchange_rate.HealthcheckRequest
	1, // 2: exchange_rate.ExchangeRateService.GetRates:output_type -> exchange_rate.GetRatesResponse
	3, // 3: exchange_rate.ExchangeRateService.Healthcheck:output_type -> exchange_rate.HealthcheckResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_exchange_rate_proto_init() }
func file_proto_exchange_rate_proto_init() {
	if File_proto_exchange_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_exchange_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatesRequest); i {
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
		file_proto_exchange_rate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatesResponse); i {
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
		file_proto_exchange_rate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthcheckRequest); i {
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
		file_proto_exchange_rate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthcheckResponse); i {
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
			RawDescriptor: file_proto_exchange_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_exchange_rate_proto_goTypes,
		DependencyIndexes: file_proto_exchange_rate_proto_depIdxs,
		MessageInfos:      file_proto_exchange_rate_proto_msgTypes,
	}.Build()
	File_proto_exchange_rate_proto = out.File
	file_proto_exchange_rate_proto_rawDesc = nil
	file_proto_exchange_rate_proto_goTypes = nil
	file_proto_exchange_rate_proto_depIdxs = nil
}
