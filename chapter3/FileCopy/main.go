package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	destFile, err := os.Create("input2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		panic(err)
	}
}