package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/sajari/fuzzy"
)

// Learn generate a fuzzy model for spell checking and saves it for later use
func Learn(wordsDictJsonPath string, outputPath string) {
	// load and parse json
	jsonFile, err := os.Open(wordsDictJsonPath)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}

	var dictionary map[string]string
	json.Unmarshal([]byte(byteValue), &dictionary)

	// train model on provided values
	model := fuzzy.NewModel()

	// For testing only, this is not advisable on production
	model.SetThreshold(1)

	// This expands the distance searched, but costs more resources (memory and time).
	// For spell checking, "2" is typically enough, for query suggestions this can be higher
	model.SetDepth(2)

	// Train multiple words simultaneously by passing an array of strings to the "Train" function
	var words []string

	for key, _ := range dictionary {
		words = append(words, key)
	}

	// TODO: this is the most resource heavy task.
	model.Train(words)

	model.Save(outputPath)
}
