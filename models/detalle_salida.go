package models

type Detalle_Salida struct {
	ID          int16 `json:"id" gorm:"primary_key;auto_increment"`
	Id_producto int16 `json:"id_producto"`
	Cantidad    int16 `json:"cantidad"`
	Precio      int16 `json:"Precio"`
}
