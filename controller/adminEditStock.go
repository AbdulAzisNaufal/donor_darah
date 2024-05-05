package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UpdateStokDarah digunakan untuk mengedit stok darah
func UpdateStokDarah(c echo.Context) error {
    // Ambil ID stok darah yang akan diperbarui dari parameter URL
	stokID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid stok ID"))
	}
	
	// Bind request ke objek model StokDarah
    stok := models.Stok{}
    if err := c.Bind(&stok); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    // Periksa apakah stok darah dengan ID yang diberikan ada dalam database
	existingStok := models.Stok{}
	if err := config.DB.First(&existingStok, stokID).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Stok darah dengan ID tersebut tidak ditemukan"))
	}

	// Update stok darah yang ada dengan data yang baru
	if err := config.DB.Model(&existingStok).Updates(&stok).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal memperbarui stok darah"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Stok darah berhasil diperbarui", existingStok))
}

