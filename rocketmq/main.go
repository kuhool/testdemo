package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"log"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	// 配置 NameServer 地址
	nameServer := []string{"your-nameserver-address"}

	// 创建客户端实例
	client, err := rocketmq.NewClient(rocketmq.NewClientOptions())
	if err != nil {
		log.Fatalf("create rocketmq client failed, err: %v", err)
	}
	defer client.Shutdown()

	// 创建消费者实例
	consumerInstance := func(id int) {
		c, err := rocketmq.NewPushConsumer(
			consumer.WithNameServer(nameServer),
			consumer.WithGroupName(fmt.Sprintf("my_test_group_%d", id)),
		)
		if err != nil {
			log.Printf("create push consumer failed, err: %v", err)
			return
		}

		// 订阅主题
		err = c.Subscribe("my_test_topic", consumer.MessageSelector{},
			func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
				// 消费逻辑
				for _, msg := range msgs {
					fmt.Printf("Consumer %d: Received message - %s\n", id, string(msg.Body))
				}
				return consumer.ConsumeSuccess, nil
			},
		)
		if err != nil {
			log.Println(err)
			return
		}

		// 启动消费者
		if err = c.Start(); err != nil {
			log.Println(err)
			return
		}
		defer c.Shutdown()

		// 等待消费者运行，这里使用阻塞的方式，实际使用中可以根据需要进行调整
		select {}
	}

	// 定义并发消费者数量
	concurrentConsumers := 10
	for i := 0; i < concurrentConsumers; i++ {
		go consumerInstance(i)
	}

	// 主函数阻塞，等待 Goroutine 结束
	select {}
}
