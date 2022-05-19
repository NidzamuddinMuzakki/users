package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	FindAll(ctx echo.Context)
	FindByUsername(ctx echo.Context)
	Insert(ctx echo.Context)
	Update(ctx echo.Context)
	Delete(ctx echo.Context)
	Login(ctx echo.Context)
}
type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
