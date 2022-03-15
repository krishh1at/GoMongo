package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishh1at/app/helpers"
	"github.com/krishh1at/app/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	result, err := user.CreateUser()
	helpers.RenderJson(w, result, err)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	err := user.Verify()
	if err != nil {
		err = errors.New("invalid email id or password")
	}

	helpers.RenderJson(w, map[string]string{"success": "Signed In successfully!"}, err)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := GetUser(w, r)
	json.NewDecoder(r.Body).Decode(user)

	result, err := models.Update(user)
	helpers.RenderJson(w, result, err)
}

func GetUser(w http.ResponseWriter, r *http.Request) *models.User {
	params := mux.Vars(r)
	result, err := models.Find(&models.User{}, params["id"])
	if err != nil {
		helpers.RenderJson(w, result, err)
		return nil
	}

	return result.(*models.User)
}
