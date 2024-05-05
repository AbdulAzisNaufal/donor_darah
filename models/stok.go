package models

import (
	"gorm.io/gorm"
	"time"
)

type Stok struct {
	ID        		uint `gorm:"primarykey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	ID_Pusat 		uint `gorm:"foreignKey:ID_Pusat"`
	ID_GolDarah		uint `gorm:"foreignKey:ID_GolDarah"`
	Jumlah			int64
	Expired			time.Time
}

func (u *Stok) TableName() string {
    return "Stok"
}