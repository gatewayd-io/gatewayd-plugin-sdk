// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: plugin/v1/plugin.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HookName int32

const (
	HookName_HOOK_NAME_UNSPECIFIED            HookName = 0
	HookName_HOOK_NAME_ON_CONFIG_LOADED       HookName = 1
	HookName_HOOK_NAME_ON_NEW_LOGGER          HookName = 2
	HookName_HOOK_NAME_ON_NEW_POOL            HookName = 3
	HookName_HOOK_NAME_ON_NEW_CLIENT          HookName = 4
	HookName_HOOK_NAME_ON_NEW_PROXY           HookName = 5
	HookName_HOOK_NAME_ON_NEW_SERVER          HookName = 6
	HookName_HOOK_NAME_ON_SIGNAL              HookName = 7
	HookName_HOOK_NAME_ON_RUN                 HookName = 8
	HookName_HOOK_NAME_ON_BOOTING             HookName = 9
	HookName_HOOK_NAME_ON_BOOTED              HookName = 10
	HookName_HOOK_NAME_ON_OPENING             HookName = 11
	HookName_HOOK_NAME_ON_OPENED              HookName = 12
	HookName_HOOK_NAME_ON_CLOSING             HookName = 13
	HookName_HOOK_NAME_ON_CLOSED              HookName = 14
	HookName_HOOK_NAME_ON_TRAFFIC             HookName = 15
	HookName_HOOK_NAME_ON_TRAFFIC_FROM_CLIENT HookName = 16
	HookName_HOOK_NAME_ON_TRAFFIC_TO_SERVER   HookName = 17
	HookName_HOOK_NAME_ON_TRAFFIC_FROM_SERVER HookName = 18
	HookName_HOOK_NAME_ON_TRAFFIC_TO_CLIENT   HookName = 19
	HookName_HOOK_NAME_ON_SHUTDOWN            HookName = 20
	HookName_HOOK_NAME_ON_TICK                HookName = 21
	HookName_HOOK_NAME_ON_HOOK                HookName = 22
)

// Enum value maps for HookName.
var (
	HookName_name = map[int32]string{
		0:  "HOOK_NAME_UNSPECIFIED",
		1:  "HOOK_NAME_ON_CONFIG_LOADED",
		2:  "HOOK_NAME_ON_NEW_LOGGER",
		3:  "HOOK_NAME_ON_NEW_POOL",
		4:  "HOOK_NAME_ON_NEW_CLIENT",
		5:  "HOOK_NAME_ON_NEW_PROXY",
		6:  "HOOK_NAME_ON_NEW_SERVER",
		7:  "HOOK_NAME_ON_SIGNAL",
		8:  "HOOK_NAME_ON_RUN",
		9:  "HOOK_NAME_ON_BOOTING",
		10: "HOOK_NAME_ON_BOOTED",
		11: "HOOK_NAME_ON_OPENING",
		12: "HOOK_NAME_ON_OPENED",
		13: "HOOK_NAME_ON_CLOSING",
		14: "HOOK_NAME_ON_CLOSED",
		15: "HOOK_NAME_ON_TRAFFIC",
		16: "HOOK_NAME_ON_TRAFFIC_FROM_CLIENT",
		17: "HOOK_NAME_ON_TRAFFIC_TO_SERVER",
		18: "HOOK_NAME_ON_TRAFFIC_FROM_SERVER",
		19: "HOOK_NAME_ON_TRAFFIC_TO_CLIENT",
		20: "HOOK_NAME_ON_SHUTDOWN",
		21: "HOOK_NAME_ON_TICK",
		22: "HOOK_NAME_ON_HOOK",
	}
	HookName_value = map[string]int32{
		"HOOK_NAME_UNSPECIFIED":            0,
		"HOOK_NAME_ON_CONFIG_LOADED":       1,
		"HOOK_NAME_ON_NEW_LOGGER":          2,
		"HOOK_NAME_ON_NEW_POOL":            3,
		"HOOK_NAME_ON_NEW_CLIENT":          4,
		"HOOK_NAME_ON_NEW_PROXY":           5,
		"HOOK_NAME_ON_NEW_SERVER":          6,
		"HOOK_NAME_ON_SIGNAL":              7,
		"HOOK_NAME_ON_RUN":                 8,
		"HOOK_NAME_ON_BOOTING":             9,
		"HOOK_NAME_ON_BOOTED":              10,
		"HOOK_NAME_ON_OPENING":             11,
		"HOOK_NAME_ON_OPENED":              12,
		"HOOK_NAME_ON_CLOSING":             13,
		"HOOK_NAME_ON_CLOSED":              14,
		"HOOK_NAME_ON_TRAFFIC":             15,
		"HOOK_NAME_ON_TRAFFIC_FROM_CLIENT": 16,
		"HOOK_NAME_ON_TRAFFIC_TO_SERVER":   17,
		"HOOK_NAME_ON_TRAFFIC_FROM_SERVER": 18,
		"HOOK_NAME_ON_TRAFFIC_TO_CLIENT":   19,
		"HOOK_NAME_ON_SHUTDOWN":            20,
		"HOOK_NAME_ON_TICK":                21,
		"HOOK_NAME_ON_HOOK":                22,
	}
)

