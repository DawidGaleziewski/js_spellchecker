package main

import (
	"fmt"
	"io/fs"
	"js_tools/spellchecker/model"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const helpText = `
spellchecker v0.0.1

# arguments:
- model [path to json file with words] - train spellcheck on provided words in json file
`

const defaultModelPath = "assets/fuzzy_model.go"

func main() {
	if len(os.Args) == 1 {
		fmt.Println(helpText)
		return
	}

	firstParam := os.Args[1]

	switch firstParam {
	case "model":
		dictionaryJSONPath := os.Args[2]
		model.Learn(dictionaryJSONPath, defaultModelPath)
	case "js":
		var filePaths []string 
		filepath.Walk("./example", func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return err
			}
			filePaths = append(filePaths, path)
			fmt.Println(path)
			
			return nil
		})

		 jsFilePaths := filterByRegex(filePaths, ".*ts")
		 fmt.Println(jsFilePaths)
		// codeTextForSpellCheck := os.Args[2]
		// CP := code_parser.CodeParser{}
		// CP.Search
	default:
		fmt.Println(helpText)
	}
}

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
