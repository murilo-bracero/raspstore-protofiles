// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: file-manager.proto

package pb

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

type FileMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri       string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	CreatedBy string `protobuf:"bytes,4,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedBy string `protobuf:"bytes,5,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
}

func (x *FileMetadata) Reset() {
	*x = FileMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadata) ProtoMessage() {}

func (x *FileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadata.ProtoReflect.Descriptor instead.
func (*FileMetadata) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{0}
}

func (x *FileMetadata) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *FileMetadata) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *FileMetadata) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *FileMetadata) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *FileMetadata `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Chunk []byte        `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetInfo() *FileMetadata {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *File) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type FileRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uri       string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	Size      uint32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	UpdatedAt int64  `protobuf:"varint,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	CreatedBy string `protobuf:"bytes,5,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedBy string `protobuf:"bytes,6,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
}

func (x *FileRef) Reset() {
	*x = FileRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRef) ProtoMessage() {}

func (x *FileRef) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRef.ProtoReflect.Descriptor instead.
func (*FileRef) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{2}
}

func (x *FileRef) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileRef) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *FileRef) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileRef) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *FileRef) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *FileRef) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{3}
}

func (x *GetFileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFileResponse) Reset() {
	*x = DeleteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileResponse) ProtoMessage() {}

func (x *DeleteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateFileRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename  string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	CreatedBy string `protobuf:"bytes,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *CreateFileRequestData) Reset() {
	*x = CreateFileRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileRequestData) ProtoMessage() {}

func (x *CreateFileRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileRequestData.ProtoReflect.Descriptor instead.
func (*CreateFileRequestData) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{5}
}

func (x *CreateFileRequestData) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *CreateFileRequestData) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type UpdateFileRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdatedBy string `protobuf:"bytes,2,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Filename  string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *UpdateFileRequestData) Reset() {
	*x = UpdateFileRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileRequestData) ProtoMessage() {}

func (x *UpdateFileRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileRequestData.ProtoReflect.Descriptor instead.
func (*UpdateFileRequestData) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFileRequestData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateFileRequestData) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *UpdateFileRequestData) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type CreateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*CreateFileRequest_Filedata
	//	*CreateFileRequest_Chunk
	Data isCreateFileRequest_Data `protobuf_oneof:"data"`
}

func (x *CreateFileRequest) Reset() {
	*x = CreateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileRequest) ProtoMessage() {}

func (x *CreateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileRequest.ProtoReflect.Descriptor instead.
func (*CreateFileRequest) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{7}
}

func (m *CreateFileRequest) GetData() isCreateFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CreateFileRequest) GetFiledata() *CreateFileRequestData {
	if x, ok := x.GetData().(*CreateFileRequest_Filedata); ok {
		return x.Filedata
	}
	return nil
}

func (x *CreateFileRequest) GetChunk() []byte {
	if x, ok := x.GetData().(*CreateFileRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isCreateFileRequest_Data interface {
	isCreateFileRequest_Data()
}

type CreateFileRequest_Filedata struct {
	Filedata *CreateFileRequestData `protobuf:"bytes,1,opt,name=filedata,proto3,oneof"`
}

type CreateFileRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*CreateFileRequest_Filedata) isCreateFileRequest_Data() {}

func (*CreateFileRequest_Chunk) isCreateFileRequest_Data() {}

type UpdateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UpdateFileRequest_Filedata
	//	*UpdateFileRequest_Chunk
	Data isUpdateFileRequest_Data `protobuf_oneof:"data"`
}

func (x *UpdateFileRequest) Reset() {
	*x = UpdateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileRequest) ProtoMessage() {}

func (x *UpdateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{8}
}

func (m *UpdateFileRequest) GetData() isUpdateFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UpdateFileRequest) GetFiledata() *UpdateFileRequestData {
	if x, ok := x.GetData().(*UpdateFileRequest_Filedata); ok {
		return x.Filedata
	}
	return nil
}

func (x *UpdateFileRequest) GetChunk() []byte {
	if x, ok := x.GetData().(*UpdateFileRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isUpdateFileRequest_Data interface {
	isUpdateFileRequest_Data()
}

type UpdateFileRequest_Filedata struct {
	Filedata *UpdateFileRequestData `protobuf:"bytes,1,opt,name=filedata,proto3,oneof"`
}

type UpdateFileRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UpdateFileRequest_Filedata) isUpdateFileRequest_Data() {}

func (*UpdateFileRequest_Chunk) isUpdateFileRequest_Data() {}

type FindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllRequest) Reset() {
	*x = FindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_manager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRequest) ProtoMessage() {}

