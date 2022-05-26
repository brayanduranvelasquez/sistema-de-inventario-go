package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brayanduranvelasquez/sistema-de-inventario-go/commons"
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/models"
	"github.com/gorilla/mux"
)

func GetTodasDetalleSalida(writer http.ResponseWriter, request *http.Request) {
	detalleSalida := []models.Detalle_Salida{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&detalleSalida)
	json, _ := json.Marshal(detalleSalida)
	commons.SendReponse(writer, http.StatusOK, json)
}

func GetDetalleSalida(writer http.ResponseWriter, request *http.Request) {
	detalleSalida := models.Detalle_Salida{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&detalleSalida, id)

	if detalleSalida.ID > 0 {
		json, _ := json.Marshal(detalleSalida)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func SaveDetalleSalida(writer http.ResponseWriter, request *http.Request) {
	detalleSalida := models.Detalle_Salida{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&detalleSalida)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&detalleSalida).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(detalleSalida)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func DeleteDetalleSalida(writer http.ResponseWriter, request *http.Request) {
	detalleSalida := models.Detalle_Salida{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&detalleSalida, id)

	if detalleSalida.ID > 0 {
		db.Delete(detalleSalida)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
