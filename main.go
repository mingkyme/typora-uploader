package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, file := range args {
		byteArray, err := ioutil.ReadFile(file)
		if err != nil {
			return
		}
		res, err := http.Post("http://localhost:3000/upload", "image/png", bytes.NewReader(byteArray))
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		bodyByteArray, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error")
		}
		str := string(bodyByteArray)
		fmt.Println(str)
	}
}
