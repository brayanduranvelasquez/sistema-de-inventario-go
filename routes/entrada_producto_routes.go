package routes

import (
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/controllers"
	"github.com/gorilla/mux"
)

func SetEntradaProductoRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/entrada-producto/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetTodasEntradaProducto).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveEntradaProducto).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteEntradaProducto).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetEntradaProducto).Methods("GET")
}
