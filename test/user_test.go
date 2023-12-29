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

	"github.com/google/uuid"
)

func TestUser(t *testing.T) {
	db := config.Connection()

	user := model.User{
		ID:       uuid.NewString(),
		Name:     "Jordan",
		Email:    "jordan@mail.com",
		Password: "12345678",
	}

	db.Create(&user)
}

func TestUserRepository(t *testing.T) {
	db := config.Connection()
	user := model.User{
		ID:       uuid.NewString(),
		Name:     "Alexia",
		Email:    "alexia@mail.com",
		Password: "12345678",
	}

	r := repositories.NewUserReopsitory(db)
	r.Create(&user)
}

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

func TestUserController(t *testing.T) {
	app := Init()

	responseBody := model.UserRequest{
		ID:       uuid.NewString(),
		Name:     "Suna",
		Email:    "suna@mail.com",
		Password: "12345678",
	}

	body, err := json.Marshal(responseBody)
	fmt.Println("body:", string(body))
	if err != nil {
		fmt.Println("error marshal:", err)
	}
	request := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	if err != nil {
		fmt.Println("error fiber:", err)
	}

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("error read all:", err)
	}

	err = json.Unmarshal(bytes, response)
	if err != nil {
		fmt.Println("error unmarshal:", err)
	}
	fmt.Println("status code:", response.StatusCode)
}
