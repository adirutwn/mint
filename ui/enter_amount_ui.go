package ui

import (
	"fmt"
	"github.com/adirutwn/mint/validators"
	"github.com/dustin/go-humanize"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"
)

func EnterAmount(assetCode string) string {
	amountPrompt := promptui.Prompt{
		Label: fmt.Sprintf("[💸] please enter the amount of %s you want to mint ", assetCode),
		Validate: validators.Amount,
	}
	amount, _ := amountPrompt.Run()
	amount = strings.TrimSpace(strings.Replace(amount, ",", "", -1))
	paramsAmountFloat, _ := strconv.ParseFloat(amount, 64)
	fmt.Printf("[👌] got it, you want to mint %s %s\n", humanize.Commaf(paramsAmountFloat), assetCode)

	return  amount
}
