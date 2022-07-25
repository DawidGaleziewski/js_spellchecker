package msg

import (
	"fmt"
	"js_tools/spellchecker/dictionary"
)

var ANSIColor = map[string] string{
	"reset" : "\033[0m",
    "red" : "\033[31m",
    "green" : "\033[32m",
    "yellow" : "\033[33m",
    "blue" : "\033[34m",
    "purple" :"\033[35m",
    "cyan" : "\033[36m",
    "white" : "\033[37m",
}

func Chalk(word string, color string) string {
	colorCode, exists := ANSIColor[color];
	if(!exists){
		return word;
	}

	return fmt.Sprint(colorCode ,word, ANSIColor["reset"])
}

func Inform(suggestions []dictionary.Suggestion){
	fmt.Println(Chalk("----------------", "red"),"\n Found potential spelling errors in: \n");
	for _, suggestion := range suggestions {
		fmt.Println("# Definition: ", Chalk(suggestion.Name, "blue"), " for words:")

		for _, word := range suggestion.IncorrectWords {
			fmt.Println("- ", Chalk("\"" + word + "\"", "red"),)
		}

		fmt.Println("")
	}

}