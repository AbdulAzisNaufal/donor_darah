package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DeleteStokDarah digunakan untuk menghapus stok darah
func DeleteStokDarah(c echo.Context) error {
	// Ambil ID stok darah dari parameter URL
	stokID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid stok ID"))
	}

	// Hapus data stok darah dari database berdasarkan ID
	if err := config.DB.Where("id = ?", stokID).Delete(&models.Stok{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menghapus stok darah"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Stok darah berhasil dihapus", nil))
}
