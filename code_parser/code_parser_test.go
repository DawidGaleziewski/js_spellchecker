package code_parser

import (
	"fmt"
	"testing"
)

// func TestFindFilesTS(t *testing.T) {
// 	fmt.Println("######test", CodeParser{})
// 	// var CP = CodeParser{}
// 	// result := CP.FindFiles("../example", ".*\\.ts");
// 	// fmt.Println(result)

// 	// if(false){
// 	// 	t.Error("wrong result")
// 	// }
// }

func TestFindDefinitions(t *testing.T) {
	//fmt.Println("######test", CodeParser{})
	// var CP = CodeParser{}
	// result := CP.FindFiles("../example", ".*\\.ts");
	// fmt.Println(result)

	// if(false){
	// 	t.Error("wrong result")
	// }
	results := Toolbox.FindDefinitions(CodeBlob{Origin{}, "const TestFunction = () => null"}, "(const) ([^ \n]*)")
	fmt.Println("###", results)
}