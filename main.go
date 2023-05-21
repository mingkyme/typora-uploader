package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// typora-uploader http://upload.com a.jpg
func main() {
	if len(os.Args) <= 1 {
		fmt.Println("need arguments")
		fmt.Println("example: typora-uploader http://upload.com a.jpg")
		fmt.Println("check https://github.com/mingkyme/typora-uploader")
		return
	}
	serverURL := os.Args[1]
	args := os.Args[2:]

	for _, filePath := range args {
		file, _ := os.Open(filePath)
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
		io.Copy(part, file)
		writer.Close()

		r, _ := http.NewRequest("POST", serverURL, body)
		r.Header.Add("Content-Type", writer.FormDataContentType())
		client := &http.Client{}
		respone, err := client.Do(r)
		if err != nil {
			fmt.Println(err)
		}
		bodyBytes, err := ioutil.ReadAll(respone.Body)
		if err != nil {
			fmt.Println(err)
		}
		str := string(bodyBytes)
		fmt.Println(str)
	}
}
