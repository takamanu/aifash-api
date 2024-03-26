package main

import (
	"aifash-api/configs"
	"aifash-api/helper"
	email "aifash-api/helper/email"
	encrypt "aifash-api/helper/encrypt"
	"aifash-api/routes"
	"aifash-api/utils/database"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	dataUser "aifash-api/features/users/data"
	handlerUser "aifash-api/features/users/handler"
	serviceUser "aifash-api/features/users/service"
)

func main() {
	e := echo.New()

	var config = configs.InitConfig()

	db, err := database.InitDB(*config)
	if err != nil {
		e.Logger.Fatal("Cannot run database: ", err.Error())
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Endpoint not found", nil))
	})

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Endpoint not found", nil))
	})

	e.GET("/api/v1", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Endpoint not found", nil))
	})

	var encrypt = encrypt.New()
	var email = email.New(*config)

	jwtInterface := helper.New(config.Secret, config.RefSecret)

	userModel := dataUser.New(db)

	userServices := serviceUser.New(userModel, jwtInterface, email, encrypt)

	userController := handlerUser.NewHandler(userServices, jwtInterface)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	group := e.Group("/api/v1")

	routes.RouteUser(group, userController, *config)

	e.Logger.Debug(db)

	e.Logger.Info(fmt.Sprintf("Listening in port :%d", config.ServerPort))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
