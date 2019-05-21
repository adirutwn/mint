package validators

import (
	"github.com/pkg/errors"
	"github.com/asaskevich/govalidator"
)

func Url(input string) error {
	ok := govalidator.IsDNSName(input)
	if !ok {
		return errors.New("well, it's not url no?")
	}


	return nil
}