func (x *FindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_manager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRequest.ProtoReflect.Descriptor instead.
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return file_file_manager_proto_rawDescGZIP(), []int{9}
}

var File_file_manager_proto protoreflect.FileDescriptor

var file_file_manager_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x7a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x22, 0x42, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x99, 0x01, 0x0a, 0x07, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x22, 0x62, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x6c, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x10, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x32, 0x86, 0x02, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x28, 0x01, 0x12, 0x2a, 0x0a, 0x08, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x28, 0x01, 0x12, 0x2e, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x30, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_manager_proto_rawDescOnce sync.Once
	file_file_manager_proto_rawDescData = file_file_manager_proto_rawDesc
)

func file_file_manager_proto_rawDescGZIP() []byte {
	file_file_manager_proto_rawDescOnce.Do(func() {
		file_file_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_manager_proto_rawDescData)
	})
	return file_file_manager_proto_rawDescData
}

var file_file_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_file_manager_proto_goTypes = []interface{}{
	(*FileMetadata)(nil),          // 0: pb.FileMetadata
	(*File)(nil),                  // 1: pb.File
	(*FileRef)(nil),               // 2: pb.FileRef
	(*GetFileRequest)(nil),        // 3: pb.GetFileRequest
	(*DeleteFileResponse)(nil),    // 4: pb.DeleteFileResponse
	(*CreateFileRequestData)(nil), // 5: pb.CreateFileRequestData
	(*UpdateFileRequestData)(nil), // 6: pb.UpdateFileRequestData
	(*CreateFileRequest)(nil),     // 7: pb.CreateFileRequest
	(*UpdateFileRequest)(nil),     // 8: pb.UpdateFileRequest
	(*FindAllRequest)(nil),        // 9: pb.FindAllRequest
}
var file_file_manager_proto_depIdxs = []int32{
	0, // 0: pb.File.info:type_name -> pb.FileMetadata
	5, // 1: pb.CreateFileRequest.filedata:type_name -> pb.CreateFileRequestData
	6, // 2: pb.UpdateFileRequest.filedata:type_name -> pb.UpdateFileRequestData
	7, // 3: pb.FileManagerService.Upload:input_type -> pb.CreateFileRequest
	3, // 4: pb.FileManagerService.Download:input_type -> pb.GetFileRequest
	3, // 5: pb.FileManagerService.Delete:input_type -> pb.GetFileRequest
	8, // 6: pb.FileManagerService.Update:input_type -> pb.UpdateFileRequest
	9, // 7: pb.FileManagerService.ListFiles:input_type -> pb.FindAllRequest
	2, // 8: pb.FileManagerService.Upload:output_type -> pb.FileRef
	1, // 9: pb.FileManagerService.Download:output_type -> pb.File
	4, // 10: pb.FileManagerService.Delete:output_type -> pb.DeleteFileResponse
	2, // 11: pb.FileManagerService.Update:output_type -> pb.FileRef
	2, // 12: pb.FileManagerService.ListFiles:output_type -> pb.FileRef
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_file_manager_proto_init() }
func file_file_manager_proto_init() {
	if File_file_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetadata); i {
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
		file_file_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_file_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRef); i {
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
		file_file_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_file_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileResponse); i {
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
		file_file_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileRequestData); i {
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
		file_file_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileRequestData); i {
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
		file_file_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileRequest); i {
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
		file_file_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileRequest); i {
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
		file_file_manager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRequest); i {
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
	file_file_manager_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*CreateFileRequest_Filedata)(nil),
		(*CreateFileRequest_Chunk)(nil),
	}
	file_file_manager_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*UpdateFileRequest_Filedata)(nil),
		(*UpdateFileRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_manager_proto_goTypes,
		DependencyIndexes: file_file_manager_proto_depIdxs,
		MessageInfos:      file_file_manager_proto_msgTypes,
	}.Build()
	File_file_manager_proto = out.File
	file_file_manager_proto_rawDesc = nil
	file_file_manager_proto_goTypes = nil
	file_file_manager_proto_depIdxs = nil
}
