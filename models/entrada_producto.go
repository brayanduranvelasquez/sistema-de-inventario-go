package models

import (
	"github.com/jinzhu/gorm"
)

type Entrada_Producto struct {
	gorm.Model

	Numero_Entrada int16  `json:"numero_entrada"`
	Descripcion    string `json:"descripcion"`
	Status         int    `json:"status"`
}
