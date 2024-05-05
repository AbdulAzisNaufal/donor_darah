package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateStokDarah digunakan untuk menambahkan stok darah baru oleh admin
func CreateStokDarah(c echo.Context) error {
    // Bind request ke objek model StokDarah
    stok := models.Stok{}
    if err := c.Bind(&stok); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    // Simpan stok darah baru ke database
    if err := config.DB.Create(&stok).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menambahkan stok darah"))
    }

    return c.JSON(http.StatusCreated, utils.SuccessResponse("Stok darah berhasil ditambahkan", stok))
}
