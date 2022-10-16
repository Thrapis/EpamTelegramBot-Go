package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	helloPath = "/hello"
)

func HelloRequest(server ServerConfig) (result string) {
	var url = fmt.Sprintf("http://%s:%d%s", server.Host, server.Port, helloPath)
	if response, err := http.Get(url); err != nil {
		result = fmt.Sprintf("%s", err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		result = string(contents)
	}
	return
}
