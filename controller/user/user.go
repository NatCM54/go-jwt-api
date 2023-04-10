package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"naatceeam/jwt_api/orm"
)

func ReadAll(c *gin.Context) {
	var users []orm.UserLogin
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.UserLogin
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "user": user})
}
