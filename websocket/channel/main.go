package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动一个goroutine，向ch1发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- fmt.Sprintf("数据 %d 来自 ch1", i)
			time.Sleep(1 * time.Second) // 模拟耗时操作
		}
		close(ch1) // 发送完毕后关闭channel
	}()

	// 启动另一个goroutine，向ch2发送数据
	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- fmt.Sprintf("数据 %d 来自 ch2", i)
			time.Sleep(1 * time.Second)
		}
		close(ch2) // 发送完毕后关闭channel
	}()

	// 使用select等待来自ch1或ch2的数据
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			// 如果ch1和ch2都为空，则打印"没有数据"
			fmt.Println("没有数据")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
