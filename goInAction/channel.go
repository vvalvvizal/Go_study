package main

import (
	"fmt"
)

func worker(ch chan string) {
	ch <- "작업 완료!" // 채널에 값 보내기
}

func main() {
	ch := make(chan string) // 채널 생성

	go worker(ch) // 고루틴 실행

	msg := <-ch // 채널에서 값 받기 (여기서 대기함)
	fmt.Println(msg)
}