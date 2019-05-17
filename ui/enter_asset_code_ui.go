package ui

import (
	"fmt"
	"github.com/adirutwn/mint/validators"
	"github.com/manifoldco/promptui"
	"strings"
)

func EnterAssetCode() string {
	assetCodePrompt := promptui.Prompt{
		Label: "[💸] please enter the asset code ",
		Validate: validators.AssetCode,
	}
	assetCode, _ := assetCodePrompt.Run()
	assetCode = strings.TrimSpace(strings.ToUpper(assetCode))
	fmt.Printf("[👍] cool, you're about to mint %s\n", assetCode)

	return assetCode
}
