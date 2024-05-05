package web

type RegistRequest struct {
	Name     		string `json:"name" form:"name"`
	ID_Jadwal       uint   `json:"id_jadwal" form:"id_jadwal"`
}
