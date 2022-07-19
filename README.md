# spellchecker

Goal of the package is to check all variable devlaration in a js file and check for spelling errors.

# credit

list of english words:
https://github.com/dwyl/english-words


# composition

## data structures

Definition stores data regarding collected definitions from code

```go
type Origin struct{
    file string;
    line int;
    lang string;
}
```

```go
type Definition struct {
    Origin;
    name string;
    words []string;
}

type Suggestion struct{
    words [][]string;
    phrases []string;
}
```

CodeBlob is a blob of code with defined type
```go
type CodeBlob struct{
    Origin;
    code string;
}
```


## Behaviours lv 1

```go
    // FileSearcher interface defines behaviour for finding files in provided directory
    // interface FileSearcher {
    //     FileSearch(path string, regexPattern string) []string
    // }

    interface DefinitionFinder {
        FindDefinitions(code CodeBlob) []Definition
    }

    interface WordFinder {
        FindWords(blob string) []string
    }

    interface WordChecker {
        CheckWord(word string) bool
    }

    interface WordSuggester {
        SuggestWord(string) []string
    }

    interface PhraseSuggester {
        SuggestPhrase(string) []string
    }

    // CodeParser defines behaviour for parsing code from provided blob
    interface CodeParser {
        ParseCode(code CodeBlob) []Definition
    }
```

## basic implementations lv 1
```go
    func (code CodeBlob)getJSDefinitions(){

    }
```


