package routes

import (
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/controllers"
	"github.com/gorilla/mux"
)

func SetDetalleEntradaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/detalle-salida/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetTodasDetalleEntrada).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveDetalleEntrada).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteDetalleEntrada).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetDetalleEntrada).Methods("GET")
}
