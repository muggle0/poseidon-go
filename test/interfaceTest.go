package main

import (
	"fmt"
	"time"
)

// 接口

// 字符串

// 异常

// go 关键字

// 通道和缓冲

func slowFunc(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C
	}

}
func main() {
	messages := make(chan string)
	go slowFunc(messages)
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
