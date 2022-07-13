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

const helpText = `
spellchecker v0.0.1

# arguments:
- model [path to json file with words] - train spellcheck on provided words in json file
`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(helpText)
		return
	}

	firstParam := os.Args[1]

	switch firstParam {
	case "model":
		dictionaryJSONPath := os.Args[2]
		GenerateFuzzyModel(dictionaryJSONPath, "assets/fuzzy_model.go")
	case "js":
		codeTextForSpellCheck := os.Args[2]
		fmt.Println(SpellCheckJavaScriptVariables(codeTextForSpellCheck, "assets/fuzzy_model.json"))
	default:
		fmt.Println(helpText)
	}
}

func SpellCheckJavaScriptVariables(code string, trainedModelPath string) string {
	// TODO: this is the most resource heavy task. Need to at leas run this before suggesting each word and maybe decrese the number of dictionary words
	model, err := fuzzy.Load(trainedModelPath)
	if err != nil {
		log.Println(err)
	}

	return model.SpellCheck(code)
}

// SuggestWord is a live checkup that takes dictionary and trains the model in runtime. It is quite costly
func SuggestWord(searchTerm string, dictionary map[string]string) string {
	model := fuzzy.NewModel()

	// For testing only, this is not advisable on production
	model.SetThreshold(1)

	// This expands the distance searched, but costs more resources (memory and time).
	// For spell checking, "2" is typically enough, for query suggestions this can be higher
	model.SetDepth(2)

	// Train multiple words simultaneously by passing an array of strings to the "Train" function

	var words []string

	for key, _ := range dictionary {
		words = append(words, key)
	}

	// TODO: this is the most resource heavy task. Need to at leas run this before suggesting each word and maybe decrese the number of dictionary words
	model.Train(words)

	return model.SpellCheck(searchTerm)
}

// Generate a fuzzy model for spell checking and saves it for later use
func GenerateFuzzyModel(wordsDictJsonPath string, outputPath string) {
	// load and parse json
	jsonFile, err := os.Open(wordsDictJsonPath)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}

	var dictionary map[string]string
	json.Unmarshal([]byte(byteValue), &dictionary)

	// train model on provided values
	model := fuzzy.NewModel()

	// For testing only, this is not advisable on production
	model.SetThreshold(1)

	// This expands the distance searched, but costs more resources (memory and time).
	// For spell checking, "2" is typically enough, for query suggestions this can be higher
	model.SetDepth(2)

	// Train multiple words simultaneously by passing an array of strings to the "Train" function
	var words []string

	for key, _ := range dictionary {
		words = append(words, key)
	}

	// TODO: this is the most resource heavy task.
	model.Train(words)

	model.Save(outputPath)
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
