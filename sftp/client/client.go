package main

import (
	"context"
	"io"
	"log"
	"os"
	pb "sftp/proto"
	"time"

	"google.golang.org/grpc"
)

func upload(ctx context.Context, client pb.GrpcRwServiceClient) error {
	stream, err := client.Upload(ctx)
	if err != nil {
		log.Fatal("error uploading to stream: ", err)
	}

	file, err := os.Open("/home/balaji/go/src/testint/fun1.txt")
	if err != nil {
		log.Fatal("error opening file: ", err)
	}

	buffer := make([]byte, 14)
	batchNumber := 1
	for {
		num, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error reading file: ", err)
		}
		chunk := buffer[:num]
		if err := stream.Send(&pb.FileUploadRequest{FileName:"", Chunk: chunk}); err != nil {
			log.Fatal("error sending file: ", err)
		}
		log.Printf("Sent - batch #%v - size - %v\n", batchNumber, len(chunk))
		batchNumber += 1
	}
	return nil
}

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGrpcRwServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	err1 := upload(ctx, client)
	if err1 != nil {
		log.Fatal("error client service: ", err1)
	}
	select {
	case <-ctx.Done():
		// The context has been canceled; stop uploading.
		log.Println("Context canceled. Stopping upload.")
	}
}
