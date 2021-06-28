package middlewares

import (
	"app/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GenerateToken(userId int) (string, error) {

	// payload
	mapClaim := jwt.MapClaims{}
	mapClaim["userId"] = userId
	mapClaim["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaim)

	// key
	return token.SignedString([]byte(constants.KEY_JWT))
}

func ExtractUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}