func (x HookName) Enum() *HookName {
	p := new(HookName)
	*p = x
	return p
}

func (x HookName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HookName) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_v1_plugin_proto_enumTypes[0].Descriptor()
}

func (HookName) Type() protoreflect.EnumType {
	return &file_plugin_v1_plugin_proto_enumTypes[0]
}

func (x HookName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HookName.Descriptor instead.
func (HookName) EnumDescriptor() ([]byte, []int) {
	return file_plugin_v1_plugin_proto_rawDescGZIP(), []int{0}
}

type PluginID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version       string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	RemoteUrl     string                 `protobuf:"bytes,3,opt,name=remote_url,json=remoteUrl,proto3" json:"remote_url,omitempty"`
	Checksum      string                 `protobuf:"bytes,4,opt,name=checksum,proto3" json:"checksum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginID) Reset() {
	*x = PluginID{}
	mi := &file_plugin_v1_plugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginID) ProtoMessage() {}

func (x *PluginID) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_v1_plugin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginID.ProtoReflect.Descriptor instead.
func (*PluginID) Descriptor() ([]byte, []int) {
	return file_plugin_v1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *PluginID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginID) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginID) GetRemoteUrl() string {
	if x != nil {
		return x.RemoteUrl
	}
	return ""
}

func (x *PluginID) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

type PluginConfig struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	Id          *PluginID              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Authors     []string               `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	License     string                 `protobuf:"bytes,4,opt,name=license,proto3" json:"license,omitempty"`
	ProjectUrl  string                 `protobuf:"bytes,5,opt,name=project_url,json=projectUrl,proto3" json:"project_url,omitempty"`
	// internal and external config options
	Config map[string]string `protobuf:"bytes,6,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// hooks it attaches to
	Hooks []HookName `protobuf:"varint,7,rep,packed,name=hooks,proto3,enum=plugin.v1.HookName" json:"hooks,omitempty"`
	// required plugins
	Requires      map[string]string `protobuf:"bytes,8,rep,name=requires,proto3" json:"requires,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tags          []string          `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	Categories    []string          `protobuf:"bytes,10,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginConfig) Reset() {
	*x = PluginConfig{}
	mi := &file_plugin_v1_plugin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfig) ProtoMessage() {}

func (x *PluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_v1_plugin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfig.ProtoReflect.Descriptor instead.
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return file_plugin_v1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *PluginConfig) GetId() *PluginID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *PluginConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PluginConfig) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *PluginConfig) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *PluginConfig) GetProjectUrl() string {
	if x != nil {
		return x.ProjectUrl
	}
	return ""
}

