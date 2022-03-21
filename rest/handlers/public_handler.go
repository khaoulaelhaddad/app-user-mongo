package handlers

import "github.com/gin-gonic/gin"

func PublicRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/health", Health())
}

func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		health := gin.H{
			"server":      "ok",
			"grpc_server": "ok",
		}

		c.JSON(200, health)
	}
}
