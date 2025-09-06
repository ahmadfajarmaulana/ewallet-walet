package interfaces

import (
	"context"
	"ewallet-wallet/internal/models"

	"github.com/gin-gonic/gin"
)

type IWalletRepo interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
}

type IWalletService interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
}

type IWalletHandler interface {
	Create(*gin.Context)
}
