package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")

	// TeeReader: 読みながら buffer にもコピー
	teeReader := io.TeeReader(reader, &buffer)

	// io.ReadAllでteeReaderから全部読み取る（結果は捨てる）
	_, err := io.ReadAll(teeReader)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	// Teeされた内容がbufferにコピーされている
	fmt.Println(buffer.String())
}
