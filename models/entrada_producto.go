package models

type Entrada_Producto struct {
	ID             int64  `json:"id" gorm:"primary_key;auto_increment"`
	Numero_Entrada string `json:"numero_entrada"`
	Fecha          string `json:"fecha"`
	Descripcion    string `json:"descripcion"`
	Status         string `json:"status"`
}
