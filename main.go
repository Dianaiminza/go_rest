package main

import (
	"net/http"
	"simple-res-api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})
	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()
	e.Logger.Fatal(e.Start(":8080"))
}
