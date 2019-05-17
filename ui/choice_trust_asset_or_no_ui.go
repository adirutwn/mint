package ui

import (
	"fmt"
	"github.com/adirutwn/mint/constants"
	"github.com/manifoldco/promptui"
)

func ChoiceTrustAssetOrNo() bool {
	trustAssetOrNo := promptui.Select{
		Label:     "[💸] the distributor account hasn't been trusted the asset yet, do you want to trust it?",
		Items:     constants.YesNoMenuItems,
	}
	selectedIndex, _, _ := trustAssetOrNo.Run()
	if selectedIndex == 0 {
		return true
	} else {
		fmt.Println("[🤨] the distributor account need to trust the asset first")
		return false
	}
}
