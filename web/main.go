package main

import (
	"fmt"
	"log"
	"net/http"

	// "io/ioutil"

	"github.com/gorilla/mux"
)

const (
	helloPath = "/hello"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	if users, err := getUsers(); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Hello, world!\n"))
		w.Write([]byte("We have users:\n"))
		for _, user := range users {
			ustr := fmt.Sprintf(" -> %d - %s\n", user.Id, user.Name)
			w.Write([]byte(ustr))
		}
		if len(users) == 0 {
			w.Write([]byte(" No users"))
		}
	}
}

func Start(host string, port int32) {
	router := mux.NewRouter()

	router.HandleFunc(helloPath, getHello).Methods(http.MethodGet)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	addr := fmt.Sprintf("%s:%d", host, port)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := GetConfig()

	Start(config.Host, config.Port)
}
