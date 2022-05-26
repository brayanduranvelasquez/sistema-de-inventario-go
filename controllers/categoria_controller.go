package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brayanduranvelasquez/sistema-de-inventario-go/commons"
	"github.com/brayanduranvelasquez/sistema-de-inventario-go/models"
	"github.com/gorilla/mux"
)

func GetTodasCategorias(writer http.ResponseWriter, request *http.Request) {
	categoria := []models.Categoria{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&categoria)
	json, _ := json.Marshal(categoria)
	commons.SendReponse(writer, http.StatusOK, json)
}

func GetCategoria(writer http.ResponseWriter, request *http.Request) {
	categoria := models.Categoria{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&categoria, id)

	if categoria.ID > 0 {
		json, _ := json.Marshal(categoria)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func SaveCategoria(writer http.ResponseWriter, request *http.Request) {
	categoria := models.Categoria{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&categoria)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&categoria).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(categoria)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func DeleteCategoria(writer http.ResponseWriter, request *http.Request) {
	categoria := models.Categoria{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&categoria, id)

	if categoria.ID > 0 {
		db.Delete(categoria)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