func (x *PluginConfig) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *PluginConfig) GetHooks() []HookName {
	if x != nil {
		return x.Hooks
	}
	return nil
}

func (x *PluginConfig) GetRequires() map[string]string {
	if x != nil {
		return x.Requires
	}
	return nil
}

func (x *PluginConfig) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PluginConfig) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_plugin_v1_plugin_proto protoreflect.FileDescriptor

var file_plugin_v1_plugin_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x16, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x08, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x22, 0x81, 0x04, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x3b, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x05, 0x68, 0x6f,
	0x6f, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05,
	0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x2a, 0x92, 0x05, 0x0a, 0x08, 0x48, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a,
	0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x47, 0x5f, 0x4c, 0x4f, 0x41, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x57,
	0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x52, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x4f, 0x4f,
	0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x50, 0x4f,
	0x4f, 0x4c, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x04, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f,
	0x4e, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10, 0x05, 0x12, 0x1b, 0x0a,
	0x17, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x4e, 0x45,
	0x57, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x4f,
	0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41,
	0x4c, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45,
	0x5f, 0x4f, 0x4e, 0x5f, 0x52, 0x55, 0x4e, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x4f, 0x4f,
	0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x42, 0x4f, 0x4f, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45,
	0x5f, 0x4f, 0x4e, 0x5f, 0x42, 0x4f, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14,
	0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x45, 0x44, 0x10, 0x0c, 0x12,
	0x18, 0x0a, 0x14, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f,
	0x43, 0x4c, 0x4f, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x0d, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x4f, 0x4f,
	0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44,
	0x10, 0x0e, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x10, 0x0f, 0x12, 0x24, 0x0a, 0x20,
	0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41,
	0x46, 0x46, 0x49, 0x43, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54,
	0x10, 0x10, 0x12, 0x22, 0x0a, 0x1e, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x10, 0x11, 0x12, 0x24, 0x0a, 0x20, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x46,
	0x52, 0x4f, 0x4d, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x12, 0x12, 0x22, 0x0a, 0x1e,
	0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41,
	0x46, 0x46, 0x49, 0x43, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x13,
	0x12, 0x19, 0x0a, 0x15, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e,
	0x5f, 0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x14, 0x12, 0x15, 0x0a, 0x11, 0x48,
	0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x43, 0x4b,
	0x10, 0x15, 0x12, 0x15, 0x0a, 0x11, 0x48, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x4f, 0x4b, 0x10, 0x16, 0x32, 0xd5, 0x09, 0x0a, 0x15, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x36, 0x0a, 0x0e,
	0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x11,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x4f, 0x6e, 0x4e, 0x65, 0x77, 0x4c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x4f, 0x6e, 0x4e,
	0x65, 0x77, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x0b,
	0x4f, 0x6e, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x12, 0x32, 0x0a, 0x0a, 0x4f, 0x6e, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x4f, 0x6e, 0x4e, 0x65, 0x77, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x4f, 0x6e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x05,
	0x4f, 0x6e, 0x52, 0x75, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x4f,
	0x6e, 0x42, 0x6f, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x30,
	0x0a, 0x08, 0x4f, 0x6e, 0x42, 0x6f, 0x6f, 0x74, 0x65, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x12, 0x31, 0x0a, 0x09, 0x4f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x4f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12,
	0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x4f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x69,
	0x6e, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x4f, 0x6e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x4f, 0x6e,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x3b, 0x0a,
	0x13, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x39, 0x0a, 0x11, 0x4f, 0x6e,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x13, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a,
	0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x12, 0x39, 0x0a, 0x11, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54,
	0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x32, 0x0a,
	0x0a, 0x4f, 0x6e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x12, 0x2e, 0x0a, 0x06, 0x4f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x11, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x12, 0x2e, 0x0a, 0x06, 0x4f, 0x6e, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x11, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x11,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x64, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_plugin_v1_plugin_proto_rawDescOnce sync.Once
	file_plugin_v1_plugin_proto_rawDescData []byte
)

