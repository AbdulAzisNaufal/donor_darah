package models

import (
	"gorm.io/gorm"
	"time"
)

type Pusat struct {
	ID        		uint `gorm:"primarykey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Alamat 			string 
	No_telp 		string 
	Jam_Operasional	time.Time
}

func (u *Pusat) TableName() string {
    return "Pusat"
}
