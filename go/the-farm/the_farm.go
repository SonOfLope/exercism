package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (m *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", m.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	food, err := weightFodder.FodderAmount()

	if food < 0 && (err == nil || err == ErrScaleMalfunction) {
		return 0, errors.New("negative fodder")
	} else if food < 0 {
		return 0, errors.New("non-scale error")
	} else if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	} else if food >= 0 && err == nil {
		return food / float64(cows), err
	} else if food >= 0 && err.Error() == ErrScaleMalfunction.Error() {
		return food * 2 / float64(cows), nil
	}

	return 0, err
}
