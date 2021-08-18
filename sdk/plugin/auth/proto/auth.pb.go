// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: auth/proto/auth.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CheckUserAndPassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Protocol string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	User     []byte `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"` // SFTPGo user JSON serialized
}

func (x *CheckUserAndPassRequest) Reset() {
	*x = CheckUserAndPassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAndPassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAndPassRequest) ProtoMessage() {}

func (x *CheckUserAndPassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAndPassRequest.ProtoReflect.Descriptor instead.
func (*CheckUserAndPassRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *CheckUserAndPassRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CheckUserAndPassRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CheckUserAndPassRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CheckUserAndPassRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *CheckUserAndPassRequest) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

type CheckUserAndTLSCertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	TlsCert  string `protobuf:"bytes,2,opt,name=tlsCert,proto3" json:"tlsCert,omitempty"` // tls certificate pem encoded
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Protocol string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	User     []byte `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"` // SFTPGo user JSON serialized
}

func (x *CheckUserAndTLSCertRequest) Reset() {
	*x = CheckUserAndTLSCertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAndTLSCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAndTLSCertRequest) ProtoMessage() {}

func (x *CheckUserAndTLSCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAndTLSCertRequest.ProtoReflect.Descriptor instead.
func (*CheckUserAndTLSCertRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *CheckUserAndTLSCertRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CheckUserAndTLSCertRequest) GetTlsCert() string {
	if x != nil {
		return x.TlsCert
	}
	return ""
}

func (x *CheckUserAndTLSCertRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CheckUserAndTLSCertRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *CheckUserAndTLSCertRequest) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

type CheckUserAndPublicKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PubKey   string `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Protocol string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	User     []byte `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"` // SFTPGo user JSON serialized
}

func (x *CheckUserAndPublicKeyRequest) Reset() {
	*x = CheckUserAndPublicKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAndPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAndPublicKeyRequest) ProtoMessage() {}

func (x *CheckUserAndPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAndPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*CheckUserAndPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CheckUserAndPublicKeyRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CheckUserAndPublicKeyRequest) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *CheckUserAndPublicKeyRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CheckUserAndPublicKeyRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *CheckUserAndPublicKeyRequest) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

type CheckUserAndKeyboardInteractiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Ip       string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	User     []byte `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"` // SFTPGo user JSON serialized
}

func (x *CheckUserAndKeyboardInteractiveRequest) Reset() {
	*x = CheckUserAndKeyboardInteractiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAndKeyboardInteractiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAndKeyboardInteractiveRequest) ProtoMessage() {}

func (x *CheckUserAndKeyboardInteractiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAndKeyboardInteractiveRequest.ProtoReflect.Descriptor instead.
func (*CheckUserAndKeyboardInteractiveRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CheckUserAndKeyboardInteractiveRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CheckUserAndKeyboardInteractiveRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CheckUserAndKeyboardInteractiveRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *CheckUserAndKeyboardInteractiveRequest) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

type KeyboardAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestID string   `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Username  string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password  string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Ip        string   `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Answers   []string `protobuf:"bytes,5,rep,name=answers,proto3" json:"answers,omitempty"`
	Questions []string `protobuf:"bytes,6,rep,name=questions,proto3" json:"questions,omitempty"`
	Step      int32    `protobuf:"varint,7,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *KeyboardAuthRequest) Reset() {
	*x = KeyboardAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardAuthRequest) ProtoMessage() {}

func (x *KeyboardAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardAuthRequest.ProtoReflect.Descriptor instead.
func (*KeyboardAuthRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *KeyboardAuthRequest) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *KeyboardAuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *KeyboardAuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *KeyboardAuthRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *KeyboardAuthRequest) GetAnswers() []string {
	if x != nil {
		return x.Answers
	}
	return nil
}

func (x *KeyboardAuthRequest) GetQuestions() []string {
	if x != nil {
		return x.Questions
	}
	return nil
}

func (x *KeyboardAuthRequest) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

type KeyboardAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instructions  string   `protobuf:"bytes,1,opt,name=instructions,proto3" json:"instructions,omitempty"`
	Questions     []string `protobuf:"bytes,2,rep,name=questions,proto3" json:"questions,omitempty"`
	Echos         []bool   `protobuf:"varint,3,rep,packed,name=echos,proto3" json:"echos,omitempty"`
	AuthResult    int32    `protobuf:"varint,4,opt,name=auth_result,json=authResult,proto3" json:"auth_result,omitempty"`
	CheckPassword int32    `protobuf:"varint,5,opt,name=check_password,json=checkPassword,proto3" json:"check_password,omitempty"`
}

func (x *KeyboardAuthResponse) Reset() {
	*x = KeyboardAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardAuthResponse) ProtoMessage() {}

func (x *KeyboardAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardAuthResponse.ProtoReflect.Descriptor instead.
func (*KeyboardAuthResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *KeyboardAuthResponse) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *KeyboardAuthResponse) GetQuestions() []string {
	if x != nil {
		return x.Questions
	}
	return nil
}

func (x *KeyboardAuthResponse) GetEchos() []bool {
	if x != nil {
		return x.Echos
	}
	return nil
}

func (x *KeyboardAuthResponse) GetAuthResult() int32 {
	if x != nil {
		return x.AuthResult
	}
	return 0
}

func (x *KeyboardAuthResponse) GetCheckPassword() int32 {
	if x != nil {
		return x.CheckPassword
	}
	return 0
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []byte `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"` // SFTPGo user JSON serialized
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_auth_proto_rawDescGZIP(), []int{6}
}

