package exception

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"nidzamTest.com/helper"
	"nidzamTest.com/model"
)

func ErrorHandler(c echo.Context, err interface{}) {

	if badRequestErrorsss(c, err) {
		return
	}
	internalServerError(c, err)
}
func badRequestErrorsss(c echo.Context, err interface{}) bool {

	execption, ok := err.(BadRequestErrors)
	fmt.Println(ok)
	if ok {
		webResponse := model.WebResponse{
			Code:   406,
			Status: "Bad Request",
			Data:   execption.Error,
		}
		helper.WriteToResponseBody(c, webResponse, 200)
		return true
	} else {
		return false
	}
}
func internalServerError(c echo.Context, err interface{}) {
	webResponse := model.WebResponseError{}
	webResponse = model.WebResponseError{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  err,
	}

	helper.WriteToResponseBody(c, webResponse, webResponse.Code)
}
