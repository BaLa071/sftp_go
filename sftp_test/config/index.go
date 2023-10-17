package config

import (
	"fmt"
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func Sftp_connection() *sftp.Client {
	username := "username"
	password := "password"
	clientConfig := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", "18.216.98.58:22", clientConfig)
	if err != nil {
		log.Fatalf("SSH DAIL FAILED:%v", err)
	}
	sftpClient, err1 := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("SFTP NEW CLIENT FAILED:%v\n", err1)
	} else {
		fmt.Println("done")
	}

	return sftpClient
}
