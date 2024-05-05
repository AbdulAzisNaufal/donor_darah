package controller

import (
    "app/config"
    "app/models"
	"app/models/web"
    "app/utils"
	"app/middleware"
    "github.com/labstack/echo/v4"
    "net/http"
)

func AdminLogin(c echo.Context) error {
    var loginRequest web.LoginRequest

    if err := c.Bind(&loginRequest); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    var admin models.Admin
    if err := config.DB.Where("email = ?", loginRequest.Email).First(&admin).Error; err != nil {
        return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
    }

    if err := middleware.ComparePassword(admin.Password, loginRequest.Password); err != nil {
        return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
    }

    token := middleware.CreateToken(int(admin.ID), admin.Email, true)

    // Buat respons dengan data yang diminta
    response := web.AdminLoginResponse{
        Email:    admin.Email,
        Token:    token,
    }

    return c.JSON(http.StatusOK, utils.SuccessResponse("Admin login successful", response))
}
