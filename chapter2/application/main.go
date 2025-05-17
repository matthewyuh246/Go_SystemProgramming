package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

/*
1. 汎用コピー関数
   Reader から Writer へデータを転送するだけのシンプルな関数。
*/
func CopyData(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

/*
2. ヘルパー関数
   bufio を使ってバッファ付きで転送する例。バッファサイズを調整できる。
*/
func BufferedCopy(dst io.Writer, src io.Reader, bufSize int) (int64, error) {
	br := bufio.NewReaderSize(src, bufSize)
	bw := bufio.NewWriterSize(dst, bufSize)
	defer bw.Flush()
	return io.Copy(bw, br)
}

/*
3. ラッパー構造体 + ログ付き Writer
   Write のたびに時刻と書き込んだバイト数を標準出力へログ出力する例。
*/
type LoggingWriter struct {
	w io.Writer
}

func NewLoggingWriter(w io.Writer) *LoggingWriter {
	return &LoggingWriter{w: w}
}

func (lw *LoggingWriter) Write(p []byte) (n int, err error) {
	start := time.Now()
	n, err = lw.w.Write(p)
	fmt.Printf("[%s] wrote %d bytes\n", start.Format(time.RFC3339), n)
	return
}

func main() {
	// ファイルを「読み書きモード＋存在しなければ作成」でオープン
	f, err := os.OpenFile("input.txt", 	
		os.O_RDWR|os.O_CREATE,			// 読み書き＋なければ作成
		0644,							// パーミッション rw-r--r--
	)
	if err != nil {
		panic(fmt.Errorf("ファイルオープン失敗: %w", err))
	}
	defer f.Close()

	// 出力先を標準出力（Writer）に設定
    // さらに LoggingWriter でラップ
	logger := NewLoggingWriter(os.Stdout)

	// 1) シンプルコピー
	fmt.Println("---- CopyData ----")
	if _, err := CopyData(logger, f); err != nil {
		fmt.Fprintln(os.Stderr, "CopyData error:", err)
	}

	// ファイル位置を先頭に戻す
	f.Seek(0, io.SeekStart)

	// 2) バッファ付きコピー（ヘルパー関数）
	fmt.Println("---- BufferedCopy (4KB) ----")
	if _, err := BufferedCopy(logger, f, 4*1024); err != nil {
		fmt.Fprintln(os.Stderr, "BufferedCopy error:", err)
	}
}