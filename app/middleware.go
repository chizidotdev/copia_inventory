package app

import (
	"github.com/chizidotdev/copia/dto"
	"github.com/chizidotdev/copia/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) isAuth(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	if reqToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse("Unauthorized"))
		return
	}

	user, err := server.TokenManager.Parse(reqToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err.Error()))
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}

func (server *Server) getUser(ctx *gin.Context) *dto.Claims {
	user := ctx.MustGet("user").(*dto.Claims)

	return user
}