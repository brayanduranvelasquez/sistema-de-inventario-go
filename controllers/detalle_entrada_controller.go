package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brayanduranvelasquez/sistema-de-inventario-go/commons"
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/models"
	"github.com/gorilla/mux"
)

func GetTodasDetalleEntrada(writer http.ResponseWriter, request *http.Request) {
	detalleEntrada := []models.Detalle_Entrada{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&detalleEntrada)
	json, _ := json.Marshal(detalleEntrada)
	commons.SendReponse(writer, http.StatusOK, json)
}

func GetDetalleEntrada(writer http.ResponseWriter, request *http.Request) {
	detalleEntrada := models.Detalle_Entrada{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&detalleEntrada, id)

	if detalleEntrada.ID > 0 {
		json, _ := json.Marshal(detalleEntrada)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func SaveDetalleEntrada(writer http.ResponseWriter, request *http.Request) {
	detalleEntrada := models.Detalle_Entrada{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&detalleEntrada)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&detalleEntrada).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(detalleEntrada)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func DeleteDetalleEntrada(writer http.ResponseWriter, request *http.Request) {
	detalleEntrada := models.Detalle_Entrada{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&detalleEntrada, id)

	if detalleEntrada.ID > 0 {
		db.Delete(detalleEntrada)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
