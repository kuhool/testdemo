package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"reflect"
)

// https://pkg.go.dev/github.com/olivere/elastic
var client *elastic.Client

func init() {
	client = NewEsClient()
}
func NewEsClient() *elastic.Client {
	url := fmt.Sprintf("http://10.2.1.104:9200/")
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(url),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		log.Fatalln("Failed to create elastic client")
	}

	return client
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func GetValue() interface{} {
	return 1
}

type person struct {
	name string
}

func Test() {
	var arr = [5]int{1, 2, 3, 4, 5}                           // 定义一个长度为5的整型数组
	slice := arr[1:4]                                         // 创建一个切片，从索引1开始到索引3结束
	fmt.Println("Capacity of slice:", cap(slice), len(slice)) // 输出容量
	//create()
	//delete()
	//update()
	//gets()
	//query()
	//list(3, 1)
	//client, err := elastic.NewClient(elastic.SetURL("http://10.2.1.104:9200/"))
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//client1 := NewEsClient()
	//fmt.Println(client)
	//fmt.Println("connect to es success")
	//p1 := Person{Name: "lmh111", Age: 1811, Married: false}
	//put1, err := client.Index().
	//	Index("user").
	//	BodyJson(p1).
	//	Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//j, _ := json.Marshal(put1)
	//fmt.Println(string(j))

}

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

// 初始化

/*下面是简单的CURD*/

// 创建
func create() {

	//使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().
		Index("employee").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("employee").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().
		Index("employee").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)

}

// 删除
func delete() {

	res, err := client.Delete().Index("employee").
		//Type("employee").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

// 修改
func update() {
	res, err := client.Update().
		Index("employee").
		//Type("employee").
		Id("2").
		Doc(map[string]interface{}{"age": 881}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	j, e := json.Marshal(res)
	fmt.Println(string(j), e)
	fmt.Printf("update age %s\n", res.Result)

}

// 查找
func gets() {
	//通过id查找
	get1, err := client.Get().Index("employee").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
	j, e := json.Marshal(get1)
	fmt.Println(string(j), e)
}

// 搜索
func query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("employee").Do(context.Background())
	printEmployee(res, err)

	//字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res, err = client.Search("employee").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("employee").Query(q).Do(context.Background())
	printEmployee(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("employee").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests.keyword")
	res, err = client.Search("employee").Aggregation("all_interests", aggs).Do(context.Background())

	////分析 interests
	//aggs := elastic.NewTermsAggregation().Field("interests")
	//res, err = client.Search("employee").Aggregation("all_interests", aggs).Do(context.Background())
	fmt.Println("==============")
	j, e := json.Marshal(res)
	fmt.Println(string(j), e, err)
	printEmployee(res, err)

}

// 简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("employee").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	j, e := json.Marshal(res)
	fmt.Println(string(j), e, err)
	printEmployee(res, err)

}

// 打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}
