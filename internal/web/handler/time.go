package handler

import (
	"time"

	"net/http"
)

func GetTime(w http.ResponseWriter, r *http.Request) {
	current := time.Now().Format(time.UnixDate)
	w.Write([]byte(current))
}
