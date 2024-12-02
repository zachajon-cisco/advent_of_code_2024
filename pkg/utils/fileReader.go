package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInFile(fileName string) (fileLines []string, err error) {

	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return
}
