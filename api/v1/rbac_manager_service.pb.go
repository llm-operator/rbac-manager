// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: api/v1/rbac_manager_service.proto

package v1

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

type AuthorizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token          string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	AccessResource string `protobuf:"bytes,2,opt,name=access_resource,json=accessResource,proto3" json:"access_resource,omitempty"`
	Capability     string `protobuf:"bytes,3,opt,name=capability,proto3" json:"capability,omitempty"`
	OrganizationId string `protobuf:"bytes,4,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	ProjectId      string `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *AuthorizeRequest) Reset() {
	*x = AuthorizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRequest) ProtoMessage() {}

func (x *AuthorizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthorizeRequest) GetAccessResource() string {
	if x != nil {
		return x.AccessResource
	}
	return ""
}

func (x *AuthorizeRequest) GetCapability() string {
	if x != nil {
		return x.Capability
	}
	return ""
}

func (x *AuthorizeRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *AuthorizeRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type AuthorizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized   bool          `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	User         *User         `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Organization *Organization `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
	Project      *Project      `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
	TenantId     string        `protobuf:"bytes,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *AuthorizeResponse) Reset() {
	*x = AuthorizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeResponse) ProtoMessage() {}

func (x *AuthorizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeResponse) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

func (x *AuthorizeResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *AuthorizeResponse) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *AuthorizeResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *AuthorizeResponse) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type AuthorizeWorkerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthorizeWorkerRequest) Reset() {
	*x = AuthorizeWorkerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeWorkerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeWorkerRequest) ProtoMessage() {}

func (x *AuthorizeWorkerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeWorkerRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeWorkerRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizeWorkerRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthorizeWorkerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool     `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	Cluster    *Cluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	TenantId   string   `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *AuthorizeWorkerResponse) Reset() {
	*x = AuthorizeWorkerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeWorkerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeWorkerResponse) ProtoMessage() {}

func (x *AuthorizeWorkerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeWorkerResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeWorkerResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorizeWorkerResponse) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

func (x *AuthorizeWorkerResponse) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *AuthorizeWorkerResponse) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{5}
}

func (x *Organization) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string                           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AssignedKubernetesEnvs []*Project_AssignedKubernetesEnv `protobuf:"bytes,2,rep,name=assigned_kubernetes_envs,json=assignedKubernetesEnvs,proto3" json:"assigned_kubernetes_envs,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{6}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetAssignedKubernetesEnvs() []*Project_AssignedKubernetesEnv {
	if x != nil {
		return x.AssignedKubernetesEnvs
	}
	return nil
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{7}
}

func (x *Cluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Project_AssignedKubernetesEnv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Project_AssignedKubernetesEnv) Reset() {
	*x = Project_AssignedKubernetesEnv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rbac_manager_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project_AssignedKubernetesEnv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project_AssignedKubernetesEnv) ProtoMessage() {}

func (x *Project_AssignedKubernetesEnv) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_manager_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project_AssignedKubernetesEnv.ProtoReflect.Descriptor instead.
func (*Project_AssignedKubernetesEnv) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_manager_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *Project_AssignedKubernetesEnv) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Project_AssignedKubernetesEnv) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_api_v1_rbac_manager_service_proto protoreflect.FileDescriptor

var file_api_v1_rbac_manager_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22,
	0xb9, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x93, 0x02, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62,
	0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62, 0x61, 0x63,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x2e, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x95, 0x01, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x3d, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62, 0x61,
	0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xe4, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x73, 0x0a,
	0x18, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62,
	0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x76, 0x52, 0x16, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x76, 0x73, 0x1a, 0x54, 0x0a, 0x15, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x76, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xfb, 0x01, 0x0a, 0x13, 0x52, 0x62, 0x61, 0x63, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x09, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7a, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x32, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6c,
	0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x6c, 0x6d, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x62, 0x61,
	0x63, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_rbac_manager_service_proto_rawDescOnce sync.Once
	file_api_v1_rbac_manager_service_proto_rawDescData = file_api_v1_rbac_manager_service_proto_rawDesc
)

