package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"todoApi/models"
)

func PostTodo(c *gin.Context, db *gorm.DB) {
	var user models.User
	value, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
	}

	email := value.(string)
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = user.ID
	if err := db.Create(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully!"})
}
