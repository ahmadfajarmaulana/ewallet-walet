package cmd

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()
	r := gin.Default()

	r.GET("/healt", d.HealtCheckApi.HealtcheckHandlerHTTP)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletHandler.Create)

	err := r.Run(":" + helpers.GetEnv("PORT", "8081"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealtCheckApi interfaces.IHealthcheckHandler

	WalletHandler interfaces.IWalletHandler
}

func dependencyInject() Dependency {
	healtcheckSvc := &services.Healcheck{}
	healcheckAPI := &api.Healthcheck{HealthcheckServices: healtcheckSvc}

	walletRepo := &repository.WalletRepository{DB: helpers.DB}
	walletService := &services.WalletService{WalletRepo: walletRepo}
	walletHandler := &api.WalletHandler{WalletService: walletService}

	return Dependency{
		HealtCheckApi: healcheckAPI,
		WalletHandler: walletHandler,
	}
}
