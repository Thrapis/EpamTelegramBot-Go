package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() ([]*User, error) {
	config := GetConfig()

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Database.User, config.Database.Password, config.Database.Host,
		config.Database.Port, config.Database.Instance)

	// "inst_user:secret@tcp(db:3306)/inst_database"

	db, err := sql.Open("mysql", datasourceName)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []*User
	for results.Next() {
		var u User

		err = results.Scan(&u.Id, &u.Name)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}
