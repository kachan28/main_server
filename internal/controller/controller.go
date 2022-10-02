package controller

import (
	"main_server/internal/service"

	validator "github.com/go-playground/validator/v10"
)

type Controller struct {
	validate *validator.Validate
	srv      *service.Service
}

func InitController() (*Controller, error) {
	validate := validator.New()

	srv, err := service.InitializeService()
	if err != nil {
		return nil, err
	}

	return &Controller{validate: validate, srv: srv}, nil
}
