package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zanjs/go-echo-rest/app/models"
)

// AllArticles is get all articles
func AllArticles(c echo.Context) error {
	var (
		data        models.ArticlePage
		queryparams models.QueryParams
		err         error
	)

	qps := c.QueryParams()

	limitq := c.QueryParam("limit")
	offsetq := c.QueryParam("offset")
	startTimeq := c.QueryParam("start_time")
	endTime := c.QueryParam("end_time")

	limit, _ := strconv.Atoi(limitq)
	offset, _ := strconv.Atoi(offsetq)
	fmt.Println(qps)
	fmt.Println(limit)
	fmt.Println(offset)

	if limit == 0 {
		limit = 10
	}

	queryparams.Limit = limit
	queryparams.Offset = offset
	queryparams.StartTime = startTimeq
	queryparams.EndTime = endTime

	data, err = models.GetArticles(queryparams)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, data)
}

// get all articles
// func AllArticles2(c echo.Context) error {
// 	var (
// 		articles []models.Article
// 		err      error
// 	)
// 	articles, err = models.GetArticles()
// 	if err != nil {
// 		return c.JSON(http.StatusForbidden, err)
// 	}
// 	return c.JSON(http.StatusOK, articles)
// }

// get one article
func ShowArticle(c echo.Context) error {
	var (
		article models.Article
		err     error
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	article, err = models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, article)
}

// CreateArticle is create article
func CreateArticle(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["id"].(float64))

	article := new(models.Article)

	article.UserID = userID
	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	err := models.CreateArticle(article)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusCreated, article)
}

// UpdateArticle update article
func UpdateArticle(c echo.Context) error {
	// Parse the content
	article := new(models.Article)

	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update article data
	err = m.UpdateArticle(article)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

// DeleteArticle delete article
func DeleteArticle(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteArticle()
	return c.JSON(http.StatusNoContent, err)
}
