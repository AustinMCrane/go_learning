package models

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

var (
	// RootFilePath default volume path
	RootFilePath  = "appvolumes"
	FileServerURL = "http://localhost:8080/"
)

func init() {
	// check if user wants to use something besides default file path
	filePath := os.Getenv("ROOT_FILE_PATH")
	if filePath != "" {
		RootFilePath = filePath
	}

	// check if user wants to use something besides default file path
	fileServerURL := os.Getenv("FILE_SERVER_URL")
	if fileServerURL != "" {
		FileServerURL = fileServerURL
	}

	_, err := os.Stat(RootFilePath)
	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		err := os.Mkdir(RootFilePath, os.FileMode(0700))
		if err != nil {
			log.Panic(err)
		}
	}
}

func SaveString(s string) (string, error) {
	err := ioutil.WriteFile(RootFilePath+"/something.txt", []byte(s), 0644)
	if err != nil {
		return "", err
	}
	return RootFilePath + "/something.txt", nil
}

func SaveBase64Image(fileName string, s string) (string, error) {
	fmt.Println("TESTING THAT ITS A PNG")
	unbased, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(RootFilePath+"/"+fileName, os.O_WRONLY|os.O_CREATE, 0777)
	fmt.Println(f)
	fmt.Println(im)
	if err != nil {
		return "", err
	}

	png.Encode(f, im)
	return RootFilePath + "/" + fileName, nil
}
