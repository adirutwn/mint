package validators

import (
	"github.com/pkg/errors"
	"github.com/stellar/go/strkey"
	"strings"
)

func StellarAddress(input string) error {
	trimmedInput := strings.TrimSpace(input)

	_, err := strkey.Decode(strkey.VersionByteAccountID, trimmedInput)
	if err != nil {
		return errors.New("wait, this is not a stellar address no?")
	}

	return nil
}
