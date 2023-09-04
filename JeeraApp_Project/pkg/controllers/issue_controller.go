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

func CreateIssue(w http.ResponseWriter, r *http.Request) {
	Issue := &models.Issue{}
	utils.ParseBody(r, Issue)
	b := Issue.CreateIssue()
	if b.Error != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(b.Value)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetIssueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["IssueId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	IssueDetails, _ := models.GetIssueById(ID)
	res, _ := json.Marshal(IssueDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllIssue(w http.ResponseWriter, r *http.Request) {
	allIssue := models.GetAllIssue()
	res, _ := json.Marshal(allIssue)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["IssueId"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedIssue := models.DeleteIssue(ID)
	res, _ := json.Marshal(deletedIssue)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}
