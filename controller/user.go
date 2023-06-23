package controller

import (
	"naro-webapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	user.Name = c.QueryParam("name")

	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func GetUserById(c echo.Context) error {
	user := model.User{}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user.ID = uint(id)
	model.DB.First(&user, user.ID)
	return c.JSON(http.StatusOK, user)
}

func GetAllUser(c echo.Context) error {
	users := []model.User{}
	model.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}