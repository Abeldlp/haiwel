package service

import (
	"github.com/Abeldlp/haiwel/user-service/config"
	"github.com/Abeldlp/haiwel/user-service/model"
)

func GetAllRankings() ([]model.UserRanking, error) {
	qry := "SELECT * FROM user_rankings"
	var userRankings []model.UserRanking

	rows, err := config.DB.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var userRanking model.UserRanking
		err := rows.Scan(&userRanking.Id, &userRanking.GameId)
		if err != nil {
			panic(err.Error())
		}
		userRankings = append(userRankings, userRanking)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return userRankings, nil
}
