package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(name string, ch chan<- string) {
	delay := time.Duration(rand.Intn(3000)) * time.Millisecond
	time.Sleep(delay)
	ch <- fmt.Sprintf("%s finished in %v", name, delay)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go worker("Worker A", ch1)
	go worker("Worker B", ch2)
	go worker("Worker C", ch3)

	timeout := time.After(2 * time.Second)

	received := 0
	for received < 3 {
		select {
		case msg := <- ch1:
			fmt.Println("[ch1]", msg)
			received++
		case msg := <- ch2:
			fmt.Println("[ch2]", msg)
			received++
		case msg := <- ch3:
			fmt.Println("[ch3]", msg)
			received++
		case <- timeout:
			fmt.Println("タイムアウト発生！処理を中断します。")
			return
		}
	}

	fmt.Println("すべてのワーカーの完了を確認しました。")
}