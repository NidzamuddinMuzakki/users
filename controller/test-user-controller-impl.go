package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"nidzamTest.com/helper"
	"nidzamTest.com/model"
	"nidzamTest.com/service"
)

type UserControllerImpl struct {
	UserService service.UserService
	JwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &UserControllerImpl{
		UserService: userService,
		JwtService:  jwtService,
	}
}

func (service *UserControllerImpl) Login(ctx echo.Context) {
	loginDTO := LoginDTO{}
	helper.ReadFromRequestBody(ctx, &loginDTO)

	authResult, Role := service.UserService.VerifyCredential(ctx.Request().Context(), loginDTO.Username, loginDTO.Password)
	if authResult == true {
		fmt.Println(loginDTO.Username, Role)
		generatedToken, refresh_token := service.JwtService.GenerateToken(loginDTO.Username, Role)
		DataLogin := LoginResponse{}
		DataLogin.Token = generatedToken
		DataLogin.RefreshToken = refresh_token
		webResponse := model.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   DataLogin,
		}

		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)

	} else {
		webResponse := model.WebResponse{
			Code:   401,
			Status: "OK",
			Data:   "gagal",
		}

		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)

	}
}

func (controller *UserControllerImpl) FindAll(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)
	fmt.Println(auth)
	getall := model.ReqList{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)

	resultData := controller.UserService.FindAll(ctx.Request().Context(), getall.Page, getall.Perpage, getall.Filter, getall.Order)
	// fmt.Println(beli)
	webResponse := model.WebResponseListAndDetail{
		Code: 200,
		Data: resultData,
		Info: "",
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}

func (controller *UserControllerImpl) FindByUsername(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)
	fmt.Println(auth)
	getall := model.ReqListByUsername{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)

	resultData := controller.UserService.FindByUsername(ctx.Request().Context(), getall.Username)
	// fmt.Println(beli)
	webResponse := model.WebResponseListAndDetail{
		Code: 200,
		Data: resultData,
		Info: "",
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
func (controller *UserControllerImpl) Insert(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)

	CreateRequest := model.UserEntity{}

	helper.ReadFromRequestBody(ctx, &CreateRequest)
	resultData := controller.UserService.Insert(ctx.Request().Context(), CreateRequest, auth.Username, auth.Role)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
func (controller *UserControllerImpl) Update(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)

	CreateRequest := model.UserEntity{}

	helper.ReadFromRequestBody(ctx, &CreateRequest)
	resultData := controller.UserService.Update(ctx.Request().Context(), CreateRequest, auth.Username, auth.Role)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}

func (controller *UserControllerImpl) Delete(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)
	getall := model.ReqListByUsername{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)
	resultData := controller.UserService.Delete(ctx.Request().Context(), getall.Username, auth.Role)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
