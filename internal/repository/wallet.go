package repository

import (
	"context"
	"ewallet-wallet/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	return r.DB.Create(wallet).Error
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, userID int, amount float64) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var currenctBalance float64
		err := tx.Raw("SELECT balance FROM wallets WHERE user_id = ? FOR UPDATE", userID).Scan(&currenctBalance).Error
		if err != nil {
			return err
		}

		if currenctBalance+amount < 0 {
			return fmt.Errorf("current balance is not enough to perform this transaction: %f - %f", currenctBalance, amount)
		}

		err = tx.Exec("UPDATE wallets SET balance = balance + ? WHERE user_id = ?", amount, userID).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *WalletRepository) CreateWalletHistory(ctx context.Context, walletHistory *models.WalletTransaction) error {
	return r.DB.Create(walletHistory).Error
}
