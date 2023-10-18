package models

type GetFileRequest struct {
	SftpServer string `json:"sftp_server"`
	RemotePath string `json:"remote_path"`
	LocalPath  string `json:"local_path"`
}

type GetFileResponse struct {
	Message string `json:"message"`
}

type PutFileRequest struct {
	SftpServer string `json:"sftp_server"`
	LocalPath  string `json:"local_path"`
	RemotePath string `json:"remote_path"`
}

type PutFileResponse struct {
	Message string `json:"message"`
}
