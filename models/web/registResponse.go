package web

import "time"

type RegistResponse struct {
	Id                int       `json:"id"`
	Name              string    `json:"name" form:"name"`
	ID_Jadwal         uint      `json:"id_jadwal" form:"id_jadwal"`
	Waktu_Pendaftaran time.Time `json:"waktu_pendaftaran" form:"waktu_pendafatran"`
	Status            string    `json:"status" form:"status"`
}