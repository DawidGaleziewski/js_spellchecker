package dictionary

import (
	"js_tools/spellchecker/code_parser"
	"log"
	"os"
	"regexp"
)

// INTERFACES lV1
type WordChecker interface {
	CheckWord(word string) bool
}

type WordSuggester interface {
	SuggestWord(string) []string
}

type PhraseSuggester interface {
	SuggestPhrase(string) []string
}

type SpellChecker struct {
	WordChecker
	WordSuggester
	PhraseSuggester
	dictionary []string
}

type Dictionary struct {}
func (SC SpellChecker)CheckForEnglish(word string) bool{
	return SC.CheckWord(word)
}

type Suggestion struct {
	code_parser.Definition
	IncorrectWords []string
}

func SuggestEnglish(definitions []code_parser.Definition) []Suggestion {
	var suggestions []Suggestion

	data, err := os.ReadFile("./dictionary/assets/eng_words.txt");
	if(err != nil){
	   log.Println(err)
	}
	englisWords := string(data);

	for _, def := range definitions {
		var incorrectWords []string 
		for _ , word := range def.Words {
			isAWord, err := regexp.MatchString(word, englisWords)

			if(err != nil){
				log.Println(err)
			}

			if !isAWord {
				incorrectWords = append(incorrectWords, word)
			}
		}

		if len(incorrectWords) > 0 {
			suggestions = append(suggestions, Suggestion{Definition: def, IncorrectWords: incorrectWords})
		}
	}

	return suggestions
}

