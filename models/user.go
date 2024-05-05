package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string
	Email         string
	Password      string
	ID_GolDarah	  uint `gorm:"foreignKey:ID_GolDarah"`
	Tanggal_Lahir time.Time
	Gender        string
	Alamat        string
}

func (u *User) TableName() string {
    return "User"
}