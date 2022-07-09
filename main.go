package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) == 1 {
		log.Println("No argument provided")
		return
	}

	inputBlob := os.Args[1]

	fmt.Println(inputBlob)
}

// GetVariableNames accepts string and returns variable declarations in javascript if there are any
func GetVariableNames(textBlob string) []string {
	r := regexp.MustCompile("(const) ([^ \n]*)")
	captureGroups := r.FindAllStringSubmatch(textBlob, -1)
	var variableNames []string

	if captureGroups == nil {
		fmt.Println("no matches found")
		return variableNames
	}

	for _, group := range captureGroups {
		variableNames = append(variableNames, group[2])
	}

	return variableNames
}

// CamelCaseToWords converts string written in cammel case to slice of strings with words
func CamelCaseToWords(variableName string) []string {
	return []string{"test", "test"}
}
