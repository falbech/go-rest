package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest/database"
	"go-rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}
func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}
func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(&personality)
}
