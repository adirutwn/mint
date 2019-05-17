package validators

import (
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

func Amount(input string) error {
	trimmedInput := strings.TrimSpace(input)

	if len(trimmedInput) == 0 {
		return errors.New("hey, amount is required")
	}

	if !regexp.MustCompile(`^[0-9,.]+$`).MatchString(trimmedInput) {
		return errors.New("hey, amount must only contains numbers, commas, or dot")
	}

	removedCommas := strings.Replace(trimmedInput, ",", "", -1)
	spittedInput := strings.Split(removedCommas, ".")
	decimals := ""

	if len(spittedInput) == 2 {
		decimals = spittedInput[1]
	}

	if len(decimals) > 7 {
		return errors.New("hey, amount's decimal places must not over 7 places")
	}

	amountFloat, err := strconv.ParseFloat(removedCommas, 64)
	if err != nil {
		return errors.New("hey, I cannot change amount into floating number")
	}

	if amountFloat > 922337203685.4775807 {
		return errors.New("hey, amount must not over 922337203685.4775807")
	}

	return nil
}
