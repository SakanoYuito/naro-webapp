package main

import (
	// "net/http"

	"github.com/labstack/echo/v4"

)

// func connect(c echo.Context) error {
// 	db, _ := model.DB.DB()
// 	defer db.Close()
// 	err := db.Ping()
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, "Connection Failed")
// 	} else {
// 		return c.String(http.StatusOK, "connection established")
// 	}
// }

func main() {
	e := echo.New()
	// e.GET("/", connect)

	e.Logger.Fatal(e.Start(":8000"))
}