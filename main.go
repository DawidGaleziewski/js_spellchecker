package main

import (
	"fmt"
	"js_tools/spellchecker/model"
	"log"
	"os"

	"github.com/sajari/fuzzy"
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
		fmt.Println(SpellCheckJavaScriptVariables(codeTextForSpellCheck, defaultModelPath))
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
