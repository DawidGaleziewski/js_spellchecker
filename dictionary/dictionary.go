package dictionary

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

