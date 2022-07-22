package code_parser

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

// INTERFACES lV1
type DefinitionFinder interface{
	FindDefinitions(code CodeBlob, regexPattern string) []Definition
}

type WordSpliter interface  {
	SplitWords(variableName string, regexPattern string) []string
}	

// INTERFACES LV2
type Parser interface{
	Parse(string) []Definition
}	

type CodeParser struct {
	DefinitionFinder
	WordSpliter
	Parser
	definitions []Definition
}


// IMPLEMENTATION
type Search struct {}
var DeclarationPattern = map[string]string{
	"JS": "(const) ([^ \n]*)",
}

func (Search)FindDefinitions(code CodeBlob, regexPattern string) []Definition{
	r := regexp.MustCompile(regexPattern)
	captureGroups := r.FindAllStringSubmatch(code.blob, -1)
	var definitions []Definition

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

type Split struct {}
var CaseNewWordPattern = map[string]string{
	"CAMEL_CASE": "[A-Z]",
}
func (Split)SplitWords(variableName string, regexPattern string) []string{
	r := regexp.MustCompile(regexPattern)
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


type Seaker struct{}
func (CP CodeParser)ParseJS(codeBlob CodeBlob) []Definition {
	definitions := CP.FindDefinitions(codeBlob, DeclarationPattern["JS"]);

	for i, def := range definitions {
		words := CP.SplitWords(def.name, CaseNewWordPattern["CAMEL_CASE"]);
		definitions[i].words = append(definitions[i].words, words...) 
	}

	return definitions
}

var Toolbox = CodeParser{
		DefinitionFinder: Search{},
		WordSpliter: Split{},
	}


// Bootstraping all together
// func main(){

	
// 	definitions := JS.ParseJS(CodeBlob{blob: "const TestVariable1("})
// 	fmt.Println(definitions[0].words)

// }