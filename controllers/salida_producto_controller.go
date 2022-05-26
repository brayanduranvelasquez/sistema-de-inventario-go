package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brayanduranvelasquez/sistema-de-inventario-go/commons"
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/models"
	"github.com/gorilla/mux"
)

func GetTodasSalidaProducto(writer http.ResponseWriter, request *http.Request) {
	salidaProducto := []models.Salida_Producto{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&salidaProducto)
	json, _ := json.Marshal(salidaProducto)
	commons.SendReponse(writer, http.StatusOK, json)
}

func GetSalidaProducto(writer http.ResponseWriter, request *http.Request) {
	salidaProducto := models.Salida_Producto{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&salidaProducto, id)

	if salidaProducto.ID > 0 {
		json, _ := json.Marshal(salidaProducto)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func SaveSalidaProducto(writer http.ResponseWriter, request *http.Request) {
	salidaProducto := models.Salida_Producto{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&salidaProducto)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&salidaProducto).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(salidaProducto)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func DeleteSalidaProducto(writer http.ResponseWriter, request *http.Request) {
	salidaProducto := models.Salida_Producto{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&salidaProducto, id)

	if salidaProducto.ID > 0 {
		db.Delete(salidaProducto)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
