package config

import (
	"app/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	username := "donor_darah"
	password := "B1common"
	host := "db4free.net"
	port := "3306"
	name := "donor_darah"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username,
		password,
		host,
		port,
		name,
	)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if errDB != nil {
		panic("Failed to Connect Database")
	}

	fmt.Println("Connected to Database")
}

func AutoMigrate() {
    DB.AutoMigrate(&models.User{}, &models.Admin{}, &models.DaftarDonor{},
	&models.Gol_Darah{}, &models.Jadwal{}, &models.Order{}, &models.Pusat{}, &models.Stok{})
}


