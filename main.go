package main

import (
	"fmt"
	"github.com/adirutwn/mint/constants"
	"github.com/adirutwn/mint/controller"
	"github.com/manifoldco/promptui"
	"github.com/stellar/go/clients/horizonclient"
)

func main() {
	fmt.Println(
	`
		███╗   ███╗██╗███╗   ██╗████████╗
		████╗ ████║██║████╗  ██║╚══██╔══╝
		██╔████╔██║██║██╔██╗ ██║   ██║   
		██║╚██╔╝██║██║██║╚██╗██║   ██║   
		██║ ╚═╝ ██║██║██║ ╚████║   ██║   
		╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝   ╚═╝
	`)

	fmt.Println("Welcome to mint, a cmd tool for minting an asset on Stellar")

	selectNetworkPrompt := promptui.Select{
		Label:     "[✨] select the network",
		Items:     constants.NetworkMenuItems,
	}
	selectedIndex, _, _ := selectNetworkPrompt.Run()
	var horizonClient *horizonclient.Client
	if selectedIndex == 0 {
		horizonClient = horizonclient.DefaultTestNetClient
	} else {
		horizonClient = horizonclient.DefaultPublicNetClient
	}

	controller.Mint(horizonClient)
}
