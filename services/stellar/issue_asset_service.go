package stellar

import (
	"fmt"
	"github.com/adirutwn/mint/entities"
	"github.com/adirutwn/mint/utils"
	"github.com/briandowns/spinner"
	"github.com/pkg/errors"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/network"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/txnbuild"
	"time"
)

func IssueAsset(hClient *horizonclient.Client, params entities.MintParams) (*horizon.TransactionSuccess, error) {
	sp := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	sp.Start()

	var operations []txnbuild.Operation

	issuerKp, err := utils.KpFromSeedString(params.IssuerSeedKey)
	if err != nil {
		sp.Stop()
		return nil, err
	}

	if params.IsDistributorAccountNeedToCreate {
		createAccountOp := txnbuild.CreateAccount{
			Destination: params.DistributorAddress,
			Amount: "1",
		}
		operations = append(operations, &createAccountOp)
	}

	if params.IsDistributorAccountNeedToTrust {
		paymentOp := txnbuild.Payment{
			Destination: params.DistributorAddress,
			Amount: "0.5",
			Asset: txnbuild.NativeAsset{},
		}
		operations = append(operations, &paymentOp)

		changeTrustOp := txnbuild.ChangeTrust{
			Line: txnbuild.CreditAsset{
				Code: params.AssetCode,
				Issuer: params.IssuerAddress,
			},
			Limit: txnbuild.MaxTrustlineLimit,
			SourceAccount: &horizon.Account {
				HistoryAccount: horizon.HistoryAccount {
					AccountID: params.DistributorAddress,
				},
			},
		}
		operations = append(operations, &changeTrustOp)
	}

	issueAssetOp := txnbuild.Payment{
		Destination: params.DistributorAddress,
		Amount: params.Amount,
		Asset: txnbuild.CreditAsset{
			Code: params.AssetCode,
			Issuer: params.IssuerAddress,
		},
	}
	operations = append(operations, &issueAssetOp)

	if params.IsIssuerAccountLock {
		highestThreshold := txnbuild.Threshold(255)
		lowestThreshold := txnbuild.Threshold(0)

		lockAccountOp := txnbuild.SetOptions{
			MasterWeight: &lowestThreshold,
			LowThreshold: &highestThreshold,
			MediumThreshold: &highestThreshold,
			HighThreshold: &highestThreshold,
		}

		operations = append(operations, &lockAccountOp)
	}

	issuerAccount, err := hClient.AccountDetail(horizonclient.AccountRequest{ AccountID: params.IssuerAddress })
	if err != nil {
		sp.Stop()
		return nil, errors.New("[😱] some problems with horizon server")
	}

	networkPassphrase := ""
	if hClient.HorizonURL == horizonclient.DefaultPublicNetClient.HorizonURL {
		networkPassphrase = network.PublicNetworkPassphrase
	} else {
		networkPassphrase = network.TestNetworkPassphrase
	}

	tx := txnbuild.Transaction{
		SourceAccount: &issuerAccount,
		Operations: operations,
		Timebounds: txnbuild.NewTimeout(300),
		Network: networkPassphrase,
	}

	txeB64 := ""
	if params.IsDistributorAccountNeedToTrust || params.IsDistributorAccountNeedToCreate {
		distributorKp, err := utils.KpFromSeedString(params.DistributorSeedKey)
		if err != nil {
			sp.Stop()
			return nil, err
		}

		txeB64, err = tx.BuildSignEncode(issuerKp, distributorKp)
		if err != nil {
			sp.Stop()
			return nil, errors.New("[😱] unable to sign the tx")
		}
	} else {
		txeB64, err = tx.BuildSignEncode(issuerKp)
		if err != nil {
			sp.Stop()
			return nil, errors.New("[😱] unable to sign the tx")
		}
	}

	resp, err := hClient.SubmitTransactionXDR(txeB64)
	if err != nil {
		herr, isHorizonError := err.(*horizonclient.Error)
		if !isHorizonError {
			sp.Stop()
			return nil, errors.New("[😱] some problems with horizon server")
		}
		herrString, _ := herr.ResultString()
		sp.Stop()
		return nil, errors.New(fmt.Sprintf("[😱] failed to confirm tx with reason: %s", herrString))
	}
	sp.Stop()

	return &resp, nil
}
