package routes

import (
	"github.com/gorilla/mux"
	"github.com/yogish-git/JeeraApp_Project/pkg/controllers"
)

var RegisterProjectRoutes = func(router *mux.Router) {
	router.HandleFunc("/project/", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/project/", controllers.GetProject).Methods("GET")
	router.HandleFunc("/project/{projectId}", controllers.GetProjectById).Methods("GET")
	router.HandleFunc("/project/{projectId}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/project/{projectId}", controllers.DeleteProject).Methods("DELETE")

}
