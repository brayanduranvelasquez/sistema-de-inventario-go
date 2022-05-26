package models

type Detalle_Salida struct {
	ID          int64  `json:"id" gorm:"primary_key;auto_increment"`
	Id_producto string `json:"Id_producto"`
	Cantidad    string `json:"cantidad"`
	Precio      string `json:"Precio"`
}