func file_plugin_v1_plugin_proto_rawDescGZIP() []byte {
	file_plugin_v1_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_v1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_plugin_v1_plugin_proto_rawDesc), len(file_plugin_v1_plugin_proto_rawDesc)))
	})
	return file_plugin_v1_plugin_proto_rawDescData
}

var file_plugin_v1_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_plugin_v1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_plugin_v1_plugin_proto_goTypes = []any{
	(HookName)(0),        // 0: plugin.v1.HookName
	(*PluginID)(nil),     // 1: plugin.v1.PluginID
	(*PluginConfig)(nil), // 2: plugin.v1.PluginConfig
	nil,                  // 3: plugin.v1.PluginConfig.ConfigEntry
	nil,                  // 4: plugin.v1.PluginConfig.RequiresEntry
	(*Struct)(nil),       // 5: plugin.v1.Struct
}
var file_plugin_v1_plugin_proto_depIdxs = []int32{
	1,  // 0: plugin.v1.PluginConfig.id:type_name -> plugin.v1.PluginID
	3,  // 1: plugin.v1.PluginConfig.config:type_name -> plugin.v1.PluginConfig.ConfigEntry
	0,  // 2: plugin.v1.PluginConfig.hooks:type_name -> plugin.v1.HookName
	4,  // 3: plugin.v1.PluginConfig.requires:type_name -> plugin.v1.PluginConfig.RequiresEntry
	5,  // 4: plugin.v1.GatewayDPluginService.GetPluginConfig:input_type -> plugin.v1.Struct
	5,  // 5: plugin.v1.GatewayDPluginService.OnConfigLoaded:input_type -> plugin.v1.Struct
	5,  // 6: plugin.v1.GatewayDPluginService.OnNewLogger:input_type -> plugin.v1.Struct
	5,  // 7: plugin.v1.GatewayDPluginService.OnNewPool:input_type -> plugin.v1.Struct
	5,  // 8: plugin.v1.GatewayDPluginService.OnNewClient:input_type -> plugin.v1.Struct
	5,  // 9: plugin.v1.GatewayDPluginService.OnNewProxy:input_type -> plugin.v1.Struct
	5,  // 10: plugin.v1.GatewayDPluginService.OnNewServer:input_type -> plugin.v1.Struct
	5,  // 11: plugin.v1.GatewayDPluginService.OnSignal:input_type -> plugin.v1.Struct
	5,  // 12: plugin.v1.GatewayDPluginService.OnRun:input_type -> plugin.v1.Struct
	5,  // 13: plugin.v1.GatewayDPluginService.OnBooting:input_type -> plugin.v1.Struct
	5,  // 14: plugin.v1.GatewayDPluginService.OnBooted:input_type -> plugin.v1.Struct
	5,  // 15: plugin.v1.GatewayDPluginService.OnOpening:input_type -> plugin.v1.Struct
	5,  // 16: plugin.v1.GatewayDPluginService.OnOpened:input_type -> plugin.v1.Struct
	5,  // 17: plugin.v1.GatewayDPluginService.OnClosing:input_type -> plugin.v1.Struct
	5,  // 18: plugin.v1.GatewayDPluginService.OnClosed:input_type -> plugin.v1.Struct
	5,  // 19: plugin.v1.GatewayDPluginService.OnTraffic:input_type -> plugin.v1.Struct
	5,  // 20: plugin.v1.GatewayDPluginService.OnTrafficFromClient:input_type -> plugin.v1.Struct
	5,  // 21: plugin.v1.GatewayDPluginService.OnTrafficToServer:input_type -> plugin.v1.Struct
	5,  // 22: plugin.v1.GatewayDPluginService.OnTrafficFromServer:input_type -> plugin.v1.Struct
	5,  // 23: plugin.v1.GatewayDPluginService.OnTrafficToClient:input_type -> plugin.v1.Struct
	5,  // 24: plugin.v1.GatewayDPluginService.OnShutdown:input_type -> plugin.v1.Struct
	5,  // 25: plugin.v1.GatewayDPluginService.OnTick:input_type -> plugin.v1.Struct
	5,  // 26: plugin.v1.GatewayDPluginService.OnHook:input_type -> plugin.v1.Struct
	5,  // 27: plugin.v1.GatewayDPluginService.GetPluginConfig:output_type -> plugin.v1.Struct
	5,  // 28: plugin.v1.GatewayDPluginService.OnConfigLoaded:output_type -> plugin.v1.Struct
	5,  // 29: plugin.v1.GatewayDPluginService.OnNewLogger:output_type -> plugin.v1.Struct
	5,  // 30: plugin.v1.GatewayDPluginService.OnNewPool:output_type -> plugin.v1.Struct
	5,  // 31: plugin.v1.GatewayDPluginService.OnNewClient:output_type -> plugin.v1.Struct
	5,  // 32: plugin.v1.GatewayDPluginService.OnNewProxy:output_type -> plugin.v1.Struct
	5,  // 33: plugin.v1.GatewayDPluginService.OnNewServer:output_type -> plugin.v1.Struct
	5,  // 34: plugin.v1.GatewayDPluginService.OnSignal:output_type -> plugin.v1.Struct
	5,  // 35: plugin.v1.GatewayDPluginService.OnRun:output_type -> plugin.v1.Struct
	5,  // 36: plugin.v1.GatewayDPluginService.OnBooting:output_type -> plugin.v1.Struct
	5,  // 37: plugin.v1.GatewayDPluginService.OnBooted:output_type -> plugin.v1.Struct
	5,  // 38: plugin.v1.GatewayDPluginService.OnOpening:output_type -> plugin.v1.Struct
	5,  // 39: plugin.v1.GatewayDPluginService.OnOpened:output_type -> plugin.v1.Struct
	5,  // 40: plugin.v1.GatewayDPluginService.OnClosing:output_type -> plugin.v1.Struct
	5,  // 41: plugin.v1.GatewayDPluginService.OnClosed:output_type -> plugin.v1.Struct
	5,  // 42: plugin.v1.GatewayDPluginService.OnTraffic:output_type -> plugin.v1.Struct
	5,  // 43: plugin.v1.GatewayDPluginService.OnTrafficFromClient:output_type -> plugin.v1.Struct
	5,  // 44: plugin.v1.GatewayDPluginService.OnTrafficToServer:output_type -> plugin.v1.Struct
	5,  // 45: plugin.v1.GatewayDPluginService.OnTrafficFromServer:output_type -> plugin.v1.Struct
	5,  // 46: plugin.v1.GatewayDPluginService.OnTrafficToClient:output_type -> plugin.v1.Struct
	5,  // 47: plugin.v1.GatewayDPluginService.OnShutdown:output_type -> plugin.v1.Struct
	5,  // 48: plugin.v1.GatewayDPluginService.OnTick:output_type -> plugin.v1.Struct
	5,  // 49: plugin.v1.GatewayDPluginService.OnHook:output_type -> plugin.v1.Struct
	27, // [27:50] is the sub-list for method output_type
	4,  // [4:27] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_plugin_v1_plugin_proto_init() }
func file_plugin_v1_plugin_proto_init() {
	if File_plugin_v1_plugin_proto != nil {
		return
	}
	file_plugin_v1_struct_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_plugin_v1_plugin_proto_rawDesc), len(file_plugin_v1_plugin_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_v1_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_v1_plugin_proto_depIdxs,
		EnumInfos:         file_plugin_v1_plugin_proto_enumTypes,
		MessageInfos:      file_plugin_v1_plugin_proto_msgTypes,
	}.Build()
	File_plugin_v1_plugin_proto = out.File
	file_plugin_v1_plugin_proto_goTypes = nil
	file_plugin_v1_plugin_proto_depIdxs = nil
}
