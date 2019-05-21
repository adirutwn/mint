package ui

import (
	"github.com/adirutwn/mint/validators"
	"github.com/manifoldco/promptui"
	"strings"
)

func EnterHomeDomain() string {
	homeDomainPrompt := promptui.Prompt{
		Label: "[💸] please enter the home domain of issuer account, home domain is a url which we provided metadata of our token (if any) ",
		Validate: validators.Url,
		Default: "",
	}
	homeDomainUri, _ := homeDomainPrompt.Run()
	return strings.TrimSpace(homeDomainUri)
}
