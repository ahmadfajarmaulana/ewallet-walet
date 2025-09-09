package cmd

import (
	"ewallet-wallet/external"
	"ewallet-wallet/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateToken(ctx *gin.Context) {
	var (
		log = helpers.Logger
	)
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization is empty")
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	tokenData, err := external.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Error(err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", tokenData)
	ctx.Next()
}
