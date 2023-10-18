// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: grpc/list.proto

package BaLa071

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

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SftpServer string `protobuf:"bytes,1,opt,name=sftp_server,json=sftpServer,proto3" json:"sftp_server,omitempty"`
	RemotePath string `protobuf:"bytes,2,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
	LocalPath  string `protobuf:"bytes,3,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_list_proto_msgTypes[0]
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
	return file_grpc_list_proto_rawDescGZIP(), []int{0}
}

func (x *GetFileRequest) GetSftpServer() string {
	if x != nil {
		return x.SftpServer
	}
	return ""
}

func (x *GetFileRequest) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

func (x *GetFileRequest) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_grpc_list_proto_rawDescGZIP(), []int{1}
}

func (x *GetFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PutFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SftpServer string `protobuf:"bytes,1,opt,name=sftp_server,json=sftpServer,proto3" json:"sftp_server,omitempty"`
	LocalPath  string `protobuf:"bytes,2,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
	RemotePath string `protobuf:"bytes,3,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
}

func (x *PutFileRequest) Reset() {
	*x = PutFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileRequest) ProtoMessage() {}

func (x *PutFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileRequest.ProtoReflect.Descriptor instead.
func (*PutFileRequest) Descriptor() ([]byte, []int) {
	return file_grpc_list_proto_rawDescGZIP(), []int{2}
}

func (x *PutFileRequest) GetSftpServer() string {
	if x != nil {
		return x.SftpServer
	}
	return ""
}

func (x *PutFileRequest) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

func (x *PutFileRequest) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

type PutFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PutFileResponse) Reset() {
	*x = PutFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileResponse) ProtoMessage() {}

func (x *PutFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileResponse.ProtoReflect.Descriptor instead.
func (*PutFileResponse) Descriptor() ([]byte, []int) {
	return file_grpc_list_proto_rawDescGZIP(), []int{3}
}

func (x *PutFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_list_proto protoreflect.FileDescriptor

var file_grpc_list_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22,
	0x71, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x66, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61,
	0x74, 0x68, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x71, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x66, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x22, 0x2b, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xa5, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x4c, 0x61, 0x30, 0x37, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_list_proto_rawDescOnce sync.Once
	file_grpc_list_proto_rawDescData = file_grpc_list_proto_rawDesc
)

func file_grpc_list_proto_rawDescGZIP() []byte {
	file_grpc_list_proto_rawDescOnce.Do(func() {
		file_grpc_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_list_proto_rawDescData)
	})
	return file_grpc_list_proto_rawDescData
}

var file_grpc_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_list_proto_goTypes = []interface{}{
	(*GetFileRequest)(nil),  // 0: filetransfer.GetFileRequest
	(*GetFileResponse)(nil), // 1: filetransfer.GetFileResponse
	(*PutFileRequest)(nil),  // 2: filetransfer.PutFileRequest
	(*PutFileResponse)(nil), // 3: filetransfer.PutFileResponse
}
var file_grpc_list_proto_depIdxs = []int32{
	0, // 0: filetransfer.FileTransferService.GetFile:input_type -> filetransfer.GetFileRequest
	2, // 1: filetransfer.FileTransferService.PutFile:input_type -> filetransfer.PutFileRequest
	1, // 2: filetransfer.FileTransferService.GetFile:output_type -> filetransfer.GetFileResponse
	3, // 3: filetransfer.FileTransferService.PutFile:output_type -> filetransfer.PutFileResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_list_proto_init() }
func file_grpc_list_proto_init() {
	if File_grpc_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResponse); i {
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
		file_grpc_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileRequest); i {
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
		file_grpc_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileResponse); i {
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
			RawDescriptor: file_grpc_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_list_proto_goTypes,
		DependencyIndexes: file_grpc_list_proto_depIdxs,
		MessageInfos:      file_grpc_list_proto_msgTypes,
	}.Build()
	File_grpc_list_proto = out.File
	file_grpc_list_proto_rawDesc = nil
	file_grpc_list_proto_goTypes = nil
	file_grpc_list_proto_depIdxs = nil
}