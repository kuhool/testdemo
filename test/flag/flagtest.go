package flag

import (
	"flag"
	"fmt"
)

func Test1() {
	// 定义一个命令行参数
	flag.Int("test.v", 0, "a test variable")

	// 解析命令行参数
	flag.Parse()

	// 查找名为 "test.v" 的参数
	flagVar := flag.Lookup("test.v")
	if flagVar != nil {
		fmt.Println("Found flag:", flagVar.Name)
		fmt.Println("Current value:", flagVar.Value.String())
	} else {
		fmt.Println("Flag not found")
	}
}
