package main

import (
	restapi "baa-telebot/internal/api"
	cfg "baa-telebot/internal/config"
	handler "baa-telebot/internal/web/handler"
	"fmt"
	"log"
	"net/http"

	// "io/ioutil"

	"github.com/gorilla/mux"
)

func Start(host string, port int32) {
	router := mux.NewRouter()

	router.HandleFunc(restapi.HelloUrl, handler.GetHello).Methods(http.MethodGet)
	router.HandleFunc(restapi.TimeUrl, handler.GetTime).Methods(http.MethodGet)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	addr := fmt.Sprintf("%s:%d", host, port)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := cfg.GetConfig()
	Start(config.Server.Host, config.Server.Port)
}
