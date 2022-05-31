package models

type Producto struct {
	ID           int16  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre       string `json:"nombre"`
	Existencia   int16  `json:"existencia"`
	Precio       int16  `json:"precio"`
	Status       int    `json:"status"`
	Id_Categoria uint   `json:"id_categoria"`
}
