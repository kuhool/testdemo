package str

import (
	"fmt"
	"sort"
)

func Test1() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])
	}
}
func Test() {
	//fmt.Println(strings.ContainsAny("failure", "u & i"))
	//fmt.Println(strings.Count("谷歌中国", ""))
	//fmt.Println(strings.Count("google谷歌中国China", ""))
	//fmt.Println(strings.Count("China", ""))
	//fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))
	//fmt.Println(strings.HasPrefix("Gopher", "Go"))
	//fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	//
	//han := func(c rune) bool {
	//	return unicode.Is(unicode.Han, c) // 汉字
	//}
	//fmt.Println(strings.IndexFunc("Hello, world", han))
	//fmt.Println(strings.IndexFunc("Hello, 世界", han))

	var ce []student //定义一个切片类型的结构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	//a := make(map[string]interface{})
	a := map[string]interface{}{"name": "xiaoming"}

	fmt.Println(ce, a)
	demo(ce, a)
	fmt.Println(ce, a)
}

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student, s map[string]interface{}) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	s["name"] = 991
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}
