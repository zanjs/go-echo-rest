package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zanjs/go-echo-rest/app/controllers"
	"github.com/zanjs/go-echo-rest/config"
)

var appConfig = config.Config.App
var jwtConfig = config.Config.JWT

type disk struct {
	read  string
	write string
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// var b = map[string]disk{
	// 	"xvda": disk{"5656418", "22438120"},
	// 	"xvdb": disk{"8493386", "1149266272"},
	// }
	// // var b = map[string]disk{
	// // // "xvda": disk{"11", "22"},
	// // // "xvdb": disk{"33", "44"},
	// // }

	// // for key, val := range a {
	// // 	b[key] = disk{val.read, val.write}
	// // 	// b[key].read = val.read
	// // 	// b[key].write = val.write
	// // }
	// // s := make(map[string]disk)

	// a := []disk{}
	// // l := len(a)
	// s := a[0:]
	// var dd disk
	// dd.read = "22"
	// dd.write = "33"
	// s = append(s, dd)

	// for _, val := range b {
	// 	// b[key] = val

	// 	s = append(s, val)
	// }
	// fmt.Println(s)

	// Routes
	e.GET("/", controllers.GetHome)

	e.POST("/user/add", controllers.CreateUser)

	e.GET("/records/jobs", controllers.AllProductWareroom)

	v0 := e.Group("/v0")

	v0.GET("/", controllers.CreateTable)

	v1 := e.Group("/v1")
	v1.POST("/login", controllers.PostLogin)

	v1.Use(middleware.JWT([]byte(jwtConfig.Secret)))

	// Users
	v1.GET("/users", controllers.AllUsers)
	v1.POST("/users", controllers.CreateUser)
	v1.GET("/users/:id", controllers.ShowUser)
	v1.PUT("/users/:id", controllers.UpdateUser)
	v1.DELETE("/users/:id", controllers.DeleteUser)

	// Articles
	v1.GET("/articles", controllers.AllArticles)
	v1.POST("/articles", controllers.CreateArticle)
	v1.GET("/articles/:id", controllers.ShowArticle)
	v1.PUT("/articles/:id", controllers.UpdateArticle)
	v1.DELETE("/articles/:id", controllers.DeleteArticle)

	// Products
	v1.GET("/products", controllers.AllProducts)
	v1.POST("/products", controllers.CreateProduct)
	v1.GET("/products/:id", controllers.ShowProduct)
	v1.PUT("/products/:id", controllers.UpdateProduct)
	v1.DELETE("/products/:id", controllers.DeleteProduct)

	// Wareroom
	v1.GET("/warerooms", controllers.AllWarerooms)
	v1.POST("/warerooms", controllers.CreateWareroom)
	v1.GET("/warerooms/:id", controllers.ShowWareroom)
	v1.PUT("/warerooms/:id", controllers.UpdateWareroom)
	v1.DELETE("/warerooms/:id", controllers.DeleteWareroom)

	// qm 库存销量更新

	v1.GET("/records", controllers.AllRecords)
	v1.GET("/records/excel", controllers.AllProductWareroomRecords)
	v1.DELETE("/records/:id", controllers.DeleteRecord)
	// Server
	if err := e.Start(fmt.Sprintf("%s:%s", appConfig.HttpAddr, appConfig.HttpPort)); err != nil {
		e.Logger.Fatal(err.Error())
	}

}
