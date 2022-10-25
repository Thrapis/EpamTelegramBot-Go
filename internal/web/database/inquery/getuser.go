package inquery

import (
	model "baa-telebot/internal/database"
	src "baa-telebot/internal/web/database/source"
	sql "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetUsers() ([]*model.User, error) {

	datasourceName := src.GetSource()

	db, err := sql.Open("mysql", datasourceName)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for results.Next() {
		var u model.User

		err = results.Scan(&u.Id, &u.Name)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}
