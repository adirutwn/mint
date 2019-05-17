package ui

import (
	"fmt"
	"github.com/adirutwn/mint/constants"
	"github.com/manifoldco/promptui"
)

func ChoiceIssuerAccountLockOrNo(assetCode string) bool {
	issueAccountLockOrNoPrompt := promptui.Select{
		Label:     fmt.Sprintf("[💸] should we lock the issuer account? (by locking it means this issuer won't be able to do anything on Stellar after issued %s", assetCode),
		Items:     constants.YesNoMenuItems,
	}
	selectedIndex, _, _ := issueAccountLockOrNoPrompt.Run()
	if selectedIndex == 0 {
		return true
	} else {
		return false
	}
}
