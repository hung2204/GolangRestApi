package mdw

import "github.com/labstack/echo/v4"

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "password" {
		c.Set("username", username)
		c.Set("admin", true)
		return true, nil
	}

	if username == "hung2204" && password == "password" {
		c.Set("username", username)
		c.Set("admin", false)
		return false, nil
	}

	return true, nil
}
