package user

import (
	cfg "baa-telebot/config"
	model "baa-telebot/internal/database"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetUsers() ([]*model.User, error) {
	config := cfg.GetConfig()

	// "inst_user:secret@tcp(db:3306)/inst_database"
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Database.User, config.Database.Password, config.Database.Host,
		config.Database.Port, config.Database.Instance)

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
