package application

type Lives int

func InitLives(maxNumOfAttempts int) (*Lives, error) {
	if maxNumOfAttempts <= 0 {
		return nil, NewErrNonPositiveAttempts(maxNumOfAttempts)
	}

	lives := Lives(maxNumOfAttempts)

	return &lives, nil
}

func (lives *Lives) DecreaseLives() {
	*lives--
}

func (lives *Lives) returnNumOfLives() int {
	return int(*lives)
}

func (lives *Lives) IsHasLives() bool {
	return *lives > 0
}
