package routes

import (
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/controllers"
	"github.com/gorilla/mux"
)

func SetSalidaProductoRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/salida-producto/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetTodasSalidaProducto).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveSalidaProducto).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteSalidaProducto).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetSalidaProducto).Methods("GET")
}
