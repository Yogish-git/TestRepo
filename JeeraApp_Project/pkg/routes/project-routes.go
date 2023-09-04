package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogish-git/JeeraApp_Project/pkg/controllers"
	"github.com/yogish-git/JeeraApp_Project/pkg/signuplogin"
)

var RegisterProjectRoutes = func(router *mux.Router) {
	router.HandleFunc("/login/", signuplogin.LoginHandler).Methods("POST")
	router.HandleFunc("/signup/", signuplogin.SignupHandler).Methods("POST")

	router.Handle("/project/", signuplogin.AuthorizeMiddleware(http.HandlerFunc(controllers.GetProject))).Methods("GET")
	router.Handle("/project/", signuplogin.AuthorizeMiddleware(http.HandlerFunc(controllers.CreateProject))).Methods("POST")
	router.Handle("/project/{projectId}", signuplogin.AuthorizeMiddleware(http.HandlerFunc(controllers.GetProjectById))).Methods("GET")
	router.Handle("/project/{projectId}", signuplogin.AuthorizeMiddleware(http.HandlerFunc(controllers.UpdateProject))).Methods("PUT")
	router.Handle("/project/{projectId}", signuplogin.AuthorizeMiddleware(http.HandlerFunc(controllers.DeleteProject))).Methods("DELETE")

	router.HandleFunc("/userstory/", controllers.CreateUST).Methods("POST")
	router.HandleFunc("/userstory/{ustId}", controllers.GetUSTById).Methods("GET")
	router.HandleFunc("/userstory/", controllers.GetAllUST).Methods("GET")
	router.HandleFunc("/userstory/{ustId}", controllers.DeleteUST).Methods("DELETE")

	router.HandleFunc("/effortlogging/", controllers.CreateEffort).Methods("POST")
	router.HandleFunc("/effortlogging/{effortId}", controllers.GetEffortById).Methods("GET")
	router.HandleFunc("/effortlogging/", controllers.GetAllEffort).Methods("GET")
	router.HandleFunc("/effortlogging/{effortId}", controllers.DeleteEffort).Methods("DELETE")
}
