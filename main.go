package main

import (
	"log"
	"net/http"

	"github.com/brayanduranvelasquez/sistema-de-inventario-go/commons"
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetProductoRoutes(router)
	routes.SetCategoriaRoutes(router)
	routes.SetDetalleEntradaRoutes(router)
	routes.SetDetalleSalidaRoutes(router)
	routes.SetEntradaProductoRoutes(router)
	routes.SetSalidaProductoRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")
	log.Println(server.ListenAndServe())
}
