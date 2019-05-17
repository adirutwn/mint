package stellar_test

import (
	"github.com/adirutwn/mint/entities"
	"github.com/adirutwn/mint/services/stellar"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssueAsset(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		hc := horizonclient.DefaultTestNetClient
		params := entities.MintParams{
			AssetCode: "HELLO",
			Amount: "1",
			IssuerAddress: "GCWCFW2DC2IU2AQVHQDME34TIHBU4E2X66BWWMEKWECGUYPBCJPSIHLG",
			IssuerSeedKey: "SC5SZJMQ277NU7YWYIIOVSNQVDFIWAABQESYDVOPAD47C6X73MM7G745",
			IsIssuerAccountLock: false,
			DistributorAddress: "GCCCMTJHFLCE53AYVQJB44PHJRYRZ3A5EMRA2G4LWKENTJQ57K3SDRLA",
			DistributorSeedKey: "SBNVVHEWVKNNDPBJTK3ZJOHFMS5WQKN2EBJHWMBNWPON3FA6RXCEMX7N",
			IsDistributorAccountNeedToCreate: true,
			IsDistributorAccountNeedToTrust: true,
		}

		resp, err := stellar.IssueAsset(hc, params)

		assert.NotEmpty(t, resp)
		assert.NoError(t, err)
	})
}
