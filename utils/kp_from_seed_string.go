package utils

import (
	"github.com/pkg/errors"
	"github.com/stellar/go/keypair"
)

func KpFromSeedString(seed string) (*keypair.Full, error) {
	seedKeyb, err := SeedStringToByte32(seed)
	if err != nil {
		return nil, errors.New("[😱] hmmm.. I'm not able to convert seed key to byte")
	}

	kp, err := keypair.FromRawSeed(*seedKeyb)
	if err != nil {
		return nil, errors.New("[😱] hmmm.. I'm not able to get keypair from issuer's seed key")
	}

	return kp, nil
}
