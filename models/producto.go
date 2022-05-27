package models

type Producto struct {
	ID           int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre       string `json:"nombre"`
	Existencia   int    `json:"existencia"`
	Precio       int    `json:"precio"`
	Status       int    `json:"status"`
	Id_Categoria uint   `json:"id_categoria"`
}
