package handler

import (
	dbiq "baa-telebot/internal/web/database/inquery"

	"fmt"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	if users, err := dbiq.GetUsers(); err != nil {
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
