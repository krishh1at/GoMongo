package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/krishh1at/app/helpers"
	"github.com/krishh1at/app/models"
	"github.com/krishh1at/app/services"
)

// User signup API
func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	result, err := user.CreateUser()
	helpers.RenderJson(w, result, err)
}

// User signin API
func SignIn(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	varifiedUser, err := user.Verify()
	if err != nil {
		err = errors.New("invalid email id or password")
	}

	token, err := services.GetJWT(varifiedUser.ID.String())
	helpers.RenderJson(w, map[string]string{"token": token}, err)
}
