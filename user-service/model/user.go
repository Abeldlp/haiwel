package model

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Rank     int    `json:"user_ranking_id"`
}
