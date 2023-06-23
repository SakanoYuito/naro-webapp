package main

import (
	// "net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"

	"naro-webapp/controller"
	"naro-webapp/model"
)


func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	db, _ := model.DB.DB()
	defer db.Close()

	e.GET("/users/:id", controller.GetUserById)
	e.GET("/users", controller.GetAllUser)
	e.POST("/users", controller.CreateUser)

	e.POST("/posts/:id", controller.CreatePost)
	e.GET("/posts/:id", controller.GetPostByUserId)
	e.GET("/posts", controller.GetAllPosts)

	e.Logger.Fatal(e.Start(":8000"))
}