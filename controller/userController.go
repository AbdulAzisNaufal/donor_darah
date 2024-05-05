package controller

import (
	"app/config"
	"app/middleware"
	"app/models"
	"app/models/web"
	"app/utils"
	"app/utils/req"
	"app/utils/res"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndex(users)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", response))
}

func Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	response := res.ConvertGeneral(&user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully retrieved", response))
}

//Registrasi User
func Store(c echo.Context) error {
	var user web.UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	userDb := req.PassBody(user)

	// Hash the user's password before storing it
	userDb.Password = middleware.HashPassword(userDb.Password)

	var golDarah models.Gol_Darah
	if err := config.DB.Where("Gol_Darah = ?", user.Gol_Darah).First(&golDarah).Error; err != nil {
    // Handle kesalahan jika golongan darah tidak ditemukan
    return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid blood type"))
}

	//data 
	userDb.Name = user.Name
	userDb.Email = user.Email
	userDb.Tanggal_Lahir = user.Tanggal_Lahir
	userDb.Gender = user.Gender
	userDb.Alamat = user.Alamat
	userDb.ID_GolDarah = golDarah.ID

	if err := config.DB.Create(&userDb).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store user data"))
	}

	// Return the response without including a JWT token
	response := res.ConvertGeneral(userDb)

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", response))
}


//Login User
func Login(c echo.Context) error {
	var loginRequest web.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var user models.User
	if err := config.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	token := middleware.CreateToken(int(user.ID), user.Name, false)

	// Buat respons dengan data yang diminta
	response := web.UserLoginResponse{
		Email:    user.Email,
		Password: loginRequest.Password,
		Token:    token,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", response))
}

func Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var updatedUser models.User

	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingUser models.User
	result := config.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	config.DB.Model(&existingUser).Updates(updatedUser)

	response := res.ConvertGeneral(&existingUser)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully updated", response))
}

func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var existingUser models.User
	result := config.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve user"))
	}

	config.DB.Delete(&existingUser)

	return c.JSON(http.StatusOK, utils.SuccessResponse("User data successfully deleted", nil))
}
