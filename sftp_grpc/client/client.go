package main

import (
	"context"
	"log"

	pro "sftp_grpc/grpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	} else {
		log.Println("connected")
	}

	client := pro.NewFileTransferServiceClient(conn)
	getFileRequest := &pro.GetFileRequest{

		RemotePath: "/sftp/testint/sample.json",
		LocalPath:  "/home/balaji/go/src/testint/test.json",
	}

	getFileResponse, err := client.GetFile(context.Background(), getFileRequest)
	if err != nil {
		log.Fatalf("GetFile error: %v", err)
	}
	log.Println(getFileResponse)
}
