package ui

import (
	"fmt"
	"github.com/adirutwn/mint/utils"
	"github.com/divan/num2words"
	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"strconv"
)

func DisplaySummary(assetCode string, amount string,
	issuerAddress string, distributorAddress string,
	isDistributorAccountNeedToCreate bool, isDistributorAccountNeedToTrust bool,
	isIssuerAccountLock bool) error {

	paramsAmountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return errors.New("well, I cannot change amount to floating number")
	}

	distributorAccountExtraMsg := ""
	if isDistributorAccountNeedToCreate == true {
		distributorAccountExtraMsg = "(will be created)"
	}

	needToTrustMsg := ""
	if isDistributorAccountNeedToTrust == true {
		needToTrustMsg = fmt.Sprintf("PS. The trust of %s [mint by %s] will be added to the distributor account\n", assetCode, issuerAddress[:5])
	}

	issuerAccountLockMsg := ""
	if isIssuerAccountLock == true {
		needToTrustMsg = fmt.Sprintf("[!!IMPORTANT!!] %s will be locked forever and won't be able to do anything after mint %s\n", issuerAddress[:5], assetCode)
	}

	decimalMsg := ""
	isDecimal := utils.IsDecimal(amount)
	if isDecimal {
		decimalMsg = " + some decimal places"
	}

	fmt.Println()
	fmt.Println("here's a summary:")
	fmt.Println("-------------------------------------")
	fmt.Printf("asset code:          %s\n", assetCode)
	fmt.Printf("asset issuer:        %s\n", issuerAddress)
	fmt.Printf("amount to be mint:   %s (%s%s)\n", humanize.Commaf(paramsAmountFloat), num2words.Convert(int(paramsAmountFloat)), decimalMsg)
	fmt.Printf("distributor address: %s %s\n", distributorAddress, distributorAccountExtraMsg)
	fmt.Printf("%s", needToTrustMsg)
	fmt.Printf("%s", issuerAccountLockMsg)
	fmt.Println("-------------------------------------")
	fmt.Println()

	return nil
}
