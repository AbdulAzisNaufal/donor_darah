package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func RegistToEvent(c echo.Context) error {
	name := c.FormValue("Nama")
	jadwalIDStr := c.FormValue("ID_Jadwal")
	jadwalID, err := strconv.ParseUint(jadwalIDStr, 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid jadwal ID"))
	}

	// Validasi input
	if name == "" || jadwalIDStr == "" {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Harap isi semua field yang diperlukan"))
	}

	// Bind request ke objek model StokDarah
	registration := models.DaftarDonor{}
	if err := c.Bind(&registration); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	// Periksa apakah jadwal dengan ID_Jadwal ada dalam database
	jadwal := models.Jadwal{}
	if err := config.DB.Where("id = ?", jadwalID).First(&jadwal).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Jadwal donor darah tidak ditemukan"))
	}

	// Periksa apakah jadwal penuh
	if jadwal.Kapasitas <= 0 {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Acara donor darah sudah penuh"))
	}

	// Simpan pendaftaran
    registration.Name = name
	registration.ID_Jadwal = uint(jadwalID)
	registration.Waktu_Pendaftaran = time.Now()
	registration.Status = "Pendaftaran Berhasil"

	if err := config.DB.Create(&registration).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mendaftar ke acara donor darah"))
	}

	// Perbarui kapasitas acara donor darah
	// jadwal.Kapasitas--

	// if err := config.DB.Save(&jadwal).Error; err != nil {
	// 	return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal memperbarui kapasitas acara donor darah"))
	// }

	// Kembalikan respons dengan data yang diminta
	response := struct {
		ID                uint      `json:"ID"`
		Name              string    `json:"Nama"`
		ID_Jadwal         uint      `json:"ID_Jadwal"`
		Waktu_Pendaftaran time.Time `json:"Waktu_Pendaftaran"`
		Status            string    `json:"Status"`
	}{
		ID:                registration.ID,
		Name:              registration.Name,
		ID_Jadwal:         registration.ID_Jadwal,
		Waktu_Pendaftaran: registration.Waktu_Pendaftaran,
		Status:            registration.Status,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Pendaftaran berhasil", response))
}
