package routes

import (
	"github.com/DhruvinShiroya/go-toolrentalstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterToolrentalRoutes = func(router *mux.Router) {
	router.HandleFunc("/tool/", controllers.CreateTool).Methods("POST")
	router.HandleFunc("/tool/", controllers.GetTool).Methods("GET")
	router.HandleFunc("/tool/{toolId}", controllers.GetToolById).Methods("GET")
	router.HandleFunc("/tool/{toolId}", controllers.UpdateTool).Methods("PUT")
	router.HandleFunc("/tool/{toolId}", controllers.DeleteTool).Methods("DELETE")
}
