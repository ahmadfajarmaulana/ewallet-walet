package external

import (
	"context"
	"ewallet-wallet/constants"
	"ewallet-wallet/external/proto/tokenvalidation"
	"ewallet-wallet/internal/models"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func ValidateToken(ctx context.Context, token string) (models.TokenData, error) {
	var (
		res models.TokenData
	)

	conn, err := grpc.Dial("localhost:7000", grpc.WithInsecure())
	if err != nil {
		return res, errors.Wrap(err, "failed to dial ums grpc service")
	}
	defer conn.Close()

	client := tokenvalidation.NewTokenValidationClient(conn)
	req := &tokenvalidation.TokenRequest{
		Token: token,
	}

	response, err := client.ValidateToken(ctx, req)
	if err != nil {
		return res, errors.Wrap(err, "failed to validate token")
	}

	if response.Message != constants.SuccessMessage {
		return res, fmt.Errorf("got response error from ums: %s", response.Message)
	}

	res.UserID = response.Data.UserId
	res.Username = response.Data.Username
	res.FullName = response.Data.FullName

	return res, nil
}
