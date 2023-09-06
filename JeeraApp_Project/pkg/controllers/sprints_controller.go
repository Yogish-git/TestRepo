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

func CreateSprint(w http.ResponseWriter, r *http.Request) {
	Sprint := &models.Sprint{}
	utils.ParseBody(r, Sprint)
	b := Sprint.CreateSprint()
	if b.Error != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(b.Value)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSprintById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["SprintId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	SprintDetails, _ := models.GetSprintById(ID)
	res, _ := json.Marshal(SprintDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllSprint(w http.ResponseWriter, r *http.Request) {
	allSprint := models.GetAllSprint()
	res, _ := json.Marshal(allSprint)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSprint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["SprintId"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedSprint := models.DeleteSprint(ID)
	res, _ := json.Marshal(deletedSprint)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}
