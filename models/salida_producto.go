package models

type Salida_Producto struct {
	ID            int64  `json:"id" gorm:"primary_key;auto_increment"`
	Numero_Salida string `json:"numero_salida"`
	Fecha         string `json:"fecha"`
	Descripcion   string `json:"descripcion"`
	Status        string `json:"status"`
}
