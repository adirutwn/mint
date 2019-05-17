package repository

import (
	"github.com/briandowns/spinner"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/protocols/horizon"
	"time"
)

func LoadAccount(client *horizonclient.Client, accountId string) (*horizon.Account, error){
	sp := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	sp.Start()
	acc, err := client.AccountDetail(horizonclient.AccountRequest{AccountID: accountId})
	if err != nil {
		sp.Stop()
		return nil, err
	}
	sp.Stop()

	return &acc, err
}
