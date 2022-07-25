package main

import (
	"fmt"
	"js_tools/spellchecker/code_parser"
	"js_tools/spellchecker/dictionary"
	"js_tools/spellchecker/model"
	"js_tools/spellchecker/msg"
	"os"
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
		checkDirPath := os.Args[2]
		definitions := code_parser.ParseJavaScript(checkDirPath)
		suggestions := dictionary.SuggestEnglish(definitions)

		msg.Inform(suggestions)
	default:
		fmt.Println(helpText)
	}
}


