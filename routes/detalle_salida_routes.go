package routes

import (
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/controllers"
	"github.com/gorilla/mux"
)

func SetDetalleSalidaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/detalle-salida/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetTodasDetalleSalida).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveDetalleSalida).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteDetalleSalida).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetDetalleSalida).Methods("GET")
}
