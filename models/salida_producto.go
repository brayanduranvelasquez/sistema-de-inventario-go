package models

import (
	"github.com/jinzhu/gorm"
)

type Salida_Producto struct {
	gorm.Model

	Numero_Salida int16  `json:"numero_salida"`
	Descripcion   string `json:"descripcion"`
	Status        int    `json:"status"`
}
