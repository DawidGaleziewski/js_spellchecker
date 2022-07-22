package main

import (
	"fmt"
	"js_tools/spellchecker/model"
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
		codeTextForSpellCheck := os.Args[2]
		CodeParser
		fmt.Println(SpellCheckJavaScriptVariables(codeTextForSpellCheck, defaultModelPath))
	default:
		fmt.Println(helpText)
	}
}
