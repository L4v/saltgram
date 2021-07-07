// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: email/email.proto

package premail

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

type ConfirmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ConfirmRequest) Reset() {
	*x = ConfirmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmRequest) ProtoMessage() {}

func (x *ConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmRequest.ProtoReflect.Descriptor instead.
func (*ConfirmRequest) Descriptor() ([]byte, []int) {
	return file_email_email_proto_rawDescGZIP(), []int{0}
}

func (x *ConfirmRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ConfirmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ConfirmResponse) Reset() {
	*x = ConfirmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmResponse) ProtoMessage() {}

func (x *ConfirmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmResponse.ProtoReflect.Descriptor instead.
func (*ConfirmResponse) Descriptor() ([]byte, []int) {
	return file_email_email_proto_rawDescGZIP(), []int{1}
}

func (x *ConfirmResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_email_email_proto_rawDescGZIP(), []int{2}
}

func (x *ResetRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ResetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetResponse) Reset() {
	*x = ResetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetResponse) ProtoMessage() {}

func (x *ResetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetResponse.ProtoReflect.Descriptor instead.
func (*ResetResponse) Descriptor() ([]byte, []int) {
	return file_email_email_proto_rawDescGZIP(), []int{3}
}

type ActivateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ActivateRequest) Reset() {
	*x = ActivateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_email_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateRequest) ProtoMessage() {}

func (x *ActivateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_email_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateRequest.ProtoReflect.Descriptor instead.
func (*ActivateRequest) Descriptor() ([]byte, []int) {
	return file_email_email_proto_rawDescGZIP(), []int{4}
}

func (x *ActivateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ActivateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActivateResponse) Reset() {
	*x = ActivateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_email_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateResponse) ProtoMessage() {}

func (x *ActivateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_email_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateResponse.ProtoReflect.Descriptor instead.
func (*ActivateResponse) Descriptor() ([]byte, []int) {
	return file_email_email_proto_rawDescGZIP(), []int{5}
}

type SendActivationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendActivationRequest) Reset() {
	*x = SendActivationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_email_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendActivationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendActivationRequest) ProtoMessage() {}

func (x *SendActivationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_email_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendActivationRequest.ProtoReflect.Descriptor instead.
func (*SendActivationRequest) Descriptor() ([]byte, []int) {
	return file_email_email_proto_rawDescGZIP(), []int{6}
}

func (x *SendActivationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SendActivationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendActivationResponse) Reset() {
	*x = SendActivationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_email_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendActivationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendActivationResponse) ProtoMessage() {}

func (x *SendActivationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_email_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendActivationResponse.ProtoReflect.Descriptor instead.
func (*SendActivationResponse) Descriptor() ([]byte, []int) {
	return file_email_email_proto_rawDescGZIP(), []int{7}
}

var File_email_email_proto protoreflect.FileDescriptor

var file_email_email_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x27, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x24, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xdd, 0x01, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x41, 0x0a, 0x0e, 0x53,
	0x65, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12,
	0x0d, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x0f,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x70, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_email_proto_rawDescOnce sync.Once
	file_email_email_proto_rawDescData = file_email_email_proto_rawDesc
)

func file_email_email_proto_rawDescGZIP() []byte {
	file_email_email_proto_rawDescOnce.Do(func() {
		file_email_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_email_proto_rawDescData)
	})
	return file_email_email_proto_rawDescData
}

var file_email_email_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_email_email_proto_goTypes = []interface{}{
	(*ConfirmRequest)(nil),         // 0: ConfirmRequest
	(*ConfirmResponse)(nil),        // 1: ConfirmResponse
	(*ResetRequest)(nil),           // 2: ResetRequest
	(*ResetResponse)(nil),          // 3: ResetResponse
	(*ActivateRequest)(nil),        // 4: ActivateRequest
	(*ActivateResponse)(nil),       // 5: ActivateResponse
	(*SendActivationRequest)(nil),  // 6: SendActivationRequest
	(*SendActivationResponse)(nil), // 7: SendActivationResponse
}
var file_email_email_proto_depIdxs = []int32{
	6, // 0: Email.SendActivation:input_type -> SendActivationRequest
	4, // 1: Email.Activate:input_type -> ActivateRequest
	2, // 2: Email.RequestReset:input_type -> ResetRequest
	0, // 3: Email.ConfirmReset:input_type -> ConfirmRequest
	7, // 4: Email.SendActivation:output_type -> SendActivationResponse
	5, // 5: Email.Activate:output_type -> ActivateResponse
	3, // 6: Email.RequestReset:output_type -> ResetResponse
	1, // 7: Email.ConfirmReset:output_type -> ConfirmResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_email_email_proto_init() }
func file_email_email_proto_init() {
	if File_email_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmRequest); i {
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
		file_email_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmResponse); i {
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
		file_email_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRequest); i {
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
		file_email_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetResponse); i {
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
		file_email_email_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateRequest); i {
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
		file_email_email_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateResponse); i {
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
		file_email_email_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendActivationRequest); i {
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
		file_email_email_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendActivationResponse); i {
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
			RawDescriptor: file_email_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_email_proto_goTypes,
		DependencyIndexes: file_email_email_proto_depIdxs,
		MessageInfos:      file_email_email_proto_msgTypes,
	}.Build()
	File_email_email_proto = out.File
	file_email_email_proto_rawDesc = nil
	file_email_email_proto_goTypes = nil
	file_email_email_proto_depIdxs = nil
}
