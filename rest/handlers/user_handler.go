package handlers

import (
	"fmt"
	"userApp/internal"
	"userApp/storage"
	"userApp/storage/dto"
	"userApp/svc"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserRoutes(apiGroup *gin.RouterGroup, userStorage storage.User) {

	pathByID := fmt.Sprintf("/:%s", internal.UserID)
	userGroup := apiGroup.Group("/users")
	{
		userGroup.POST("", AddUsers(userStorage))
		userGroup.GET(pathByID, GetUserByID(userStorage))
		userGroup.GET("", GetUsers(userStorage))
		userGroup.PUT(pathByID, UpdateUser(userStorage))
		userGroup.DELETE(pathByID, DeleteUser(userStorage))
	}
}

func AddUsers(userStorage storage.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := dto.UserDto{}
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		res, err := svc.InsertUser(c.Request.Context(), userStorage, user)

		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, res)
		return
	}
}

func GetUserByID(userStorage storage.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param(internal.UserID)
		id, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := svc.SelectUserByID(c.Request.Context(), id, userStorage)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, user)
		return
	}
}

func GetUsers(userStorage storage.User) gin.HandlerFunc {
	return func(c *gin.Context) {

		users, err := svc.SelectUsers(c.Request.Context(), userStorage)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, users)
		return
	}
}

func UpdateUser(userStorage storage.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param(internal.UserID)
		id, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		userDto := dto.UserDto{}
		err = c.ShouldBindJSON(&userDto)
		if err != nil {
			c.AbortWithStatusJSON(400, err)
			return
		}

		err = svc.UpdateUser(c.Request.Context(), id, userDto, userStorage)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, "Done")
		return
	}
}

func DeleteUser(userStorage storage.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param(internal.UserID)
		id, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = svc.DeleteUser(c.Request.Context(), id, userStorage)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, "Done")
		return
	}
}
