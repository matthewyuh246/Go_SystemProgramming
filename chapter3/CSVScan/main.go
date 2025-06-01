package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `1,Alice,Tokyo,25,Female,Engineer,Go,Rust,Python
2,Bob,Osaka,30,Male,Designer,JavaScript,HTML,CSS`

func main() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			continue
		}
		// 各行の 3列目 (index=2) と 7〜9列目 (index=6〜8) を出力
		fmt.Println("City:", line[2], "| Skills:", line[6:9])
	}
}
