package ui

import (
	"github.com/adirutwn/mint/constants"
	"github.com/manifoldco/promptui"
)

func ChoiceConfirmOrNo() bool {
	confirmOrNoPrompt := promptui.Select{
		Label:     "[💸] confirm?",
		Items:     constants.YesNoMenuItems,
	}
	selectedIndex, _, _ := confirmOrNoPrompt.Run()
	if selectedIndex == 0 {
		return true
	} else {
		return false
	}
}
