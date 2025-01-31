package application

import (
	"backend-academy_2024_project_1-go-progirira/internal/domain"
	"crypto/rand"
	"fmt"
	"math/big"
)

func GetRandomInt(a, b int) (randNum int, err error) {
	if a > b {
		err := NewErrInvalidBoarders(a, b)
		return 0, err
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(b-a+1)))
	if err != nil {
		return 0, err
	}

	return int(nBig.Int64()) + a, nil
}

func GetRandomCategory() (category string) {
	categoryNum, err := GetRandomInt(1, domain.CategoriesCount())

	if err != nil {
		fmt.Println(err.Error())
	}

	return domain.CategoryFileNameByIndex(categoryNum)
}
