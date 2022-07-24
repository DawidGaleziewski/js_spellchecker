package code_parser

import (
	"fmt"
	"testing"
)

func TestFindDefinitions(t *testing.T) {
	results := ParseJavaScript("../example")
	fmt.Println("###", results)
}