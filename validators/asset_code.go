package validators

import (
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func AssetCode(input string) error {
	trimmedInput := strings.TrimSpace(input)

	if len(trimmedInput) == 0 {
		return errors.New("hey, asset code is required")
	}

	if len(trimmedInput) > 12 {
		return errors.New("hey, asset code must not over 12 letters")
	}

	if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(trimmedInput) {
		return errors.New("hey, asset code must only contains letters")
	}

	return nil
}
