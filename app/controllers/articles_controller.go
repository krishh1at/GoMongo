package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/krishh1at/app/helpers"
	"github.com/krishh1at/app/models"
)

// Get All Articles API
func GetArticles(w http.ResponseWriter, r *http.Request) {
	result, err := models.All(&models.Article{})
	helpers.RenderJson(w, result, err)
}

// Create Article API
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var article models.Article
	_ = json.NewDecoder(r.Body).Decode(&article)

	result, err := models.InsertOne(&article)
	helpers.RenderJson(w, result, err)
}
