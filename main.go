package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techlateef/jwt-Auth-techies/config"
	"github.com/techlateef/jwt-Auth-techies/controllers"
	"github.com/techlateef/jwt-Auth-techies/repository"
	service "github.com/techlateef/jwt-Auth-techies/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()

	authService    service.AuthService        = service.NewAuthService(userRepository)
	authController controllers.AuthController = controllers.NewAuthService(authService, jwtService)
)

func main() {
	defer config.CloseDatabasec(db)
	r := gin.Default()
	authRoute := r.Group("api/auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}

	r.Run()
}
