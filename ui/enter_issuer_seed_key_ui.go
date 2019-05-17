package ui

import (
	"github.com/adirutwn/mint/validators"
	"github.com/manifoldco/promptui"
	"strings"
)

func EnterIssuerSeedKey() string {
	issuerSeedKeyPrompt := promptui.Prompt{
		Label: "[💸] please enter the issuer seed key ",
		Validate: validators.StellarSeedKey,
	}
	issuerSeedKey, _ := issuerSeedKeyPrompt.Run()
	return strings.TrimSpace(issuerSeedKey)
}