func file_api_v1_rbac_manager_service_proto_rawDescGZIP() []byte {
	file_api_v1_rbac_manager_service_proto_rawDescOnce.Do(func() {
		file_api_v1_rbac_manager_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_rbac_manager_service_proto_rawDescData)
	})
	return file_api_v1_rbac_manager_service_proto_rawDescData
}

var file_api_v1_rbac_manager_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_rbac_manager_service_proto_goTypes = []interface{}{
	(*AuthorizeRequest)(nil),              // 0: llmoperator.rbac.server.v1.AuthorizeRequest
	(*AuthorizeResponse)(nil),             // 1: llmoperator.rbac.server.v1.AuthorizeResponse
	(*AuthorizeWorkerRequest)(nil),        // 2: llmoperator.rbac.server.v1.AuthorizeWorkerRequest
	(*AuthorizeWorkerResponse)(nil),       // 3: llmoperator.rbac.server.v1.AuthorizeWorkerResponse
	(*User)(nil),                          // 4: llmoperator.rbac.server.v1.User
	(*Organization)(nil),                  // 5: llmoperator.rbac.server.v1.Organization
	(*Project)(nil),                       // 6: llmoperator.rbac.server.v1.Project
	(*Cluster)(nil),                       // 7: llmoperator.rbac.server.v1.Cluster
	(*Project_AssignedKubernetesEnv)(nil), // 8: llmoperator.rbac.server.v1.Project.AssignedKubernetesEnv
}
var file_api_v1_rbac_manager_service_proto_depIdxs = []int32{
	4, // 0: llmoperator.rbac.server.v1.AuthorizeResponse.user:type_name -> llmoperator.rbac.server.v1.User
	5, // 1: llmoperator.rbac.server.v1.AuthorizeResponse.organization:type_name -> llmoperator.rbac.server.v1.Organization
	6, // 2: llmoperator.rbac.server.v1.AuthorizeResponse.project:type_name -> llmoperator.rbac.server.v1.Project
	7, // 3: llmoperator.rbac.server.v1.AuthorizeWorkerResponse.cluster:type_name -> llmoperator.rbac.server.v1.Cluster
	8, // 4: llmoperator.rbac.server.v1.Project.assigned_kubernetes_envs:type_name -> llmoperator.rbac.server.v1.Project.AssignedKubernetesEnv
	0, // 5: llmoperator.rbac.server.v1.RbacInternalService.Authorize:input_type -> llmoperator.rbac.server.v1.AuthorizeRequest
	2, // 6: llmoperator.rbac.server.v1.RbacInternalService.AuthorizeWorker:input_type -> llmoperator.rbac.server.v1.AuthorizeWorkerRequest
	1, // 7: llmoperator.rbac.server.v1.RbacInternalService.Authorize:output_type -> llmoperator.rbac.server.v1.AuthorizeResponse
	3, // 8: llmoperator.rbac.server.v1.RbacInternalService.AuthorizeWorker:output_type -> llmoperator.rbac.server.v1.AuthorizeWorkerResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_rbac_manager_service_proto_init() }
func file_api_v1_rbac_manager_service_proto_init() {
	if File_api_v1_rbac_manager_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_rbac_manager_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeRequest); i {
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
		file_api_v1_rbac_manager_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeResponse); i {
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
		file_api_v1_rbac_manager_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeWorkerRequest); i {
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
		file_api_v1_rbac_manager_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeWorkerResponse); i {
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
		file_api_v1_rbac_manager_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_api_v1_rbac_manager_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
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
		file_api_v1_rbac_manager_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_api_v1_rbac_manager_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_api_v1_rbac_manager_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project_AssignedKubernetesEnv); i {
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
			RawDescriptor: file_api_v1_rbac_manager_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_rbac_manager_service_proto_goTypes,
		DependencyIndexes: file_api_v1_rbac_manager_service_proto_depIdxs,
		MessageInfos:      file_api_v1_rbac_manager_service_proto_msgTypes,
	}.Build()
	File_api_v1_rbac_manager_service_proto = out.File
	file_api_v1_rbac_manager_service_proto_rawDesc = nil
	file_api_v1_rbac_manager_service_proto_goTypes = nil
	file_api_v1_rbac_manager_service_proto_depIdxs = nil
}
