package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string `gorm:"unique" json:"email"`   
	Password  string 
}

func (u *Admin) TableName() string {
    return "Admin"
}