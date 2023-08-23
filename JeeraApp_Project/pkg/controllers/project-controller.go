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

var NewProject models.Project

func GetProject(w http.ResponseWriter, r *http.Request) {
	newProjects := models.GetAllProjects()
	res, _ := json.Marshal(newProjects)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader((http.StatusOK))
	w.Write(res)
}

func GetProjectById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProjectId := vars["projectId"]
	ID, err := strconv.ParseInt(ProjectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing = ")
		fmt.Println(err)
	}
	projectDetails, _ := models.GetProjectById(ID)
	res, _ := json.Marshal(projectDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	CreateProject := &models.Project{}
	utils.ParseBody(r, CreateProject)
	b := CreateProject.CreateProject()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["Id"]
	ID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	project := models.DeleteProject(ID)
	res, _ := json.Marshal(project)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var UpdateProject = &models.Project{}
	utils.ParseBody(r, UpdateProject)
	vars := mux.Vars(r)
	projectId := vars["Id"]
	ID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	projectDetails, db := models.GetProjectById(ID)
	if UpdateProject.Name != "" {
		projectDetails.Name = UpdateProject.Name
	}
	//other update functions you add similarly
	db.Save(&projectDetails)
	res, _ := json.Marshal(projectDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
