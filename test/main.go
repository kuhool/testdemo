package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"net/http"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type Server interface {
	http.Handler
}

func defer_call() {
	defer func() {
		fmt.Println("打印前")
		panic("触发异常")
	}()
	defer func() {
		fmt.Println("打印中")
		panic("触发异常")
	}()
	defer func() {
		fmt.Println("打印后")
		panic("触发异常")
	}()
	panic("触发异常")
}
func defer_call1() {
	defer func() {
		fmt.Println("打印前")

	}()
	defer func() {
		fmt.Println("打印中")

	}()
	defer func() {
		fmt.Println("打印后")

	}()
	panic("触发异常")
}

// 代码中只有defer和panic语句，没有recover()，执行顺序是先defer，然后是panic
// 代码中有defer和panic语句，有recover()，执行顺序是先defer，然后是panic，然后是recover(),
// 多个panic执行 recover最后一个
// 函数中有panic 函数中有defer,defer中的panic，两个都会执行
func exampleFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Starting example function.")

	defer fmt.Println("Defer statement 3.")
	defer fmt.Println("Defer statement 4.")

	panic("Panic from example function!")
	// 这行代码也不会被执行，因为上面已经触发了 panic。
	fmt.Println("This line will also never be reached.")
}

func runExample() {
	fmt.Println("Running example function...")
	exampleFunction()
	fmt.Println("Finished running example function.")
}
func test() {
	defer func() {
		fmt.Println(recover())
		fmt.Println("Running example function...1")
	}()

	defer func() {
		panic("defer panic")
		fmt.Println("Running example function...2")
	}()

	panic("test panic")
}
func forx() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "===>", *v)
	}
}

func XX() (x, y int) {
	return 1, 2
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}
func Test(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")
}

// var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// var slice0 []int = arr[2:8]
// var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
// var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
// var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
// var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
func main() {
	Test(0)

	//fmt.Printf("全局变量：arr %v\n", arr)
	//fmt.Printf("全局变量：slice0 %v\n", slice0)
	//fmt.Println(cap(slice0))
	//fmt.Printf("全局变量：slice1 %v\n", slice1)
	//fmt.Println(cap(slice1))
	//fmt.Printf("全局变量：slice2 %v\n", slice2)
	//fmt.Println(cap(slice2))
	//fmt.Printf("全局变量：slice3 %v\n", slice3)
	//fmt.Println(cap(slice3))
	//fmt.Printf("全局变量：slice4 %v\n", slice4)
	//fmt.Println(cap(slice4))
	//fmt.Printf("-----------------------------------\n")

	//es.Test()
	//forx()
	//runExample()
	//test()

	//defer_call()
	//var h Server
	//http.ListenAndServe(":8080", h)
	//str.Test()
	//str.Test1()
	//a := 1
	//b := 2
	//go func() {
	//	fmt.Println(a, b)
	//}()
	//
	//a = 3
	//
	//time.Sleep(1 * time.Second)
	//startOneProducer()

	//startOneConsumer()
}
func startOneConsumer() {

	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"10.2.1.89:9876"}),
		consumer.WithGroupName("my_test_group"),
	)
	err := c.Subscribe("my_test_topic",
		consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				fmt.Printf("==========获取到的值为: %s\n", msg.Body)
			}
			return consumer.ConsumeSuccess, nil
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = c.Start()
	defer c.Shutdown()
	time.Sleep(10000 * time.Second) //阻塞10s  等待消费 可以修改阻塞时间，下边的发送也可以使用这个接收

}
func startOneProducer() {

	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"10.2.1.89:9876"}))
	if err != nil {
		panic(err)
	}
	err = p.Start()
	defer p.Shutdown()
	if err != nil {
		panic(err)
	}

	res, err := p.SendSync(context.Background(), &primitive.Message{
		Topic: "my_test_topic",
		Body:  []byte("Hello RocketMQ!"),
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("发送成功", res.String())
	}
}
