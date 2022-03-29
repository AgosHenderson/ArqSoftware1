package div

import (
	"errors"
)

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir")
	}
	return a / b, nil
}
