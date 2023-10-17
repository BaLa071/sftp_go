package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sftp_test/models"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
)

func Put(client *sftp.Client) {

	localFile, err := os.Open("/home/balaji/go/src/demo/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file got read!", localFile)

	remoteFile, err2 := client.Create("/sftp/testint/test.txt")
	if err != nil {
		log.Fatal(err2)
	}
	log.Println("created")

	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("file uploaded")
}

func Get(client *sftp.Client) {

	remoteFile, err := client.Open("/sftp/testint/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Opened")

	localFile, err := os.Create("/home/balaji/go/src/demo/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("created")

	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("file uploaded")
}

func List(c *gin.Context, client *sftp.Client) {

	remoteDir := "/sftp/testint"
	files, err := client.ReadDir(remoteDir)
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
		// log.Println(file.Name())
	}
	response := models.FileResponse{Files: fileNames}
	c.JSON(http.StatusOK, response)
}
