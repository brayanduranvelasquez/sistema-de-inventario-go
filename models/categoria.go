package models

type Categoria struct {
	ID     int16  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre string `json:"nombre"`
	Status int    `json:"status"`
}
