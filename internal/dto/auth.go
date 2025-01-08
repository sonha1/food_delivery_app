package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct{}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token      string            `json:"token"`
	UserDetail UserLoginResponse `json:"userDetail"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

type RegisterResponse struct{}

func (r *LoginRequest) Validate() error {
	rules := map[string]string{
		"username": "required,min=4,max=32",
		"password": "required,min=6",
	}

	validate := validator.New()
	validate.RegisterStructValidationMapRules(rules, LoginRequest{})
	err := validate.Struct(r)
	if err != nil {

		return fmt.Errorf("Validate Login Request error: ", err)
	}

	return nil
}
