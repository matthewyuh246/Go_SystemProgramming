package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダーを表示してみる
	fmt.Println(res.Header)
	// ボディーを表示してみる。最後にはClose()すること
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}