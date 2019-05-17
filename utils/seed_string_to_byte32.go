package utils

import (
	"github.com/stellar/go/strkey"
)

func SeedStringToByte32(seedKey string) (*[32]byte, error) {
	rawSeed, err := strkey.Decode(strkey.VersionByteSeed, seedKey)
	if err != nil {
		return nil, err
	}

	rawSeed32 := [32]byte{}
	for i := 0; i < 32; i++ {
		rawSeed32[i] = rawSeed[i]
	}

	return &rawSeed32, nil
}
