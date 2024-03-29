package commons

import (
	"log"

	"github.com/brayanduranvelasquez/sistema-de-inventario-go/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/sistema-de-inventario?charset=utf8&parseTime=true&loc=Local")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Haciendo migracion....")

	db.AutoMigrate(&models.Categoria{}, &models.Producto{}, &models.Entrada_Producto{}, &models.Salida_Producto{}, &models.Detalle_Entrada{}, &models.Detalle_Salida{})
}
