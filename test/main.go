package main

import (
	"fmt"
	"os"
	//flagtest "test/flag"
)

// main.go

// 写一个检查demo
func main() {
	fmt.Println("hello world")
	mode := os.Getenv("GOVERSION")
	fmt.Println(mode)
	//flagtest.Test1()
	//var config string
	//flag.StringVar(&config, "c", "", "choose config file.")
	//flag.Parse()
	//fmt.Printf("参数:%s\n", config)
	////到這裏退出程序
	//os.Exit(1)
}
