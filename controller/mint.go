package controller

import (
	"fmt"
	"github.com/adirutwn/mint/entities"
	"github.com/adirutwn/mint/repository"
	"github.com/adirutwn/mint/services/stellar"
	"github.com/adirutwn/mint/ui"
	"github.com/stellar/go/clients/horizonclient"
)

func Mint(horizonClient *horizonclient.Client) {
	var params entities.MintParams

	params.IssuerSeedKey = ui.EnterIssuerSeedKey()
	issuerKp, err := stellar.IsAccountCreatedFromSeed(horizonClient, params.IssuerSeedKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	params.IssuerAddress = issuerKp.Address()

	params.AssetCode = ui.EnterAssetCode()
	params.Amount = ui.EnterAmount(params.AssetCode)
	params.DistributorAddress = ui.EnterDistributorAddress(params.AssetCode)
	if params.DistributorAddress == params.IssuerAddress {
		fmt.Println("[🤨] hmmm.. distributor address must not be the same as asset issuer")
		return
	}

	distributorAcc, err := repository.LoadAccount(horizonClient, params.DistributorAddress)
	if err != nil {
		herr, isHorizonError := err.(*horizonclient.Error)
		if !isHorizonError {
			fmt.Println("[😱] some problems with horizon server")
			return
		}
		if herr.Problem.Status == 404 {
			params.IsDistributorAccountNeedToCreate, err = ui.ChoiceCreateDistributorAccount()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			params.IsDistributorAccountNeedToTrust = true
		} else {
			fmt.Println("[😱] some problems with horizon server")
			return
		}
	}

	if params.IsDistributorAccountNeedToCreate == false {
		ok, err := stellar.IsAccountTrustedAsset(*distributorAcc, params.AssetCode, params.IssuerAddress)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if *ok == false {
			params.IsDistributorAccountNeedToTrust = ui.ChoiceTrustAssetOrNo()
			if params.IsDistributorAccountNeedToTrust == false {
				return
			}
		}
	}

	if params.IsDistributorAccountNeedToTrust == true {
		params.DistributorSeedKey = ui.EnterDistributorSeedKey(params.AssetCode)
	}

	params.IsIssuerAccountLock = ui.ChoiceIssuerAccountLockOrNo(params.AssetCode)

	err = ui.DisplaySummary(
		params.AssetCode,
		params.Amount,
		params.IssuerAddress,
		params.DistributorAddress,
		params.IsDistributorAccountNeedToCreate,
		params.IsDistributorAccountNeedToTrust,
		params.IsIssuerAccountLock,
	)

	confirm := ui.ChoiceConfirmOrNo()
	if confirm {
		resultTx, err := stellar.IssueAsset(horizonClient, params)
		if err != nil {
			fmt.Println(err)
			herr, isHorizonError := err.(*horizonclient.Error)
			if !isHorizonError {
				fmt.Println("[😱] some problems with horizon server")
				return
			}
			fmt.Println(herr.ResultString())
		}

		fmt.Println()
		fmt.Printf("[🎉] congratuation! %s %s has been minted to %s.\n", params.Amount, params.AssetCode, params.DistributorAddress[:5])
		fmt.Printf("[🦄] check %s on any Stellar explorer\n", resultTx.Hash)
		fmt.Println()
	} else {
		fmt.Println("[👋] bye bye..")
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}