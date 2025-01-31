package infrastructure

import (
	"backend-academy_2024_project_1-go-progirira/internal/domain"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func InfoBeforeGame(maxNumOfAttempts int) {
	fmt.Println("Здравствуй, дорогой игрок! Тебя приветствует Виселица.\nВ процессе игры ты поборешься"+
		" за жизнь куклы, одна неверная буква - и она на шаг ближе к повешанию.\nУвлекательного путешествия...\n"+
		"Максимальное количество попыток: ", maxNumOfAttempts)
}

func ChooseParametersByAskingUser() (category string, difficulty int) {
	return askForCategory(), askForDifficulty()
}

func askForCategory() string {
	var numOfCategory string

	fmt.Println("Выберите категорию, введя соответствующую цифру")

	for i := domain.FirstCategory(); i <= domain.LastCategory(); i++ {
		fmt.Println(i, domain.CategoryInRussian(i))
	}

	_, err := fmt.Scanf("%s", &numOfCategory)
	if err != nil {
		return ""
	}

	numOfCategory = strings.TrimSpace(numOfCategory)

	categoryInt, errAtoi := strconv.Atoi(numOfCategory)
	if errAtoi != nil {
		return ""
	}

	return domain.CategoryFileNameByIndex(categoryInt)
}

func askForDifficulty() int {
	var difficulty string

	fmt.Println("Выберите уровень сложности, введя соответствующую цифру")

	for i := domain.TheLightestDifficulty(); i <= domain.TheHardestDifficulty(); i++ {
		fmt.Println(i, domain.DifficultyInRussian(i))
	}

	_, err := fmt.Scanf("%s", &difficulty)
	if err != nil {
		return 0
	}

	difficultyInt, err := strconv.Atoi(strings.TrimSpace(difficulty))
	if err != nil {
		return 0
	}

	return difficultyInt
}

func DisplayStart() {
	fmt.Println("Начнём игру!")
}

func AskForLetter() rune {
	var guess string

	fmt.Println("Введите букву: ")

	_, err := fmt.Scanf("%s", &guess)
	if err != nil {
		return 0
	}

	if guess == "" {
		fmt.Println("Пожалуйста, введите букву.")

		_, err := fmt.Scanf("%s", &guess)
		if err != nil {
			return 0
		}
	}

	letter, _ := utf8.DecodeRuneInString(strings.ToLower(guess))

	return letter
}

func OutputWin(word string) {
	fmt.Println("Вы выиграли! Угаданное Вами слово: ", word)
}

func OutputLose(word string) {
	fmt.Println("Вы проиграли( Было загадано слово ", word)
}

func AlreadyGuessed(letter rune) {
	fmt.Println("Вы уже угадывали эту букву: ", string(letter))
}

func Guessed() {
	fmt.Println("Буква угадана!")
}

func NoSuchLetter() {
	fmt.Println("Такой буквы нет:(")
}

func PrintCurrentString(word string) {
	fmt.Println("Слово теперь выглядит так: ", word)
}

func AskIfHint() bool {
	fmt.Println("Хотите получить подсказку? Введите +, если нет - введите любой другой символ.")

	var guess string

	_, err := fmt.Scanf("%s", &guess)
	if err != nil {
		return false
	}

	if guess == "+" {
		return true
	}

	return false
}

func PrintString(line string) {
	fmt.Println(line)
}
