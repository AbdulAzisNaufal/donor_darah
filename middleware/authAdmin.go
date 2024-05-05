package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthorizationAdmin(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Dapatkan token dari header Authorization
        tokenString := c.Request().Header.Get("Authorization")

        if tokenString == "" {
            return c.JSON(http.StatusUnauthorized, map[string]interface{}{
                "message": "Token is missing",
            })
        }

        // Lakukan verifikasi token
        token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte("1234"), nil // Ganti dengan secret key yang sebenarnya
        })

        if err != nil {
            return c.JSON(http.StatusUnauthorized, map[string]interface{}{
                "message": "Invalid token",
            })
        }

        // Cek apakah token valid
        if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
            // Token valid, Anda dapat mengakses klaim-klaim di sini, seperti claims.Admin

            // Misalnya, jika Anda ingin memeriksa apakah pengguna adalah admin
            if claims.Admin {
                // Pengguna adalah admin, izinkan eksekusi rute berikutnya
                return next(c)
            } else {
                // Pengguna bukan admin, kembalikan kesalahan atau respons sesuai
                return c.JSON(http.StatusForbidden, map[string]interface{}{
                    "message": "Permission denied",
                })
            }
        }

        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Invalid token",
        })
    }
}