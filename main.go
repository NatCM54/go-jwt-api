package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	AuthController "naatceeam/jwt_api/controller/auth"
	UserController "naatceeam/jwt_api/controller/user"
	"naatceeam/jwt_api/middleware"
	"naatceeam/jwt_api/orm"
)

// // Binding from JSON
// type Register struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// 	Fullname string `json:"fullname" binding:"required"`
// 	Avatar   string `json:"avatar" binding:"required"`
// }

// type UserLogin struct {
// 	gorm.Model
// 	Username string
// 	Password string
// 	Fullname string
// 	Avatar   string
//}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", UserController.ReadAll)
	authorized.GET("/profile", UserController.Profile)
	r.Run("localhost:8080")
}
