package main

import (
	"regexp"
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
	FindWords(blob string) []string
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
type CodeParser interface {
	DefinitionFinder
	WordFinder
}

type CodeSuggester interface {
	WordChecker
	WordSuggester
	PhraseSuggester
}

// IMPLEMENTATION

type JS struct {}



func (JS)getDef(code CodeBlob) []Definition{
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
		}

	}

	return []Definition{}
}
