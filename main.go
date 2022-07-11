package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/sajari/fuzzy"
)

func main() {

	if len(os.Args) == 1 {
		log.Println("No argument provided")
		return
	}

	inputBlob := os.Args[1]

	IsEnglishWord(inputBlob)
	// fmt.Println()
}

func SuggestWord(word string) string {
	model := fuzzy.NewModel()

	// For testing only, this is not advisable on production
	model.SetThreshold(1)

	// This expands the distance searched, but costs more resources (memory and time).
	// For spell checking, "2" is typically enough, for query suggestions this can be higher
	model.SetDepth(5)

	// Train multiple words simultaneously by passing an array of strings to the "Train" function
	words := []string{"bob", "your", "uncle", "dynamite", "delicate", "biggest", "big", "bigger", "aunty", "you're"}
	model.Train(words)

	return model.SpellCheck(word)
}

func IsEnglishWord(searchTerm string) {
	jsonFile, err := os.Open("./assets/words_dictionary.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var englishWordsDictionary map[string]string

	json.Unmarshal([]byte(byteValue), &englishWordsDictionary)

	if _, ok := englishWordsDictionary[searchTerm]; ok {
		fmt.Println("has word")
	}
}

// GetVariableNames accepts string and returns variable declarations in javascript if there are any
func GetVariableNames(textBlob string) []string {
	r := regexp.MustCompile("(const) ([^ \n]*)")
	captureGroups := r.FindAllStringSubmatch(textBlob, -1)
	var variableNames []string

	if captureGroups == nil {
		fmt.Println("no matches found")
		return variableNames
	}

	for _, group := range captureGroups {
		variableNames = append(variableNames, group[2])
	}

	return variableNames
}

// CamelCaseToWords converts string written in cammel case to slice of strings with words
func CamelCaseToWords(variableName string) []string {
	r := regexp.MustCompile("[A-Z]")
	indexGroups := r.FindAllStringIndex(variableName, -1)
	//var words []string

	var wordStartIndexes []int

	for _, indexGroup := range indexGroups {
		wordStartIndexes = append(wordStartIndexes, indexGroup[0])
	}

	hasNoUppercaseWords := len(wordStartIndexes) == 0
	onlyUppercaseWordIsAtStart := len(wordStartIndexes) == 1 && wordStartIndexes[0] == 0

	if hasNoUppercaseWords || onlyUppercaseWordIsAtStart {
		return []string{strings.ToLower(variableName)}
	}

	var words []string

	startsWithLowerCase := wordStartIndexes[0] != 0
	if startsWithLowerCase {
		word := variableName[0:wordStartIndexes[0]]
		words = append(words, strings.ToLower(word))
	}

	for i, startIndex := range wordStartIndexes {
		var word string
		var isLastFoundIndex = i+1 == len(wordStartIndexes)

		// if there are no more indexes then one. We just have one word creating the variable
		if isLastFoundIndex { // if there is no other uppercase just return the rest of characters
			word = variableName[startIndex:]
			words = append(words, strings.ToLower(word))
		} else {
			nextUppercaseIndex := wordStartIndexes[i+1]
			word = variableName[startIndex:nextUppercaseIndex]
			words = append(words, strings.ToLower(word))
		}

	}

	return words
}