func (x *AuthResponse) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

var File_auth_proto_auth_proto protoreflect.FileDescriptor

var file_auth_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91,
	0x01, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x50,
	0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x6e, 0x64, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x1c, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x84, 0x01, 0x0a,
	0x26, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0xc7, 0x01, 0x0a, 0x13, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0xb6, 0x01,
	0x0a, 0x14, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x63, 0x68, 0x6f,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x08, 0x52, 0x05, 0x65, 0x63, 0x68, 0x6f, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xac, 0x03, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x47, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x13,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x54, 0x4c, 0x53, 0x43,
	0x65, 0x72, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x15, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65,
	0x0a, 0x1f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4b, 0x65,
	0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x65, 0x79,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_auth_proto_rawDescOnce sync.Once
	file_auth_proto_auth_proto_rawDescData = file_auth_proto_auth_proto_rawDesc
)

func file_auth_proto_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_auth_proto_rawDescData)
	})
	return file_auth_proto_auth_proto_rawDescData
}

var file_auth_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_auth_proto_auth_proto_goTypes = []interface{}{
	(*CheckUserAndPassRequest)(nil),                // 0: proto.CheckUserAndPassRequest
	(*CheckUserAndTLSCertRequest)(nil),             // 1: proto.CheckUserAndTLSCertRequest
	(*CheckUserAndPublicKeyRequest)(nil),           // 2: proto.CheckUserAndPublicKeyRequest
	(*CheckUserAndKeyboardInteractiveRequest)(nil), // 3: proto.CheckUserAndKeyboardInteractiveRequest
	(*KeyboardAuthRequest)(nil),                    // 4: proto.KeyboardAuthRequest
	(*KeyboardAuthResponse)(nil),                   // 5: proto.KeyboardAuthResponse
	(*AuthResponse)(nil),                           // 6: proto.AuthResponse
}
var file_auth_proto_auth_proto_depIdxs = []int32{
	0, // 0: proto.Auth.CheckUserAndPass:input_type -> proto.CheckUserAndPassRequest
	1, // 1: proto.Auth.CheckUserAndTLSCert:input_type -> proto.CheckUserAndTLSCertRequest
	2, // 2: proto.Auth.CheckUserAndPublicKey:input_type -> proto.CheckUserAndPublicKeyRequest
	3, // 3: proto.Auth.CheckUserAndKeyboardInteractive:input_type -> proto.CheckUserAndKeyboardInteractiveRequest
	4, // 4: proto.Auth.SendKeyboardAuthRequest:input_type -> proto.KeyboardAuthRequest
	6, // 5: proto.Auth.CheckUserAndPass:output_type -> proto.AuthResponse
	6, // 6: proto.Auth.CheckUserAndTLSCert:output_type -> proto.AuthResponse
	6, // 7: proto.Auth.CheckUserAndPublicKey:output_type -> proto.AuthResponse
	6, // 8: proto.Auth.CheckUserAndKeyboardInteractive:output_type -> proto.AuthResponse
	5, // 9: proto.Auth.SendKeyboardAuthRequest:output_type -> proto.KeyboardAuthResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_auth_proto_init() }
func file_auth_proto_auth_proto_init() {
	if File_auth_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAndPassRequest); i {
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
		file_auth_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAndTLSCertRequest); i {
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
		file_auth_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAndPublicKeyRequest); i {
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
		file_auth_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAndKeyboardInteractiveRequest); i {
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
		file_auth_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardAuthRequest); i {
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
		file_auth_proto_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardAuthResponse); i {
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
		file_auth_proto_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
			RawDescriptor: file_auth_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_auth_proto_msgTypes,
	}.Build()
	File_auth_proto_auth_proto = out.File
	file_auth_proto_auth_proto_rawDesc = nil
	file_auth_proto_auth_proto_goTypes = nil
	file_auth_proto_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	CheckUserAndPass(ctx context.Context, in *CheckUserAndPassRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	CheckUserAndTLSCert(ctx context.Context, in *CheckUserAndTLSCertRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	CheckUserAndPublicKey(ctx context.Context, in *CheckUserAndPublicKeyRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	CheckUserAndKeyboardInteractive(ctx context.Context, in *CheckUserAndKeyboardInteractiveRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	SendKeyboardAuthRequest(ctx context.Context, in *KeyboardAuthRequest, opts ...grpc.CallOption) (*KeyboardAuthResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) CheckUserAndPass(ctx context.Context, in *CheckUserAndPassRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/CheckUserAndPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckUserAndTLSCert(ctx context.Context, in *CheckUserAndTLSCertRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/CheckUserAndTLSCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckUserAndPublicKey(ctx context.Context, in *CheckUserAndPublicKeyRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/CheckUserAndPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckUserAndKeyboardInteractive(ctx context.Context, in *CheckUserAndKeyboardInteractiveRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/CheckUserAndKeyboardInteractive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendKeyboardAuthRequest(ctx context.Context, in *KeyboardAuthRequest, opts ...grpc.CallOption) (*KeyboardAuthResponse, error) {
	out := new(KeyboardAuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/SendKeyboardAuthRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	CheckUserAndPass(context.Context, *CheckUserAndPassRequest) (*AuthResponse, error)
	CheckUserAndTLSCert(context.Context, *CheckUserAndTLSCertRequest) (*AuthResponse, error)
	CheckUserAndPublicKey(context.Context, *CheckUserAndPublicKeyRequest) (*AuthResponse, error)
	CheckUserAndKeyboardInteractive(context.Context, *CheckUserAndKeyboardInteractiveRequest) (*AuthResponse, error)
	SendKeyboardAuthRequest(context.Context, *KeyboardAuthRequest) (*KeyboardAuthResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) CheckUserAndPass(context.Context, *CheckUserAndPassRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAndPass not implemented")
}
func (*UnimplementedAuthServer) CheckUserAndTLSCert(context.Context, *CheckUserAndTLSCertRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAndTLSCert not implemented")
}
func (*UnimplementedAuthServer) CheckUserAndPublicKey(context.Context, *CheckUserAndPublicKeyRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAndPublicKey not implemented")
}
func (*UnimplementedAuthServer) CheckUserAndKeyboardInteractive(context.Context, *CheckUserAndKeyboardInteractiveRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAndKeyboardInteractive not implemented")
}
func (*UnimplementedAuthServer) SendKeyboardAuthRequest(context.Context, *KeyboardAuthRequest) (*KeyboardAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKeyboardAuthRequest not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_CheckUserAndPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAndPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckUserAndPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/CheckUserAndPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckUserAndPass(ctx, req.(*CheckUserAndPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckUserAndTLSCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAndTLSCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckUserAndTLSCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/CheckUserAndTLSCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckUserAndTLSCert(ctx, req.(*CheckUserAndTLSCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckUserAndPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAndPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckUserAndPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/CheckUserAndPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckUserAndPublicKey(ctx, req.(*CheckUserAndPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckUserAndKeyboardInteractive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAndKeyboardInteractiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckUserAndKeyboardInteractive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/CheckUserAndKeyboardInteractive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckUserAndKeyboardInteractive(ctx, req.(*CheckUserAndKeyboardInteractiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendKeyboardAuthRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyboardAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendKeyboardAuthRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/SendKeyboardAuthRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendKeyboardAuthRequest(ctx, req.(*KeyboardAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUserAndPass",
			Handler:    _Auth_CheckUserAndPass_Handler,
		},
		{
			MethodName: "CheckUserAndTLSCert",
			Handler:    _Auth_CheckUserAndTLSCert_Handler,
		},
		{
			MethodName: "CheckUserAndPublicKey",
			Handler:    _Auth_CheckUserAndPublicKey_Handler,
		},
		{
			MethodName: "CheckUserAndKeyboardInteractive",
			Handler:    _Auth_CheckUserAndKeyboardInteractive_Handler,
		},
		{
			MethodName: "SendKeyboardAuthRequest",
			Handler:    _Auth_SendKeyboardAuthRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/proto/auth.proto",
}