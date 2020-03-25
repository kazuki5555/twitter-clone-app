package controllers

import (
	"github.com/kazuki5555/twitter-clone-app/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("[controllers/favorites]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewFavorites(r *http.Request) error {
	m := models.NewSqlHandler()
	users, err := NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		return err
	}

	favorites := models.GetFavorites()
	favorites.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&favorites)
	if err != nil {
		log.Println(err)
		return err
	}

	if r.Method == "POST" {
		err = CreateFavorite(favorites, m)
	} else if r.Method == "DELETE" {
		err = DeleteFavorite(favorites, m)
	}
	if err != nil {
		return err
	}

	return nil
}

func CreateFavorite(f models.Favorites, m *models.DB) error {
	err := f.CreateFavorite(m)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFavorite(f models.Favorites, m *models.DB) error {
	err := f.DeleteFavorite(m)
	if err != nil {
		return err
	}
	return nil
}
