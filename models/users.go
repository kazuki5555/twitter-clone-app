package models

import (
	"database/sql"
	"log"
)

type Users struct {
	ID         int    `json:"id"`
	ScreenName string `json:"screen_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func init() {
	log.SetPrefix("[users]]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetTimelineByUserId(users []int, m *DB) *sql.Rows {
	query := m.DB.Table("users").
		Select("users.id as user_id, users.screen_name, tweets.id as tweet_id, tweets.text, tweets.created_at").
		Joins("left join tweets on users.id = tweets.user_id").
		Where("users.id in (?) and tweets.is_daleted = 0", users)
	rows, err := query.Rows()
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
