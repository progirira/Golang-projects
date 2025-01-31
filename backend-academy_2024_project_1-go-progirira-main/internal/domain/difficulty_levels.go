package domain

type Difficulty int

const (
	light Difficulty = iota + 1
	average
	hard
)

var difficultyLevelsInRussian = [...]string{"лёгкий", "средний", "сложный"}

func TheLightestDifficulty() Difficulty {
	return light
}

func TheHardestDifficulty() Difficulty {
	return hard
}

func DifficultyInRussian(d Difficulty) string {
	return difficultyLevelsInRussian[d-1]
}
