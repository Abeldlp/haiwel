package service

import (
	"github.com/Abeldlp/haiwel/user-service/config"
	"github.com/Abeldlp/haiwel/user-service/model"
	"log"
)

func GetAll() ([]model.User, error) {
	qry := "SELECT * FROM users"
	var users []model.User

	rows, err := config.DB.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Username, &user.Rank)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return users, nil
}

func Create(user model.User) int64 {
	result, err := config.DB.Exec("INSERT INTO users (username, rank) VALUES ($1, $2);", user.Username, user.Rank)
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Error getting rows affected: %v", err)
	}

	return rows
}
