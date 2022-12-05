package controllers

import (
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techlateef/jwt-Auth-techies/dto"
	entity "github.com/techlateef/jwt-Auth-techies/entities"
	service "github.com/techlateef/jwt-Auth-techies/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthService(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDTO
	errDto := ctx.ShouldBind(&loginDto)
	if errDto != nil {
		ctx.JSON(http.StatusBadRequest, errDto)
		return
	}
	authResult := c.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GeneratedToken(strconv.FormatUint(v.ID, 10), v.Role)
		fmt.Println(v.Role)
		v.Token = generateToken
		ctx.JSON(http.StatusOK, v)
		return
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDto dto.RegisterDTO
	errDto := ctx.ShouldBind(&registerDto)
	if errDto != nil {
		ctx.JSON(http.StatusBadRequest, errDto)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDto.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email already Exist "})

	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.MinCost)
		if err != nil {
			log.Fatalf("failed to hash password %v ", err)
		}
		registerDto.Password = string(hash)

		createUser := c.authService.CreateUser(registerDto)

		token := c.jwtService.GeneratedToken(createUser.Role, strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token
		ctx.JSON(http.StatusOK, createUser)

	}
}

func (c *authController) GetUsers(ctx *gin.Context) {
	var users []entity.Users = c.authService.GetAllUsers()

	ctx.JSON(http.StatusOK, users)

}
