package request

import (
	api "baa-telebot/internal/api"
	cfg "baa-telebot/internal/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloRequest(server cfg.ServerConfig) (result string) {
	var url = fmt.Sprintf("http://%s:%d%s", server.Host, server.Port, api.HelloUrl)

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
