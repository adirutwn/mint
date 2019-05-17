package ui

import (
	"github.com/adirutwn/mint/validators"
	"github.com/manifoldco/promptui"
	"strings"
)

func EnterDistributorSeedKey(assetCode string) string {
	distributorAddressPrompt := promptui.Prompt{
		Label: "[💸] the seed key of a distributor is needed ",
		Validate: validators.StellarSeedKey,
	}
	distributorAddress, _ := distributorAddressPrompt.Run()
	return strings.TrimSpace(distributorAddress)
}
