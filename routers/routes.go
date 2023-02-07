package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"todoApi/authentication"
)

func Routes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	//router.GET("/users", getAlbums)
	router.POST("/users", authentication.AuthMiddleware, func(c *gin.Context) {
		PostUsers(c, db)
	})
	router.POST("/create-todo", authentication.AuthMiddleware, func(c *gin.Context) {
		PostTodo(c, db)
	})
	router.GET("/users", authentication.AuthMiddleware, func(c *gin.Context) {
		GetUsers(c, db)
	})
	router.GET("/users/:id", authentication.AuthMiddleware, func(c *gin.Context) {
		GetUser(c, db)
	})
	router.POST("/login", func(c *gin.Context) {
		authentication.Login(c, db)
	})

	return router
}
