// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sonrhq/service/module/v1/state.proto

package modulev1

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	_ "cosmossdk.io/api/cosmos/orm/v1alpha1"
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

// ServicePermissions is the specified permissions the application has with the user's account.
type ServicePermissions int32

const (
	// SERVICE_PERMISSIONS_BASE - off chain visibility
	ServicePermissions_SERVICE_PERMISSIONS_BASE ServicePermissions = 0
	// SERVICE_PERMISSIONS_READ - read access
	ServicePermissions_SERVICE_PERMISSIONS_READ ServicePermissions = 1
	// SERVICE_PERMISSIONS_WRITE - write access
	ServicePermissions_SERVICE_PERMISSIONS_WRITE ServicePermissions = 2
	// SERVICE_PERMISSIONS_OWN - ownership
	ServicePermissions_SERVICE_PERMISSIONS_OWN ServicePermissions = 3
)

// Enum value maps for ServicePermissions.
var (
	ServicePermissions_name = map[int32]string{
		0: "SERVICE_PERMISSIONS_BASE",
		1: "SERVICE_PERMISSIONS_READ",
		2: "SERVICE_PERMISSIONS_WRITE",
		3: "SERVICE_PERMISSIONS_OWN",
	}
	ServicePermissions_value = map[string]int32{
		"SERVICE_PERMISSIONS_BASE":  0,
		"SERVICE_PERMISSIONS_READ":  1,
		"SERVICE_PERMISSIONS_WRITE": 2,
		"SERVICE_PERMISSIONS_OWN":   3,
	}
)

func (x ServicePermissions) Enum() *ServicePermissions {
	p := new(ServicePermissions)
	*p = x
	return p
}

func (x ServicePermissions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServicePermissions) Descriptor() protoreflect.EnumDescriptor {
	return file_sonrhq_service_module_v1_state_proto_enumTypes[0].Descriptor()
}

func (ServicePermissions) Type() protoreflect.EnumType {
	return &file_sonrhq_service_module_v1_state_proto_enumTypes[0]
}

func (x ServicePermissions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServicePermissions.Descriptor instead.
func (ServicePermissions) EnumDescriptor() ([]byte, []int) {
	return file_sonrhq_service_module_v1_state_proto_rawDescGZIP(), []int{0}
}

// Module is the app config object of the module.
// Learn more: https://docs.cosmos.network/main/building-modules/depinject
type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_sonrhq_service_module_v1_state_proto_rawDescGZIP(), []int{0}
}

// ServiceRecord is the balance of an account.
type ServiceRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Did         string             `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Origin      string             `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Name        string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string             `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Owner       string             `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Permissions ServicePermissions `protobuf:"varint,6,opt,name=permissions,proto3,enum=sonrhq.service.module.v1.ServicePermissions" json:"permissions,omitempty"`
}

func (x *ServiceRecord) Reset() {
	*x = ServiceRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRecord) ProtoMessage() {}

func (x *ServiceRecord) ProtoReflect() protoreflect.Message {
	mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRecord.ProtoReflect.Descriptor instead.
func (*ServiceRecord) Descriptor() ([]byte, []int) {
	return file_sonrhq_service_module_v1_state_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceRecord) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *ServiceRecord) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *ServiceRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceRecord) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceRecord) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ServiceRecord) GetPermissions() ServicePermissions {
	if x != nil {
		return x.Permissions
	}
	return ServicePermissions_SERVICE_PERMISSIONS_BASE
}

// UserProfile is the total supply of the module.
type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Did       string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Origin    string `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Handle    string `protobuf:"bytes,3,opt,name=handle,proto3" json:"handle,omitempty"`
	PublicKey []byte `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_sonrhq_service_module_v1_state_proto_rawDescGZIP(), []int{2}
}

func (x *UserProfile) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *UserProfile) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *UserProfile) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

func (x *UserProfile) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Resource is the total supply of the module.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID of the resource (i.e: did:sonr:abc1234/Qm1234)
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// CID of the resource (i.e: Qm1234)
	Cid string `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	// CID type of the resource (i.e: ipfs)
	CidType string `protobuf:"bytes,3,opt,name=cid_type,json=cidType,proto3" json:"cid_type,omitempty"`
	// Origin is the domain of the resource (i.e: "sonr.io")
	Origin string `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	// Owner is the UserProfile DID of the owner of the resource (i.e: did:sonr:abc1234)
	Owner string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_sonrhq_service_module_v1_state_proto_rawDescGZIP(), []int{3}
}

func (x *Resource) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *Resource) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *Resource) GetCidType() string {
	if x != nil {
		return x.CidType
	}
	return ""
}

