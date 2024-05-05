package models

import (
	"gorm.io/gorm"
	"time"
)

type Gol_Darah struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Gol_Darah string
}

func (u *Gol_Darah) TableName() string {
    return "Gol_Darah"
}
