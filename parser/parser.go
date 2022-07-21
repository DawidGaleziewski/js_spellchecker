package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type definition struct {
	originalVariableName string
	wordsInVariable      []string
}

func Parse(lang string, codeBlob string) []definition {
	var vocebulary []definition

	variableNames := getVarNames(codeBlob)

	for _, name := range variableNames {
		words := splitToWords(name)
		vocebulary = append(vocebulary, definition{
			originalVariableName: name,
			wordsInVariable:      words,
		})
	}

	return vocebulary
}

// GetVariableNames accepts string and returns variable declarations in javascript if there are any
func getVarNames(textBlob string) []string {
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
func splitToWords(variableName string) []string {
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