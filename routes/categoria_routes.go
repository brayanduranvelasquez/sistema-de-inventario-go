package routes

import (
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/controllers"
	"github.com/gorilla/mux"
)

func SetCategoriaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/categoria/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetTodasCategorias).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveCategoria).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteCategoria).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetCategoria).Methods("GET")
}
