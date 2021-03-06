package controllers

import (
	"log"
	"net/http"

	"github.com/kazuki5555/twitter-clone-app/models"

	"github.com/gorilla/mux"
)

type Timeline struct {
	TweetID    int    `json:"tweet_id"`
	UserID     int    `json:"user_id"`
	ScreenName string `json:"screen_name"`
	ScreenID   string `json:"screen_id"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
}

func init() {
	log.SetPrefix("[controllers/timeline]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewTimeline(r *http.Request) ([]Timeline, error) {
	m := models.GetSqlHandler()
	users, err := NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		return nil, err
	}

	timeline := Timeline{}
	timelines, err := timeline.GetTimeline(users, m)
	if err != nil {
		return timelines, err
	}

	return timelines, nil
}

func (t *Timeline) GetTimeline(users models.Users, m *models.DB) ([]Timeline, error) {
	followers, err := models.GetFollowingByUid(users.ID, m)
	if err != nil {
		return nil, err
	}

	//timeline取得のために自分も入れる
	userIds := []int{users.ID}
	for _, s := range followers {
		userIds = append(userIds, s.FollowedID)
	}

	var timelines []Timeline
	rows, err := models.GetTimelineByUserId(userIds, m)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		m.DB.ScanRows(rows, &t)
		timelines = append(timelines, *t)
	}
	return timelines, nil
}
