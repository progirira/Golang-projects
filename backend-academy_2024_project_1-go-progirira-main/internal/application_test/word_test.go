package application_test_test

import (
	"backend-academy_2024_project_1-go-progirira/internal/application"
	"testing"
)

func TestInitWordName(t *testing.T) {
	tests := map[string]struct {
		word                   string
		hint                   string
		expectedGuessedLetters string
	}{
		"empty": {
			word:                   "",
			hint:                   "",
			expectedGuessedLetters: "",
		},
		"one word": {
			word:                   "лиссабон",
			hint:                   "Столица Португалии.",
			expectedGuessedLetters: "________",
		},
		"expression with whitespace": {
			word:                   "собачье сердце",
			hint:                   "Рассказ о превращении пса в человека.",
			expectedGuessedLetters: "_______ ______",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := application.InitWord(tt.word, tt.hint)
			if got.ReturnWord() != tt.word {
				t.Errorf("got: %s, expected: %s", got.ReturnWord(), tt.word)
			}

			if got.ReturnGuessedLetters() != tt.expectedGuessedLetters {
				t.Errorf("got: %s, expected: %s", got.ReturnGuessedLetters(), tt.expectedGuessedLetters)
			}

			if got.ReturnHint() != tt.hint {
				t.Errorf("got: %s, expected: %s", got.ReturnHint(), tt.hint)
			}
		})
	}
}

func TestIsAlreadyGuessed(t *testing.T) {
	tests := map[string]struct {
		input        string
		guessedInput []rune
		letter       rune
		expected     bool
	}{
		"is guessed": {
			input:        "лимасол",
			guessedInput: []rune("__м__о_"),
			letter:       'о',
			expected:     true,
		},
		"is not guessed": {
			input:        "париж",
			guessedInput: []rune("_____"),
			letter:       'а',
			expected:     false,
		},
		"is not in word": {
			input:        "париж",
			guessedInput: []rune("_____"),
			letter:       'у',
			expected:     false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			word := application.InitWord(tc.input, "")
			word.SetGuessedLetters(tc.guessedInput)
			got := word.IsAlreadyGuessed(tc.letter)

			if got != tc.expected {
				t.Errorf("got: %v, expected: %v", got, tc.expected)
			}
		})
	}
}

func TestIsLetterIsInWord(t *testing.T) {
	tests := map[string]struct {
		input    string
		letter   rune
		expected bool
	}{
		"is in word": {
			input:    "лимасол",
			letter:   'о',
			expected: true,
		},
		"not in word": {
			input:    "собачье сердце",
			letter:   'л',
			expected: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			word := application.InitWord(tc.input, "")
			got := word.IsLetterIsInWord(tc.letter)

			if got != tc.expected {
				t.Errorf("got: %v, expected: %v", got, tc.expected)
			}
		})
	}
}

func TestOpenLetter(t *testing.T) {
	tests := map[string]struct {
		input    string
		letter   rune
		expected string
	}{
		"no letter": {
			input:    "париж",
			letter:   'о',
			expected: "_____",
		},
		"one letter": {
			input:    "лиссабон",
			letter:   'о',
			expected: "______о_",
		},
		"more than one letter": {
			input:    "магнитная гора",
			letter:   'а',
			expected: "_а_____а_ ___а",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			word := application.InitWord(tt.input, "")
			word.OpenLetter(tt.letter)
			got := word.ReturnGuessedLetters()

			if got != tt.expected {
				t.Errorf("got: %s, expected: %s", got, tt.expected)
			}
		})
	}
}

func TestAddLetterToUsed(t *testing.T) {
	tests := map[string]struct {
		input          []rune
		letterToAppend rune
		expected       []rune
	}{
		"empty in the beginning": {
			input:          []rune{},
			letterToAppend: 'о',
			expected:       []rune{'о'},
		},
		"not empty in the beginning": {
			input:          []rune{'а', 'б'},
			letterToAppend: 'о',
			expected:       []rune{'а', 'б', 'о'},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			word := application.InitWord("", "")
			word.SetUsedLetters(tc.input)
			word.AddLetterToUsed(tc.letterToAppend)
			got := word.ReturnUsedLetters()

			for i := range got {
				if got[i] != tc.expected[i] {
					t.Errorf("got: %s, expected: %s", string(got), string(tc.expected))
				}
			}
		})
	}
}

func TestIsAllLettersGuessed(t *testing.T) {
	tests := map[string]struct {
		input    []rune
		expected bool
	}{
		"all are guessed": {
			input:    []rune("москва"),
			expected: true,
		},
		"all are guessed in expression with blanks": {
			input:    []rune("набережная реки урал"),
			expected: true,
		},
		"not all are guessed": {
			input:    []rune("мо__ва"),
			expected: false,
		},
		"none of letters are guessed": {
			input:    []rune("_________"),
			expected: false,
		},
		"none of letters are guessed in expression with blanks": {
			input:    []rune("____ ___ __"),
			expected: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			word := application.InitWord("", "")
			word.SetGuessedLetters(tc.input)
			got := word.IsAllLettersGuessed()

			if got != tc.expected {
				t.Errorf("got: %v, expected: %v", got, tc.expected)
			}
		})
	}
}
