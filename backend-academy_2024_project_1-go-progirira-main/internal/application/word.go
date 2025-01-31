package application

type Word struct {
	name           []rune
	guessedLetters []rune
	hint           string
	usedLetters    []rune
}

func InitWord(word, hint string) *Word {
	runes := []rune(word)
	guessedLetters := make([]rune, len(runes))

	for i, r := range runes {
		if r == ' ' {
			guessedLetters[i] = ' '
		} else {
			guessedLetters[i] = '_'
		}
	}

	return &Word{
		name:           runes,
		guessedLetters: guessedLetters,
		hint:           hint,
		usedLetters:    []rune{},
	}
}

func (word *Word) IsAlreadyGuessed(letter rune) bool {
	for _, sym := range word.guessedLetters {
		if sym == letter {
			return true
		}
	}

	return false
}

func (word *Word) IsLetterIsInWord(letter rune) bool {
	for _, sym := range word.name {
		if string(sym) == string(letter) {
			return true
		}
	}

	return false
}

func (word *Word) OpenLetter(letter rune) {
	for ind, sym := range word.name {
		if sym == letter {
			word.guessedLetters[ind] = letter
		}
	}
}

func (word *Word) ReturnGuessedLetters() string {
	return string(word.guessedLetters)
}

func (word *Word) ReturnWord() string {
	return string(word.name)
}

func (word *Word) AddLetterToUsed(letter rune) {
	word.usedLetters = append(word.usedLetters, letter)
}

func (word *Word) IsAllLettersGuessed() bool {
	for _, sym := range word.guessedLetters {
		if string(sym) == "_" {
			return false
		}
	}

	return true
}

func (word *Word) ReturnHint() string {
	return word.hint
}

func (word *Word) SetGuessedLetters(runes []rune) {
	word.guessedLetters = runes
}

func (word *Word) SetUsedLetters(runes []rune) {
	word.usedLetters = runes
}

func (word *Word) ReturnUsedLetters() []rune {
	return word.usedLetters
}
