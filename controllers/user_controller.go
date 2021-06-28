package controllers

import (
	"app/configs"
	"app/middlewares"
	"app/models/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Addition(a, b int) int {
	hasil := a + b
	if hasil < 0 {
		hasil = 0
	}
	return hasil
}

func CreateUserControllers(c echo.Context) error {
	var userCreaate user.UserRegister
	c.Bind(&userCreaate)

	var userDB user.User
	userDB.Name = userCreaate.Name
	userDB.Email = userCreaate.Email
	userDB.Passowrd = userCreaate.Passowrd
	userDB.School = userCreaate.School

	err := configs.DB.Create(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user.ResponseUser{
			false, "Failed insert database user", nil,
		})
	}

	token, err := middlewares.GenerateToken(int(userDB.ID))

	userResponse := user.UserResponse{
		userDB.ID, userDB.CreatedAt, userDB.UpdatedAt, userDB.Name, userCreaate.Email, token, userDB.School,
	}

	return c.JSON(http.StatusOK, user.ResponseUserSingle{
		true, "Success insert database user", userResponse,
	})
}

func GetUserControllers(c echo.Context) error {

	var dataUsers []user.User
	err := configs.DB.Preload("School").Find(&dataUsers).Error

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, user.ResponseUser{
			false, "Failed get database user", nil,
		})
	}

	return c.JSON(http.StatusOK, user.ResponseUser{
		true, "Success get Data User", dataUsers,
	})
}
