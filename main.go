package main

import (
	"ewallet-wallet/cmd"
	"ewallet-wallet/helpers"
)

func main() {
	helpers.SetupConfig()

	helpers.SetupLogger()

	helpers.SetupMySQL()

	cmd.ServeHTTP()
	// go cmd.ServeGRPC()
}
