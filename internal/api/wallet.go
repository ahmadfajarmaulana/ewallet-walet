package api

import (
	"ewallet-wallet/constants"
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	WalletService interfaces.IWalletService
}

func (h *WalletHandler) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.Wallet
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to bind json", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailledBadRequest, nil)
		return
	}

	if req.UserID == 0 {
		log.Error("user_id is empty")
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailledBadRequest, nil)
		return
	}

	err := h.WalletService.CreateWallet(c.Request.Context(), &req)
	if err != nil {
		log.Error("failed to create wallet", err)
		helpers.SendResponseHttp(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, req)
}
