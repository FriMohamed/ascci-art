package main
import (
	"ascii/methods"
	"fmt"
	"os"
)
func main () {
	// Check arguments :
	args := os.Args;
	if len(args) != 2 {
		fmt.Println("Expected: go run . <input_string>");
		os.Exit(1)
	}
	
	// Proccess file content :
	FileContent := methods.ProccessFileContent("standard.txt")

	// Proccess the output
	input := os.Args[1]
	if len(input) == 0 {
		return 
	} else {
		fmt.Printf(methods.ProccessOutput(input, FileContent))
	}
}