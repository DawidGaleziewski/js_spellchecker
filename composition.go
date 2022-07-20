package main

import (
	"regexp"
	"strings"
)

// DATA STRUCTURES
// Type CodeBlob defines code string that is unparsed
type CodeBlob struct{
    Origin;
    blob string;
}

// Type Origin describes the origin of the code
type Origin struct{
    filePath string;
    lineNO int;
    lang string;
}

// Type Definition describes a single definition discovered in the code and possible words its made of. Definition is a word used to declare variable, function or key in a object
type Definition struct {
    Origin;
    name string;
    words []string;
}

// Type Suggestion defines words suggestion for a definition
type Suggestion struct{
    words [][]string;
    phrases []string;
}




// INTERFACES lV1
type  DefinitionFinder interface{
	FindDefinitions(code CodeBlob) []Definition
}

type WordFinder interface  {
	FindWords(variableName string) []string
}

type WordChecker interface {
	CheckWord(word string) bool
}

type WordSuggester interface {
	SuggestWord(string) []string
}

type PhraseSuggester interface {
	SuggestPhrase(string) []string
}

// INTERFACES LV2
type CodeParser struct {
	DefinitionFinder
	WordFinder
}

type CodeSuggester interface {
	WordChecker
	WordSuggester
	PhraseSuggester
}

// IMPLEMENTATION

type JSDefFinder struct {}

func (JSDefFinder)FindDefinition(code CodeBlob) []Definition{
	r := regexp.MustCompile("(const) ([^ \n]*)")
	captureGroups := r.FindAllStringSubmatch(code.blob, -1)

	var definitions []Definition
	//var variableNames []string

	// no matches were found
	if captureGroups == nil {
		return  []Definition{}
	}

	for _, group := range captureGroups {
		variableName := group[2]
		definition := Definition{
			name: variableName,
			Origin: code.Origin,
		}

		definitions = append(definitions, definition)

	}

	return definitions
}

func (JS)FindWords(variableName string) []string{
		r := regexp.MustCompile("[A-Z]")
	indexGroups := r.FindAllStringIndex(variableName, -1)

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




// Bootstraping all together
func main_2(){
	JSParser := CodeParser{
		DefinitionFinder: JSDefFinder{},
		WordFinder: JS{},
	}
}