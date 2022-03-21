package handlers

import "github.com/gin-gonic/gin"

func PublicRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/ping", Ping())
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
