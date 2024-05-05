package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        		uint `gorm:"primarykey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	ID_Pusat 		uint `gorm:"foreignKey:ID_Pusat"`
	ID_User 		uint `gorm:"foreignKey:ID_User"`
	ID_GolDarah 	uint `gorm:"foreignKey:ID_GolDarah"`
	Jumlah 			int32
}

func (u *Order) TableName() string {
    return "Order"
}