package stellar

import (
	"github.com/adirutwn/mint/repositories"
	"github.com/adirutwn/mint/utils"
	"github.com/pkg/errors"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
)

func IsAccountCreatedFromSeed(horizonClient *horizonclient.Client, seedKey string) (*keypair.Full, error) {
	issuerKp, err := utils.KpFromSeedString(seedKey)
	if err != nil {
		return nil, err
	}

	_, err = repositories.LoadAccount(horizonClient, issuerKp.Address())
	if err != nil {
		herr, isHorizonError := err.(*horizonclient.Error)
		if !isHorizonError {
			return nil, errors.New("[😱] some problems with horizon server")
		}
		if herr.Problem.Status == 404 {
			return nil, errors.New("[😱] issuer account is not created yet")
		}
	}

	return issuerKp, nil
}