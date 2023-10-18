package service

import (
	"context"
	"io"
	"log"
	"os"
	"sftp_grpc/config"
	pro "sftp_grpc/grpc"
)

type Server struct {
	pro.UnimplementedFileTransferServiceServer
}

func (s *Server) GetFile(ctx context.Context, req *pro.GetFileRequest) (*pro.GetFileResponse, error) {

	Client := config.Sftp_connection()

	remoteFile, err := Client.Open(req.RemotePath)
	if err != nil {
		return nil, err
	}
	log.Println("opened")
	localFile, err := os.Create(req.LocalPath)
	if err != nil {
		return nil, err
	}
	log.Println("created")
	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		return nil, err
	}

	return &pro.GetFileResponse{Message: "File downloaded successfully"}, nil
}
