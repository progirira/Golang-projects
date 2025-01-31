package application

import (
	"backend-academy_2024_project_1-go-progirira/internal/infrastructure"
	"fmt"
)

type Game struct {
	maxNumOfAttempts int
	word             *Word
	remainingLives   *Lives
	hangman          *Hangman
}

func InitGame() Game {
	maxNumOfAttempts := 7
	lives, err := InitLives(maxNumOfAttempts)

	if err != nil {
		fmt.Println(err.Error())
	}

	return Game{
		maxNumOfAttempts: maxNumOfAttempts,
		remainingLives:   lives,
		hangman:          InitHangman(),
	}
}

func (game *Game) displayHangman() {
	remainedLives := game.remainingLives.returnNumOfLives()
	fmt.Println(game.hangman.returnExactHangman(game.maxNumOfAttempts - remainedLives - 1))
}

func (game *Game) Start() {
	infrastructure.InfoBeforeGame(game.maxNumOfAttempts)

	gameWord, hint := game.getANewWord()

	game.word = InitWord(gameWord, hint)

	infrastructure.DisplayStart()
	game.GuessProcess()
}

func (game *Game) GuessProcess() {
	for {
		if infrastructure.AskIfHint() {
			infrastructure.PrintString(game.word.ReturnHint())
		}

		letter := infrastructure.AskForLetter()

		if game.word.IsAlreadyGuessed(letter) {
			infrastructure.AlreadyGuessed(letter)
			continue
		}

		if game.word.IsLetterIsInWord(letter) {
			game.word.OpenLetter(letter)
			infrastructure.Guessed()

			if game.word.IsAllLettersGuessed() {
				infrastructure.OutputWin(game.word.ReturnWord())
				break
			}

			game.word.AddLetterToUsed(letter)
			infrastructure.PrintCurrentString(game.word.ReturnGuessedLetters())

			continue
		}

		infrastructure.NoSuchLetter()
		game.remainingLives.DecreaseLives()

		if !game.remainingLives.IsHasLives() {
			infrastructure.OutputLose(game.word.ReturnWord())
			game.displayHangman()

			break
		}

		game.word.AddLetterToUsed(letter)
		game.displayHangman()
		infrastructure.PrintCurrentString(game.word.ReturnGuessedLetters())
	}
}

func (game *Game) getANewWord() (word, hint string) {
	category, complexity := infrastructure.ChooseParametersByAskingUser()
	if category == "" {
		category = GetRandomCategory()
	}

	if complexity == 0 {
		var err error
		complexity, err = GetRandomInt(1, 3)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	gameWord, hint, _ := infrastructure.RandomWordFromFile(category, complexity)

	return gameWord, hint
}
