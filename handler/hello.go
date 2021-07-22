package handler

import (
	"fmt"
	"net/http"

	"github.com/GolangRestApi/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username := claims["username"].(string)
	admin := claims["admin"].(bool)

	message := fmt.Sprintf("Hello %s is admin %v", username, admin)

	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}

func Hello2(c echo.Context) error {
	return c.String(http.StatusOK, "done keo")
}
