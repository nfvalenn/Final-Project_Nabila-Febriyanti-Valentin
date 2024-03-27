package middleware

import (
	"MY-GRAM/models"
	util "MY-GRAM/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	authorizationValue := ctx.GetHeader("Authorization")
	splittedValue := strings.Split(authorizationValue, "Bearer ")
	if len(splittedValue) <= 1 {
		var r models.Response = models.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}
	jwtToken := splittedValue[1]

	claims, err := util.VerifyJWT(jwtToken)
	if err != nil {
		var r models.Response = models.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	ctx.Set("claims", claims)

	ctx.Next()
}
