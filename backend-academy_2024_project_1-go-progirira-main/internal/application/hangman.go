package application

import "backend-academy_2024_project_1-go-progirira/internal/infrastructure"

type Hangman struct {
	stages []string
}

func InitHangman() *Hangman {
	return &Hangman{
		stages: infrastructure.LoadStagesOfHangman(),
	}
}

func (hangman *Hangman) returnExactHangman(ind int) string {
	return hangman.stages[ind]
}
