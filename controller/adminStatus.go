package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateStatus(c echo.Context) error {
    // Buat sebuah struktur untuk membaca input JSON
    type Input struct {
        ID int `json:"id"`
        Status string `json:"status"`
    }

    // Inisialisasi variabel untuk menyimpan input JSON
    var input Input

    // Bind input JSON ke variabel "input"
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid JSON input"))
    }

    // Sekarang Anda bisa menggunakan "input.ID" untuk mendapatkan ID dari input JSON
    registrationID := input.ID

    // Perbarui status peserta di basis data
    var registration models.DaftarDonor
    if err := config.DB.Where("id = ?", registrationID).First(&registration).Error; err != nil {
        return c.JSON(http.StatusNotFound, utils.ErrorResponse("Registration not found"))
    }

    registration.Status = input.Status
    if err := config.DB.Save(&registration).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update registration status"))
    }

    // Beri respons yang sesuai
    return c.JSON(http.StatusOK, utils.SuccessResponse("Registration status updated successfully", registration))
}

// GetStatus memungkinkan admin untuk mengambil status peserta berdasarkan ID
func GetStatus(c echo.Context) error {
    // Buat sebuah struktur untuk membaca input JSON
    type Input struct {
        ID int `json:"id"`
    }

    // Inisialisasi variabel untuk menyimpan input JSON
    var input Input

    // Bind input JSON ke variabel "input"
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid JSON input"))
    }

    // Sekarang Anda bisa menggunakan "input.ID" untuk mendapatkan ID dari input JSON
    registID := input.ID

    // Ambil status peserta dari basis data berdasarkan ID
    var registration models.DaftarDonor
    if err := config.DB.Where("id = ?", registID).First(&registration).Error; err != nil {
        return c.JSON(http.StatusNotFound, utils.ErrorResponse("Registration not found"))
    }

    // Beri respons yang sesuai
    return c.JSON(http.StatusOK, utils.SuccessResponse("Registration status retrieved successfully", registration))
}

// CreateRegistration memungkinkan pengguna untuk membuat pendaftaran peserta donor darah
func CreateRegistration(c echo.Context) error {
    // Ambil data pendaftaran dari request body
    var registration models.DaftarDonor
    if err := c.Bind(&registration); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    // Simpan data pendaftaran ke basis data
    if err := config.DB.Create(&registration).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create registration"))
    }

    // Beri respons yang sesuai
    return c.JSON(http.StatusCreated, utils.SuccessResponse("Registration created successfully", registration))
}