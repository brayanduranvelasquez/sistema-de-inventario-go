package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brayanduranvelasquez/sistema-de-inventario-go/commons"
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/models"
	"github.com/gorilla/mux"
)

func GetTodasEntradaProducto(writer http.ResponseWriter, request *http.Request) {
	entradaProducto := []models.Entrada_Producto{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&entradaProducto)
	json, _ := json.Marshal(entradaProducto)
	commons.SendReponse(writer, http.StatusOK, json)
}

func GetEntradaProducto(writer http.ResponseWriter, request *http.Request) {
	entradaProducto := models.Entrada_Producto{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&entradaProducto, id)

	if entradaProducto.ID > 0 {
		json, _ := json.Marshal(entradaProducto)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func SaveEntradaProducto(writer http.ResponseWriter, request *http.Request) {
	entradaProducto := models.Entrada_Producto{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&entradaProducto)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&entradaProducto).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(entradaProducto)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func DeleteEntradaProducto(writer http.ResponseWriter, request *http.Request) {
	entradaProducto := models.Entrada_Producto{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&entradaProducto, id)

	if entradaProducto.ID > 0 {
		db.Delete(entradaProducto)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
