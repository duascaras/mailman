package main

import (
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
)

func main() {
	// Credentials
	tgHostname := ""
	tgUsername := ""
	tgPassword := ""

	// Connect to FTP server
	ftpClient, err := ftpConnect(tgHostname, tgUsername, tgPassword)
	if err != nil {
		fmt.Println("Error connecting to FTP server:", err)
		return
	}
	defer ftpClient.Quit()

	// List local files in the current directory
	localDir := ""
	files, err := os.ReadDir(localDir)
	if err != nil {
		fmt.Println("Error listing local files:", err)
		return
	}

	// Upload files
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		localFilePath := localDir + "/" + file.Name()
		remoteFilePath := file.Name()

		localFile, err := os.Open(localFilePath)
		if err != nil {
			fmt.Println("Error opening local file:", err)
			return
		}
		defer localFile.Close()

		err = ftpClient.Stor(remoteFilePath, localFile)
		if err != nil {
			fmt.Println("Error uploading file:", err)
			return
		}
		fmt.Println("Uploaded:", file.Name())
	}

	// List files on the FTP server
	remoteFiles, err := ftpClient.List("")
	if err != nil {
		fmt.Println("Error listing remote files:", err)
		return
	}
	fmt.Println("Remote Files:")
	for _, remoteFile := range remoteFiles {
		fmt.Println(remoteFile.Name)
	}
}

func ftpConnect(hostname, username, password string) (*ftp.ServerConn, error) {
	client, err := ftp.Dial(hostname + ":21")
	if err != nil {
		return nil, err
	}

	err = client.Login(username, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}
