package web

import (
	"time"
)

type UserReponse struct {
	Id            int       `json:"id"`
	Name          string    `json:"name" form:"name"`
	Email         string    `json:"email" form:"email"`
	Password      string    `json:"password" form:"password"`
	Tanggal_Lahir time.Time `json:"Tanggal_Lahir" form:"Tanggal_Lahir"`
	Gender        string    `json:"gender" form:"gender"`
	Alamat        string    `json:"alamat" form:"alamat"`
	Gol_Darah     string    `json:"gol_darah" form:"gol_darah"`
}

type UserLoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
