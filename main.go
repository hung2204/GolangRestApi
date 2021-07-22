package main

import (
	"github.com/GolangRestApi/handler"
	"github.com/GolangRestApi/mdw"
	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//conection
	err := orm.RegisterDataBase("default", "mysql", "root:010198@/golangrestapi?charset=utf8")

	if err != nil {
		glog.Fatal("Fail to register database %v", err)
	}
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		glog.Fatal("Fail to run sync, error: %v", err)
	}
}

func main() {
	// Echo instance
	server := echo.New()

	//middleware
	server.Use(middleware.Logger())

	isLogedIn := middleware.JWT([]byte("mysecretkey"))
	isAdmin := mdw.IsAdminMdw
	// Routes
	server.GET("/", handler.Hello, isLogedIn)
	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))
	server.GET("/admin", handler.Hello, isLogedIn, isAdmin)

	groupv2 := server.Group("/v2")
	groupv2.GET("/hello", handler.Hello2)

	groupUser := server.Group("/restapi")
	groupUser.PUT("/create", handler.CreateUser)
	groupUser.GET("/read", handler.ReadUser)
	groupUser.POST("/update", handler.UpdateUser)
	groupUser.DELETE("/delete", handler.DeleteUser)

	// Start server
	server.Logger.Fatal(server.Start(":8888"))
}
