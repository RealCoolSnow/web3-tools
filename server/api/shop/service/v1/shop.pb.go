// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/shop/service/v1/shop.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户手机号码 （在此注释生成的swagger.json倒入apifox可自动生成注释）
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	// 密码
	Pass string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	// 昵称
	NickName string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	// 年龄
	Age int64 `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *RegisterRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *RegisterRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Age      int64  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegisterReply) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterReply) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *RegisterReply) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Pass   string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LoginRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // 这里需要注意int64前端接收到的是字符串，int32才是数值
	Data *LoginReply_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Msg  string           `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginReply) GetData() *LoginReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoginReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{4}
}

type GetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data *GetUserReply_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetUserReply) GetData() *GetUserReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type DemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                             // 从URL上映射
	NickName string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"` // 从body映射
}

func (x *DemoRequest) Reset() {
	*x = DemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoRequest) ProtoMessage() {}

func (x *DemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoRequest.ProtoReflect.Descriptor instead.
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{6}
}

func (x *DemoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DemoRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

type DemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
}

func (x *DemoResponse) Reset() {
	*x = DemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoResponse) ProtoMessage() {}

func (x *DemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoResponse.ProtoReflect.Descriptor instead.
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{7}
}

func (x *DemoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DemoResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

type LoginReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginReply_Data) Reset() {
	*x = LoginReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply_Data) ProtoMessage() {}

func (x *LoginReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply_Data.ProtoReflect.Descriptor instead.
func (*LoginReply_Data) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{3, 0}
}

func (x *LoginReply_Data) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetUserReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetUserReply_Data) Reset() {
	*x = GetUserReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_service_v1_shop_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply_Data) ProtoMessage() {}

func (x *GetUserReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_service_v1_shop_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply_Data.ProtoReflect.Descriptor instead.
func (*GetUserReply_Data) Descriptor() ([]byte, []int) {
	return file_api_shop_service_v1_shop_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetUserReply_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_shop_service_v1_shop_proto protoreflect.FileDescriptor

var file_api_shop_service_v1_shop_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x0b, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x06, 0x18, 0x12, 0x52, 0x04, 0x70, 0x61, 0x73,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x22, 0x66, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98,
	0x01, 0x0b, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x61,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x06, 0x18, 0x12, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0a, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x1c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x3b, 0x0a, 0x0c, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xa8, 0x03,
	0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x6d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x21,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x12, 0x6b, 0x0a, 0x04, 0x44,
	0x65, 0x6d, 0x6f, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x1a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x09, 0x6e,
	0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x63, 0x61, 0x73, 0x73,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_shop_service_v1_shop_proto_rawDescOnce sync.Once
	file_api_shop_service_v1_shop_proto_rawDescData = file_api_shop_service_v1_shop_proto_rawDesc
)

func file_api_shop_service_v1_shop_proto_rawDescGZIP() []byte {
	file_api_shop_service_v1_shop_proto_rawDescOnce.Do(func() {
		file_api_shop_service_v1_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_shop_service_v1_shop_proto_rawDescData)
	})
	return file_api_shop_service_v1_shop_proto_rawDescData
}

var file_api_shop_service_v1_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_shop_service_v1_shop_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),   // 0: api.shop.service.v1.RegisterRequest
	(*RegisterReply)(nil),     // 1: api.shop.service.v1.RegisterReply
	(*LoginRequest)(nil),      // 2: api.shop.service.v1.LoginRequest
	(*LoginReply)(nil),        // 3: api.shop.service.v1.LoginReply
	(*GetUserRequest)(nil),    // 4: api.shop.service.v1.GetUserRequest
	(*GetUserReply)(nil),      // 5: api.shop.service.v1.GetUserReply
	(*DemoRequest)(nil),       // 6: api.shop.service.v1.DemoRequest
	(*DemoResponse)(nil),      // 7: api.shop.service.v1.DemoResponse
	(*LoginReply_Data)(nil),   // 8: api.shop.service.v1.LoginReply.Data
	(*GetUserReply_Data)(nil), // 9: api.shop.service.v1.GetUserReply.Data
}
var file_api_shop_service_v1_shop_proto_depIdxs = []int32{
	8, // 0: api.shop.service.v1.LoginReply.data:type_name -> api.shop.service.v1.LoginReply.Data
	9, // 1: api.shop.service.v1.GetUserReply.data:type_name -> api.shop.service.v1.GetUserReply.Data
	0, // 2: api.shop.service.v1.Shop.Register:input_type -> api.shop.service.v1.RegisterRequest
	2, // 3: api.shop.service.v1.Shop.Login:input_type -> api.shop.service.v1.LoginRequest
	4, // 4: api.shop.service.v1.Shop.GetUser:input_type -> api.shop.service.v1.GetUserRequest
	6, // 5: api.shop.service.v1.Shop.Demo:input_type -> api.shop.service.v1.DemoRequest
	1, // 6: api.shop.service.v1.Shop.Register:output_type -> api.shop.service.v1.RegisterReply
	3, // 7: api.shop.service.v1.Shop.Login:output_type -> api.shop.service.v1.LoginReply
	5, // 8: api.shop.service.v1.Shop.GetUser:output_type -> api.shop.service.v1.GetUserReply
	7, // 9: api.shop.service.v1.Shop.Demo:output_type -> api.shop.service.v1.DemoResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_shop_service_v1_shop_proto_init() }
func file_api_shop_service_v1_shop_proto_init() {
	if File_api_shop_service_v1_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_shop_service_v1_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReply); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoRequest); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoResponse); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply_Data); i {
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
		file_api_shop_service_v1_shop_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReply_Data); i {
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
			RawDescriptor: file_api_shop_service_v1_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_shop_service_v1_shop_proto_goTypes,
		DependencyIndexes: file_api_shop_service_v1_shop_proto_depIdxs,
		MessageInfos:      file_api_shop_service_v1_shop_proto_msgTypes,
	}.Build()
	File_api_shop_service_v1_shop_proto = out.File
	file_api_shop_service_v1_shop_proto_rawDesc = nil
	file_api_shop_service_v1_shop_proto_goTypes = nil
	file_api_shop_service_v1_shop_proto_depIdxs = nil
}
