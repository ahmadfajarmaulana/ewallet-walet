package interfaces

import (
	"context"
	"ewallet-wallet/internal/models"

	"github.com/gin-gonic/gin"
)

type IWalletRepo interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	UpdateBalance(ctx context.Context, userID int, amount float64) error
	CreateWalletHistory(ctx context.Context, walletHistory *models.WalletTransaction) error
}

type IWalletService interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
}

type IWalletHandler interface {
	Create(*gin.Context)
}
