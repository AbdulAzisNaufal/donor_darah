package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

//untuk mengambil data stok darah
func GetStok(c echo.Context) error {
    // Mengambil semua stok darah dari database
    var stok []models.Stok
    if err := config.DB.Find(&stok).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal mengambil stok darah"))
    }

    return c.JSON(http.StatusOK, utils.SuccessResponse("Data stok darah", stok))
}
