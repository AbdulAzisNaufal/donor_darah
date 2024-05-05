package controller

import (
	"app/config"
	"app/models"
	"app/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func PesanDarah(c echo.Context) error {
	// Bind request ke objek model Order
    pesanan := models.Order{}
    if err := c.Bind(&pesanan); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    // Set waktu saat pembuatan pesanan
    pesanan.CreatedAt = time.Now()

    // Simpan pesanan darah ke database
    if err := config.DB.Create(&pesanan).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal melakukan pesanan darah"))
    }
	
	// Kembalikan respons dengan data yang diminta
    response := struct {
        ID        uint      `json:"id"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
        ID_Pusat   uint      `json:"id_pusat"`
        ID_GolDarah uint      `json:"id_gol_darah"`
        Jumlah    int32      `json:"jumlah"`
    }{
        ID:        		pesanan.ID,
        CreatedAt: 		pesanan.CreatedAt,
        UpdatedAt: 		pesanan.UpdatedAt,
        ID_Pusat:   	pesanan.ID_Pusat,
        ID_GolDarah:	pesanan.ID_GolDarah,
        Jumlah:    		pesanan.Jumlah,
    }

    return c.JSON(http.StatusOK, utils.SuccessResponse("Pesanan darah berhasil", response))
}
