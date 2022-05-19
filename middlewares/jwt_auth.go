package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"nidzamTest.com/helper"
	"nidzamTest.com/model"
	"nidzamTest.com/service"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		urlString := c.Request().URL.String()

		if strings.Contains(urlString, "user/login") == false {

			// client := &http.Client{}

			authHeader := c.Request().Header["Authorization"]
			if authHeader == nil {
				webResponse := model.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
				}
				// hay := make(map[string]exception.BadRequestError)
				// var messages exception.BadRequestError
				// messages.Desc = "hay"
				// messages.DescGlob = "hay"
				// hay["hay"] = messages
				// fmt.Println(webResponse)
				// panic(exception.NewBadRequestError(hay))

				helper.WriteToResponseBody(c, webResponse, webResponse.Code)
				return nil
			}

			token, _ := service.NewJWTService().ValidateToken(authHeader[0])
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println(claims)
				// log.Println("Claim[user_id]: ", claims["user_id"])
				// log.Println("Claim[issuer] :", claims["issuer"])
			} else {
				webResponse := model.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
				}

				helper.WriteToResponseBody(c, webResponse, webResponse.Code)
				return nil
			}
		}
		// req, err := http.NewRequest("POST", "https://api-dev.adapro.tech/master/auth/me", nil)
		// req.Header.Add("Authorization", authHeader[0])
		// if err != nil || authHeader[0] == "" {

		// 	webResponse := model.WebResponse{
		// 		Code:   http.StatusUnauthorized,
		// 		Status: "UNAUTHORIZED",
		// 	}

		// 	helper.WriteToResponseBody(c, webResponse, webResponse.Code)
		// 	return nil
		// }
		// r, err := client.Do(req)
		// helper.PanicIfError(err)
		// defer r.Body.Close()
		// fmt.Println(r.Body)
		// if r.StatusCode == 401 {

		// 	webResponse := model.WebResponse{
		// 		Code:   http.StatusUnauthorized,
		// 		Status: "UNAUTHORIZED",
		// 	}

		// 	helper.WriteToResponseBody(c, webResponse, webResponse.Code)
		// 	return nil
		// }

		return next(c)
	}
}
