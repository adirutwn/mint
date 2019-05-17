package stellar

import (
	"github.com/AlekSi/pointer"
	"github.com/stellar/go/protocols/horizon"
)

func IsAccountTrustedAsset(account horizon.Account, assetCode string, assetIssuer string) (*bool, error) {
	for _, bal := range account.Balances {
		if bal.Asset.Code == assetCode && bal.Asset.Issuer == assetIssuer {
			return pointer.ToBool(true), nil
		}
	}

	return pointer.ToBool(false), nil
}
