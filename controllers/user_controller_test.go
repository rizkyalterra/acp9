package controllers

import (
	"app/configs"
	"app/models/user"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	if Addition(2, 3) != 5 { //  epected : 5
		t.Error("The result actually 5")
	}
	if Addition(6, 4) != 10 {
		t.Error("The result actually 6")
	}
}

func setupEchoDB() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

func AddUserData() bool {
	user := user.User{Name: "Alterra", Passowrd: "123", Email: "alta@gmail.com"}
	err := configs.DB.Create(&user)
	if err != nil {
		return false
	}
	return true
}

func TestGetUserControllers(t *testing.T) {
	e := setupEchoDB()
	// if AddUserData() {
	AddUserData()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/users")

	// Assertions
	if assert.NoError(t, GetUserControllers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseUser user.ResponseUser
		json.Unmarshal([]byte(body), &responseUser)
		// if err != nil {

		assert.Equal(t, responseUser.Status, true)
		assert.Equal(t, len(responseUser.Data), 1)
		assert.Equal(t, responseUser.Data[0].Name, "Alterra")
		// }
		// assert.Equal(t, userJSON, rec.Body.String())
		// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
	}

	// }
}
