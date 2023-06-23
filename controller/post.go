package controller

import (
	"naro-webapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreatePost(c echo.Context) error {
	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return err
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post.UserID = uint(id)

	model.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}

func GetPostByUserId(c echo.Context) error {
	posts := []model.Post{}
	model.DB.Where("user_id = ?", c.Param("id")).Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func GetAllPosts(c echo.Context) error {
	posts := []model.Post{}
	model.DB.Order("created_at DESC").Find(&posts)
	return c.JSON(http.StatusOK, posts)
}