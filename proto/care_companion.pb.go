// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: proto/care_companion.proto

package care_companion

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_care_companion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_care_companion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_care_companion_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_care_companion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_care_companion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_proto_care_companion_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCareSeekerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCareSeekerRequest) Reset() {
	*x = GetCareSeekerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_care_companion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCareSeekerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCareSeekerRequest) ProtoMessage() {}

func (x *GetCareSeekerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_care_companion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCareSeekerRequest.ProtoReflect.Descriptor instead.
func (*GetCareSeekerRequest) Descriptor() ([]byte, []int) {
	return file_proto_care_companion_proto_rawDescGZIP(), []int{2}
}

func (x *GetCareSeekerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCareSeekerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CareSeeker *CareSeeker `protobuf:"bytes,1,opt,name=care_seeker,json=careSeeker,proto3" json:"care_seeker,omitempty"`
}

func (x *GetCareSeekerResponse) Reset() {
	*x = GetCareSeekerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_care_companion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCareSeekerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCareSeekerResponse) ProtoMessage() {}

func (x *GetCareSeekerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_care_companion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCareSeekerResponse.ProtoReflect.Descriptor instead.
func (*GetCareSeekerResponse) Descriptor() ([]byte, []int) {
	return file_proto_care_companion_proto_rawDescGZIP(), []int{3}
}

func (x *GetCareSeekerResponse) GetCareSeeker() *CareSeeker {
	if x != nil {
		return x.CareSeeker
	}
	return nil
}

type CreateCareSeekerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCareSeekerRequest) Reset() {
	*x = CreateCareSeekerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_care_companion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCareSeekerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCareSeekerRequest) ProtoMessage() {}

func (x *CreateCareSeekerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_care_companion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCareSeekerRequest.ProtoReflect.Descriptor instead.
func (*CreateCareSeekerRequest) Descriptor() ([]byte, []int) {
	return file_proto_care_companion_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCareSeekerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCareSeekerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CareSeeker *CareSeeker `protobuf:"bytes,1,opt,name=care_seeker,json=careSeeker,proto3" json:"care_seeker,omitempty"`
}

func (x *CreateCareSeekerResponse) Reset() {
	*x = CreateCareSeekerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_care_companion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCareSeekerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCareSeekerResponse) ProtoMessage() {}

func (x *CreateCareSeekerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_care_companion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCareSeekerResponse.ProtoReflect.Descriptor instead.
func (*CreateCareSeekerResponse) Descriptor() ([]byte, []int) {
	return file_proto_care_companion_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCareSeekerResponse) GetCareSeeker() *CareSeeker {
	if x != nil {
		return x.CareSeeker
	}
	return nil
}

type CareSeeker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CareSeeker) Reset() {
	*x = CareSeeker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_care_companion_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CareSeeker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CareSeeker) ProtoMessage() {}

func (x *CareSeeker) ProtoReflect() protoreflect.Message {
	mi := &file_proto_care_companion_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CareSeeker.ProtoReflect.Descriptor instead.
func (*CareSeeker) Descriptor() ([]byte, []int) {
	return file_proto_care_companion_proto_rawDescGZIP(), []int{6}
}

func (x *CareSeeker) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CareSeeker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_care_companion_proto protoreflect.FileDescriptor

var file_proto_care_companion_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x61,
	0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65,
	0x72, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x22, 0x2d, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x65,
	0x5f, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x65, 0x53,
	0x65, 0x65, 0x6b, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65,
	0x6b, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x91, 0x02, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x63, 0x61, 0x72,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x63, 0x61,
	0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65,
	0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e,
	0x63, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x6d, 0x61, 0x79, 0x6f,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x6d, 0x6f, 0x68, 0x61, 0x6d, 0x6d, 0x61, 0x64, 0x69, 0x2f, 0x63,
	0x61, 0x72, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_care_companion_proto_rawDescOnce sync.Once
	file_proto_care_companion_proto_rawDescData = file_proto_care_companion_proto_rawDesc
)

func file_proto_care_companion_proto_rawDescGZIP() []byte {
	file_proto_care_companion_proto_rawDescOnce.Do(func() {
		file_proto_care_companion_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_care_companion_proto_rawDescData)
	})
	return file_proto_care_companion_proto_rawDescData
}

var file_proto_care_companion_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_care_companion_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: care_companion.Empty
	(*PingResponse)(nil),             // 1: care_companion.PingResponse
	(*GetCareSeekerRequest)(nil),     // 2: care_companion.GetCareSeekerRequest
	(*GetCareSeekerResponse)(nil),    // 3: care_companion.GetCareSeekerResponse
	(*CreateCareSeekerRequest)(nil),  // 4: care_companion.CreateCareSeekerRequest
	(*CreateCareSeekerResponse)(nil), // 5: care_companion.CreateCareSeekerResponse
	(*CareSeeker)(nil),               // 6: care_companion.CareSeeker
}
var file_proto_care_companion_proto_depIdxs = []int32{
	6, // 0: care_companion.GetCareSeekerResponse.care_seeker:type_name -> care_companion.CareSeeker
	6, // 1: care_companion.CreateCareSeekerResponse.care_seeker:type_name -> care_companion.CareSeeker
	2, // 2: care_companion.CareCompanion.GetCareSeeker:input_type -> care_companion.GetCareSeekerRequest
	4, // 3: care_companion.CareCompanion.CreateCareSeeker:input_type -> care_companion.CreateCareSeekerRequest
	0, // 4: care_companion.CareCompanion.Ping:input_type -> care_companion.Empty
	3, // 5: care_companion.CareCompanion.GetCareSeeker:output_type -> care_companion.GetCareSeekerResponse
	5, // 6: care_companion.CareCompanion.CreateCareSeeker:output_type -> care_companion.CreateCareSeekerResponse
	1, // 7: care_companion.CareCompanion.Ping:output_type -> care_companion.PingResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_care_companion_proto_init() }
func file_proto_care_companion_proto_init() {
	if File_proto_care_companion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_care_companion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_care_companion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_proto_care_companion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCareSeekerRequest); i {
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
		file_proto_care_companion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCareSeekerResponse); i {
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
		file_proto_care_companion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCareSeekerRequest); i {
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
		file_proto_care_companion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCareSeekerResponse); i {
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
		file_proto_care_companion_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CareSeeker); i {
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
			RawDescriptor: file_proto_care_companion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_care_companion_proto_goTypes,
		DependencyIndexes: file_proto_care_companion_proto_depIdxs,
		MessageInfos:      file_proto_care_companion_proto_msgTypes,
	}.Build()
	File_proto_care_companion_proto = out.File
	file_proto_care_companion_proto_rawDesc = nil
	file_proto_care_companion_proto_goTypes = nil
	file_proto_care_companion_proto_depIdxs = nil
}