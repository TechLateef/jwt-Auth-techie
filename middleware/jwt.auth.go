package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	service "github.com/techlateef/jwt-Auth-techies/services"
)

// AuthorizeJWT validate the token user give, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "failed to process"})
			return
		}
		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			{

				log.Println("Claim[user_id]: ", claims["user_id"])
				log.Println("Claim[issuer]: ", claims["issuer"])

			}
			role := fmt.Sprintf("%v", claims["role"])
			if role == "admin" {
				log.Println("Admin User : Access granted")
			} else {
				res := "Non Admin User : Access denied"
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			}

		} else {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
	}
}
