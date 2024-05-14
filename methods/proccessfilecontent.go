package methods

import (
	"fmt"
	"os"
	"strings"
)

// Get the file content and split it to get each character graphic representation alone, where also each character represrntation splited to 8 lines
func ProccessFileContent(fileName string) [][]string {
	// get file content............................................................
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// split the content by character graphics.....................................
	splitedContent := strings.Split(string(fileContent[1:]), "\n\n")

	// split character graphics to lines...........................................
	graphics := splitGraphics(splitedContent)

	return graphics
}

// used to help in proccessing the file content
func splitGraphics(slice []string) [][]string {

	var ret [][]string
	for i := 0; i < len(slice); i++ {
		ret = append(ret, strings.Split(slice[i], "\n"))
	}

	return ret
}
