package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yogish-git/JeeraApp_Project/pkg/models"
	"github.com/yogish-git/JeeraApp_Project/pkg/utils"
)

// create userstory
func CreateEffort(w http.ResponseWriter, r *http.Request) {
	Effort := &models.EffortLogging{}
	utils.ParseBody(r, Effort)
	b := Effort.CreateEffort()
	if b.Error != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(b.Value)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// get user story by id
func GetEffortById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["effortId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	effortDetails, _ := models.GetEffortById(ID)
	res, _ := json.Marshal(effortDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// get all user story
func GetAllEffort(w http.ResponseWriter, r *http.Request) {
	allEffort := models.Getalleffort()
	res, _ := json.Marshal(allEffort)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// delete ust by id
func DeleteEffort(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["effortId"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedEffort := models.DeleteEffort(ID)
	res, _ := json.Marshal(deletedEffort)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}
