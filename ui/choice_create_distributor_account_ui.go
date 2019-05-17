package ui

import (
	"github.com/adirutwn/mint/constants"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
)

func ChoiceCreateDistributorAccount() (bool, error) {
	createDistributorAccountOrNot := promptui.Select{
		Label:     "[💸] the distributor account hasn't been created, do you want to create it?",
		Items:     constants.YesNoMenuItems,
	}
	selectedIndex, _, _ := createDistributorAccountOrNot.Run()
	if selectedIndex == 0 {
		return true, nil
	} else {
		return false, errors.New("[🤨] the distributor account must be created first")
	}
}
