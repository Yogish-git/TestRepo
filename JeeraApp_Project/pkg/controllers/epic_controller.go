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

func CreateEpic(w http.ResponseWriter, r *http.Request) {
	Epic := &models.EPIC{}
	utils.ParseBody(r, Epic)
	b := Epic.CreateEpic()
	if b.Error != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(b.Value)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEpicById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["EpicId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	EpicDetails, _ := models.GetEpicById(ID)
	res, _ := json.Marshal(EpicDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllEpic(w http.ResponseWriter, r *http.Request) {
	allEpic := models.GetAllEpic()
	res, _ := json.Marshal(allEpic)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEpic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["EpicId"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedEpic := models.DeleteEpic(ID)
	res, _ := json.Marshal(deletedEpic)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}
