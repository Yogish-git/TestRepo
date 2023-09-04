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
func CreateUST(w http.ResponseWriter, r *http.Request) {
	UST := &models.UserStory{}
	utils.ParseBody(r, UST)
	b := UST.CreateUST()
	if b.Error != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(b.Value)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// get user story by id
func GetUSTById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ustId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	ustDetails, _ := models.GetUSTById(ID)
	res, _ := json.Marshal(ustDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// get all user story
func GetAllUST(w http.ResponseWriter, r *http.Request) {
	allUST := models.Getallust()
	res, _ := json.Marshal(allUST)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// delete ust by id
func DeleteUST(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["ustId"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedUST := models.DeleteUST(ID)
	res, _ := json.Marshal(deletedUST)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}
