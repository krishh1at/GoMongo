package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/krishh1at/app/helpers"
	"github.com/krishh1at/app/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	user := &models.User{}
	_ = json.NewDecoder(r.Body).Decode(user)

	user, err := user.CheckEmailExist()
	if err != nil {
		helpers.RenderJson(w, user, err)
	}

	user, err = user.Encrypt()
	if err != nil {
		helpers.RenderJson(w, user, err)
	}

	result, err := models.InsertOne(user)
	helpers.RenderJson(w, result, err)
}
