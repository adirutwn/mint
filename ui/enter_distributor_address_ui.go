package ui

import (
	"fmt"
	"github.com/adirutwn/mint/validators"
	"github.com/manifoldco/promptui"
	"strings"
)

func EnterDistributorAddress(assetCode string) string {
	distributorAddressPrompt := promptui.Prompt{
		Label: fmt.Sprintf("[💸] please enter the address you want to issue %s to ", assetCode),
		Validate: validators.StellarAddress,
	}
	distributorAddress, _ := distributorAddressPrompt.Run()
	return strings.TrimSpace(distributorAddress)
}
