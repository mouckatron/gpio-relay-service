// Standard setup for gin, should not need edited
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginRouter *gin.Engine

func setupGinRouter() {
	ginRouter = gin.Default()

	ginRouter.Use(cors.Default())

	ginRouter.GET("/ping", ping)

	addRoutes()
}

func runGinRouter() {

	ginRouter.Run(fmt.Sprintf("%s:%s", conf.Host, conf.Port))
	//gin.SetMode(gin.ReleaseMode)
}

func ping(c *gin.Context) {
	c.Header("X-Ping", "Pong")
	c.JSON(http.StatusOK, "")
}
