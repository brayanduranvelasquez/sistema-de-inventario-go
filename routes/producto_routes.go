package routes

import (
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/controllers"
	"github.com/gorilla/mux"
)

func SetProductoRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/producto/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetTodosProductos).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveProducto).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteProducto).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetProducto).Methods("GET")
}