func (x *Resource) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Resource) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// Property is the total supply of the module.
type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Did of the property (i.e: did:sonr:abc1234#name)
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Key is the fragment of the property (i.e: name)
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Value is the encrypted text of the property (i.e: "John Doe")
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// Origin is the domain of the property (i.e: "sonr.io")
	Origin string `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	// Schema is the schema of the property (i.e: "https://schema.org/name")
	Schema string `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
	// Owner is the UserProfile DID of the owner of the property (i.e: did:sonr:abc1234)
	Owner string `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_sonrhq_service_module_v1_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_sonrhq_service_module_v1_state_proto_rawDescGZIP(), []int{4}
}

func (x *Property) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *Property) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Property) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Property) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Property) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *Property) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

var File_sonrhq_service_module_v1_state_proto protoreflect.FileDescriptor

var file_sonrhq_service_module_v1_state_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x3a, 0x2a, 0x82, 0x9f, 0xd3, 0x8e, 0x03, 0x24, 0x0a, 0x22, 0x08, 0x01,
	0x12, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x64, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x55, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x4f,
	0x0a, 0x0c, 0x0a, 0x0a, 0x64, 0x69, 0x64, 0x2c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x0c,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x10, 0x01, 0x18, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x2c,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x10, 0x03, 0x18, 0x01, 0x12, 0x12, 0x0a, 0x0c, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x2c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x10, 0x04, 0x18, 0x01, 0x18, 0x01, 0x22,
	0xae, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x3a, 0x3e, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x38, 0x0a, 0x13, 0x0a, 0x11, 0x64, 0x69, 0x64, 0x2c,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0d, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x2c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x10, 0x02, 0x18, 0x01, 0x18, 0x02,
	0x22, 0xe9, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x69, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x3a, 0x70, 0xf2, 0x9e, 0xd3, 0x8e,
	0x03, 0x6a, 0x0a, 0x09, 0x0a, 0x07, 0x64, 0x69, 0x64, 0x2c, 0x63, 0x69, 0x64, 0x12, 0x09, 0x0a,
	0x03, 0x63, 0x69, 0x64, 0x10, 0x01, 0x18, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2c, 0x63,
	0x69, 0x64, 0x10, 0x03, 0x18, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x10,
	0x04, 0x12, 0x0f, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2c, 0x63, 0x69, 0x64, 0x10, 0x05,
	0x18, 0x01, 0x12, 0x16, 0x0a, 0x10, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2c, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x2c, 0x63, 0x69, 0x64, 0x10, 0x06, 0x18, 0x01, 0x18, 0x03, 0x22, 0x9b, 0x02, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x3a, 0x8e, 0x01, 0xf2, 0x9e, 0xd3, 0x8e,
	0x03, 0x87, 0x01, 0x0a, 0x05, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x10, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x2c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2c, 0x6b, 0x65, 0x79, 0x10, 0x01,
	0x18, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2c, 0x6b, 0x65, 0x79, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2c, 0x6b, 0x65, 0x79, 0x10,
	0x03, 0x12, 0x10, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2c, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x10, 0x05, 0x12, 0x10,
	0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2c, 0x6b, 0x65, 0x79, 0x10, 0x06, 0x18, 0x01,
	0x12, 0x16, 0x0a, 0x10, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2c, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x2c, 0x6b, 0x65, 0x79, 0x10, 0x07, 0x18, 0x01, 0x18, 0x04, 0x2a, 0x8c, 0x01, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x45, 0x52,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x10, 0x00, 0x12,
	0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a,
	0x19, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x53, 0x5f, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x42, 0xee, 0x01, 0x0a, 0x1c, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x4d, 0xaa,
	0x02, 0x18, 0x53, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x53, 0x6f, 0x6e,
	0x72, 0x68, 0x71, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x53, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x53,
	0x6f, 0x6e, 0x72, 0x68, 0x71, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sonrhq_service_module_v1_state_proto_rawDescOnce sync.Once
	file_sonrhq_service_module_v1_state_proto_rawDescData = file_sonrhq_service_module_v1_state_proto_rawDesc
)

func file_sonrhq_service_module_v1_state_proto_rawDescGZIP() []byte {
	file_sonrhq_service_module_v1_state_proto_rawDescOnce.Do(func() {
		file_sonrhq_service_module_v1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonrhq_service_module_v1_state_proto_rawDescData)
	})
	return file_sonrhq_service_module_v1_state_proto_rawDescData
}

var file_sonrhq_service_module_v1_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sonrhq_service_module_v1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sonrhq_service_module_v1_state_proto_goTypes = []interface{}{
	(ServicePermissions)(0), // 0: sonrhq.service.module.v1.ServicePermissions
	(*State)(nil),           // 1: sonrhq.service.module.v1.State
	(*ServiceRecord)(nil),   // 2: sonrhq.service.module.v1.ServiceRecord
	(*UserProfile)(nil),     // 3: sonrhq.service.module.v1.UserProfile
	(*Resource)(nil),        // 4: sonrhq.service.module.v1.Resource
	(*Property)(nil),        // 5: sonrhq.service.module.v1.Property
}
var file_sonrhq_service_module_v1_state_proto_depIdxs = []int32{
	0, // 0: sonrhq.service.module.v1.ServiceRecord.permissions:type_name -> sonrhq.service.module.v1.ServicePermissions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sonrhq_service_module_v1_state_proto_init() }
func file_sonrhq_service_module_v1_state_proto_init() {
	if File_sonrhq_service_module_v1_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sonrhq_service_module_v1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_sonrhq_service_module_v1_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRecord); i {
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
		file_sonrhq_service_module_v1_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
		file_sonrhq_service_module_v1_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_sonrhq_service_module_v1_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
			RawDescriptor: file_sonrhq_service_module_v1_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonrhq_service_module_v1_state_proto_goTypes,
		DependencyIndexes: file_sonrhq_service_module_v1_state_proto_depIdxs,
		EnumInfos:         file_sonrhq_service_module_v1_state_proto_enumTypes,
		MessageInfos:      file_sonrhq_service_module_v1_state_proto_msgTypes,
	}.Build()
	File_sonrhq_service_module_v1_state_proto = out.File
	file_sonrhq_service_module_v1_state_proto_rawDesc = nil
	file_sonrhq_service_module_v1_state_proto_goTypes = nil
	file_sonrhq_service_module_v1_state_proto_depIdxs = nil
}
