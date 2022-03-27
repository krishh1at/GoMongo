package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishh1at/app/helpers"
	"github.com/krishh1at/app/models"
)

// Update User API
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := GetUser(w, r)
	json.NewDecoder(r.Body).Decode(user)

	result, err := models.Update(user)
	helpers.RenderJson(w, result, err)
}

// GET User API
func GetUser(w http.ResponseWriter, r *http.Request) *models.User {
	params := mux.Vars(r)
	result, err := models.Find(&models.User{}, params["id"])
	if err != nil {
		helpers.RenderJson(w, result, err)
		return nil
	}

	return result.(*models.User)
}
