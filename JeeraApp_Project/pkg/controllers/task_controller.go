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

func CreateTask(w http.ResponseWriter, r *http.Request) {
	Task := &models.Task{}
	utils.ParseBody(r, Task)
	b := Task.CreateTask()
	if b.Error != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(b.Value)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["TaskId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	TaskDetails, _ := models.GetTaskById(ID)
	res, _ := json.Marshal(TaskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	allTask := models.GetAllTask()
	res, _ := json.Marshal(allTask)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["TaskId"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedTask := models.DeleteTask(ID)
	res, _ := json.Marshal(deletedTask)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}
