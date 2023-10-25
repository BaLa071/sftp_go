package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	pb "sftp/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGrpcRwServiceServer
}

type File struct {
	FilePath   string
	buffer     *bytes.Buffer
	OutputFile *os.File
}

func NewFile() *File {
	return &File{
		buffer: &bytes.Buffer{},
	}
}

func (s *server) Upload(stream pb.GrpcRwService_UploadServer) error {
	file := NewFile()
	file.OutputFile, _ = os.Create("/home/balaji/go/src/fun.txt")
	var fileSize uint32
	fileSize = 0
	defer func() {
		if err := file.OutputFile.Close(); err != nil {
			log.Fatal("err")
		}
	}()

	for {
		fmt.Println("1")
		req, err := stream.Recv()
		file.FilePath = req.GetFileName()
		fmt.Println(req.GetFileName())
		fmt.Println(file.FilePath)

		// if file.FilePath == "" {
		//  log.Fatal("server file path is empty")
		// }
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("err: ", err)
		}

		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))
		fmt.Println("Received chunk with size: ", fileSize)
		fmt.Println(file.OutputFile)
		_, err1 := file.OutputFile.Write(chunk)
		if err1 != nil {
			log.Fatal("error writing chunk: ", err1)
		}
	}

	fmt.Println(file.FilePath, fileSize)
	fileName := filepath.Base(file.FilePath)
	fmt.Printf("saved file: %s, size: %d", fileName, fileSize)
	return stream.SendAndClose(&pb.FileUploadResponse{
		FileName: fileName,
		Size:     fileSize,
	})
}

func main() {
	fmt.Println("gRPC server listening on")
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGrpcRwServiceServer(s, &server{})
	fmt.Println("Server listening on: 8000")

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
