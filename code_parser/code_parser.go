package code_parser

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
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
type FileFinder interface {
	FindFiles(dir string, regexPattern string) []string 
}

type DefinitionFinder interface{
	FindDefinitions(code CodeBlob, regexPattern string, splitWordsFn func (string)[]string) []Definition
}

type WordSpliter interface  {
	SplitWords(variableName string, regexPattern string) []string
}	

// INTERFACES LV2
type Parser interface{
	Parse(string) []Definition
}	

type CodeParser struct {
	FileFinder
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

var FilePattern = map[string]string {
	"TS":  ".*\\.ts",
}

func (Search)FindFiles(dir string, regexPattern string) []string {
	fmt.Println("From find files func")
	var filePaths []string 
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				log.Println("FindFiles error during Walk",err)
				return err
			}
			filePaths = append(filePaths, path)
			fmt.Println(path)
			
			return nil
		})

	jsFilePaths := filterByRegex(filePaths, regexPattern)
	fmt.Println(jsFilePaths)
	return jsFilePaths
}

func (Search)FindDefinitions(code CodeBlob, regexPattern string, splitWordsFN func(name string)[]string) []Definition{
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
			words: splitWordsFN(variableName),
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

var Toolbox = CodeParser{
		DefinitionFinder: Search{},
		FileFinder: Search{},
		WordSpliter: Split{},
	}


// Bootstraping all together
func ParseJavaScript(dir string){
	var definitions []Definition
	filePaths := Toolbox.FindFiles(dir, FilePattern["TS"]);

	for _, filePath := range filePaths {
		data, err := os.ReadFile(filePath)
		if(err != nil){
			fmt.Println(err)
		}
		def := Toolbox.FindDefinitions(CodeBlob{Origin: Origin{filePath: filePath}, blob: string(data) }, DeclarationPattern["JS"], func(name string) []string{
			result := Toolbox.SplitWords(name, CaseNewWordPattern["CAMEL_CASE"])
			return result
		});
		definitions = append(definitions, def...)
	}

	fmt.Println(definitions)
}


// utils
func filterByRegex (arr []string, regexPattern string) []string {
	var results []string;

	for _, item := range arr {
		doesMatch, err := regexp.MatchString(regexPattern, item)
		if err != nil {
			log.Println(err)
		}
		if(doesMatch){
			results = append(results, item)
		}
	}

	return results
}