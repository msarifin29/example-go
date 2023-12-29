package test

import (
	"encoding/json"
	"fmt"
	"gorm-example/config"
	"gorm-example/model"
	"gorm-example/repositories"
	"gorm-example/services"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	v := config.NewValidator()
	db := config.Connection()
	user := model.User{
		ID:       uuid.NewString(),
		Name:     "Aurora",
		Email:    "aurora@mail.com",
		Password: "12345678",
	}
	r := repositories.NewUserReopsitory(db)
	r.Create(&user)
	service := services.NewUserService(*r, v, db)
	fmt.Println("service :", service)

}

func TestCreateUserSuccess(t *testing.T) {
	app := Init()

	requestBody := model.UserRequest{
		ID:       uuid.NewString(),
		Name:     "Angela",
		Email:    "angela@mail.com",
		Password: "12345678",
	}

	body, err := json.Marshal(requestBody)
	assert.Nil(t, err)
	request := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(request.Body)
	assert.Nil(t, err)

	err = json.Unmarshal(bytes, response)
	assert.Nil(t, err)

	assert.Equal(t, "Angela", requestBody.Name)
	assert.Equal(t, "angela@mail.com", requestBody.Email)
	assert.Equal(t, "12345678", requestBody.Password)

}
func TestCreateUserFailed(t *testing.T) {
	app := Init()

	requestBody := model.UserRequest{
		ID:       uuid.NewString(),
		Name:     "",
		Email:    "alexa@mail.com",
		Password: "12345678",
	}

	body, err := json.Marshal(requestBody)
	assert.Nil(t, err)
	request := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 500, response.StatusCode)

}

func TestFindAllSuccess(t *testing.T) {
	app := Init()
	CreateUser()
	CreateUser()
	request := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	users := responseBody["data"].([]interface{})

	user1 := users[0].(map[string]interface{})
	user2 := users[0].(map[string]interface{})
	assert.NotNil(t, user1, "user name not null")
	assert.NotNil(t, user2, "user name not null")
}
func TestFindAllFailed(t *testing.T) {
	app := Init()

	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 404, response.StatusCode)
}
