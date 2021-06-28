package middlewares

import (
	"app/configs"
	"app/models/user"

	"github.com/labstack/echo"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user user.User
	if err := configs.DB.Where("email = ? AND passowrd = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
