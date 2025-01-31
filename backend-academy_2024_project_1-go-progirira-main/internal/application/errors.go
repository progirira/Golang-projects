package application

import "fmt"

type ErrInvalidBoarders struct {
	leftBoarder int
	rightBoader int
}

func NewErrInvalidBoarders(a, b int) error {
	return ErrInvalidBoarders{leftBoarder: a, rightBoader: b}
}

func (e ErrInvalidBoarders) Error() string {
	return fmt.Sprintf("Min cannot be greater than max: %v > %v", e.leftBoarder, e.rightBoader)
}

type ErrNonPositiveAttempts struct {
	numOfAttempts int
}

func NewErrNonPositiveAttempts(num int) error {
	return ErrNonPositiveAttempts{numOfAttempts: num}
}

func (e ErrNonPositiveAttempts) Error() string {
	return fmt.Sprintf("Max number of attempts can not be 0 and less: %v <= 0", e.numOfAttempts)
}
