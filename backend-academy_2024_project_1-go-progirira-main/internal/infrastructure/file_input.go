package infrastructure

import (
	"bufio"
	"crypto/rand"
	"errors"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	dictionaryDirectory  = "words/"
	hangmanFile          = "hangman_stages.txt"
	errInvalidString     = "неверный формат строки"
	errNotSuchComplexity = "не найдено слов с нужной сложностью"
)

func readDictionaryFromFile(filename string) (dict []string, complexitiesDict map[string]int, hintsDict map[string]string, e error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, err
	}

	defer file.Close()

	var words []string

	var complexities = make(map[string]int)

	var hints = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		if len(parts) != 3 {
			return nil, nil, nil, errors.New(errInvalidString)
		}

		chars, errAtoi := strconv.Atoi(strings.TrimSpace(parts[1]))
		if errAtoi != nil {
			return nil, nil, nil, err
		}

		wordName := strings.TrimSpace(parts[0])
		words = append(words, wordName)
		complexities[wordName] = chars
		hints[wordName] = strings.TrimSpace(parts[2])
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, nil, err
	}

	return words, complexities, hints, nil
}

func RandomWordFromFile(category string, requiredComplexity int) (word, hint string, err error) {
	filename := dictionaryDirectory + category + ".txt"
	words, complexities, hints, _ := readDictionaryFromFile(filename)

	var matchedWords []string

	for _, word := range words {
		if complexities[word] == requiredComplexity {
			matchedWords = append(matchedWords, word)
		}
	}

	if len(matchedWords) == 0 {
		return "", "", errors.New(errNotSuchComplexity)
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(matchedWords)+1)))
	if err != nil {
		return "", "", err
	}

	gameWord := words[int(nBig.Int64())]

	return gameWord, hints[gameWord], nil
}

func LoadStagesOfHangman() []string {
	filePath := hangmanFile

	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}

	defer file.Close()

	var stages []string

	var currentStage strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			if currentStage.Len() > 0 {
				stages = append(stages, currentStage.String())
				currentStage.Reset()
			}
		} else {
			currentStage.WriteString(line + "\n")
		}
	}

	if currentStage.Len() > 0 {
		stages = append(stages, currentStage.String())
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return stages
}
