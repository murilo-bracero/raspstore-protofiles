// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileManagerServiceClient is the client API for FileManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileManagerServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileManagerService_UploadClient, error)
	Download(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (FileManagerService_DownloadClient, error)
	Delete(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error)
	Update(ctx context.Context, opts ...grpc.CallOption) (FileManagerService_UpdateClient, error)
	ListFiles(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (FileManagerService_ListFilesClient, error)
}

type fileManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileManagerServiceClient(cc grpc.ClientConnInterface) FileManagerServiceClient {
	return &fileManagerServiceClient{cc}
}

func (c *fileManagerServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileManagerService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileManagerService_ServiceDesc.Streams[0], "/pb.FileManagerService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileManagerServiceUploadClient{stream}
	return x, nil
}

type FileManagerService_UploadClient interface {
	Send(*CreateFileRequest) error
	CloseAndRecv() (*FileRef, error)
	grpc.ClientStream
}

type fileManagerServiceUploadClient struct {
	grpc.ClientStream
}

func (x *fileManagerServiceUploadClient) Send(m *CreateFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileManagerServiceUploadClient) CloseAndRecv() (*FileRef, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileRef)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileManagerServiceClient) Download(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (FileManagerService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileManagerService_ServiceDesc.Streams[1], "/pb.FileManagerService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileManagerServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileManagerService_DownloadClient interface {
	Recv() (*File, error)
	grpc.ClientStream
}

type fileManagerServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *fileManagerServiceDownloadClient) Recv() (*File, error) {
	m := new(File)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileManagerServiceClient) Delete(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error) {
	out := new(DeleteFileResponse)
	err := c.cc.Invoke(ctx, "/pb.FileManagerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerServiceClient) Update(ctx context.Context, opts ...grpc.CallOption) (FileManagerService_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileManagerService_ServiceDesc.Streams[2], "/pb.FileManagerService/Update", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileManagerServiceUpdateClient{stream}
	return x, nil
}

type FileManagerService_UpdateClient interface {
	Send(*UpdateFileRequest) error
	CloseAndRecv() (*FileRef, error)
	grpc.ClientStream
}

type fileManagerServiceUpdateClient struct {
	grpc.ClientStream
}

func (x *fileManagerServiceUpdateClient) Send(m *UpdateFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileManagerServiceUpdateClient) CloseAndRecv() (*FileRef, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileRef)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileManagerServiceClient) ListFiles(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (FileManagerService_ListFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileManagerService_ServiceDesc.Streams[3], "/pb.FileManagerService/ListFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileManagerServiceListFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileManagerService_ListFilesClient interface {
	Recv() (*FileRef, error)
	grpc.ClientStream
}

type fileManagerServiceListFilesClient struct {
	grpc.ClientStream
}

func (x *fileManagerServiceListFilesClient) Recv() (*FileRef, error) {
	m := new(FileRef)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileManagerServiceServer is the server API for FileManagerService service.
// All implementations must embed UnimplementedFileManagerServiceServer
// for forward compatibility
type FileManagerServiceServer interface {
	Upload(FileManagerService_UploadServer) error
	Download(*GetFileRequest, FileManagerService_DownloadServer) error
	Delete(context.Context, *GetFileRequest) (*DeleteFileResponse, error)
	Update(FileManagerService_UpdateServer) error
	ListFiles(*FindAllRequest, FileManagerService_ListFilesServer) error
	mustEmbedUnimplementedFileManagerServiceServer()
}

// UnimplementedFileManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileManagerServiceServer struct {
}

func (UnimplementedFileManagerServiceServer) Upload(FileManagerService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileManagerServiceServer) Download(*GetFileRequest, FileManagerService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedFileManagerServiceServer) Delete(context.Context, *GetFileRequest) (*DeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFileManagerServiceServer) Update(FileManagerService_UpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFileManagerServiceServer) ListFiles(*FindAllRequest, FileManagerService_ListFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFileManagerServiceServer) mustEmbedUnimplementedFileManagerServiceServer() {}

// UnsafeFileManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileManagerServiceServer will
// result in compilation errors.
type UnsafeFileManagerServiceServer interface {
	mustEmbedUnimplementedFileManagerServiceServer()
}

func RegisterFileManagerServiceServer(s grpc.ServiceRegistrar, srv FileManagerServiceServer) {
	s.RegisterService(&FileManagerService_ServiceDesc, srv)
}

func _FileManagerService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileManagerServiceServer).Upload(&fileManagerServiceUploadServer{stream})
}

type FileManagerService_UploadServer interface {
	SendAndClose(*FileRef) error
	Recv() (*CreateFileRequest, error)
	grpc.ServerStream
}

type fileManagerServiceUploadServer struct {
	grpc.ServerStream
}

func (x *fileManagerServiceUploadServer) SendAndClose(m *FileRef) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileManagerServiceUploadServer) Recv() (*CreateFileRequest, error) {
	m := new(CreateFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileManagerService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileManagerServiceServer).Download(m, &fileManagerServiceDownloadServer{stream})
}

type FileManagerService_DownloadServer interface {
	Send(*File) error
	grpc.ServerStream
}

type fileManagerServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *fileManagerServiceDownloadServer) Send(m *File) error {
	return x.ServerStream.SendMsg(m)
}

func _FileManagerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FileManagerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServiceServer).Delete(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManagerService_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileManagerServiceServer).Update(&fileManagerServiceUpdateServer{stream})
}

type FileManagerService_UpdateServer interface {
	SendAndClose(*FileRef) error
	Recv() (*UpdateFileRequest, error)
	grpc.ServerStream
}

type fileManagerServiceUpdateServer struct {
	grpc.ServerStream
}

func (x *fileManagerServiceUpdateServer) SendAndClose(m *FileRef) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileManagerServiceUpdateServer) Recv() (*UpdateFileRequest, error) {
	m := new(UpdateFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileManagerService_ListFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileManagerServiceServer).ListFiles(m, &fileManagerServiceListFilesServer{stream})
}

type FileManagerService_ListFilesServer interface {
	Send(*FileRef) error
	grpc.ServerStream
}

type fileManagerServiceListFilesServer struct {
	grpc.ServerStream
}

func (x *fileManagerServiceListFilesServer) Send(m *FileRef) error {
	return x.ServerStream.SendMsg(m)
}

// FileManagerService_ServiceDesc is the grpc.ServiceDesc for FileManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FileManagerService",
	HandlerType: (*FileManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _FileManagerService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileManagerService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _FileManagerService_Download_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Update",
			Handler:       _FileManagerService_Update_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ListFiles",
			Handler:       _FileManagerService_ListFiles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file-manager.proto",
}
