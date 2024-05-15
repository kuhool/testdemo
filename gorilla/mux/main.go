package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	mapBooks map[string]*Book
	slcBooks []*Book
)

type Book struct {
	ISBN        string   `json:"isbn"`
	Name        string   `json:"name"`
	Authors     []string `json:"authors"`
	Press       string   `json:"press"`
	PublishedAt string   `json:"published_at"`
	Total       int      `json:"total"`
}

// Book结构体举例5条数据
var books = []*Book{
	{
		ISBN:        "9787111566417",
		Name:        "Go语言实战",
		Authors:     []string{"李明", "张华"},
		Press:       "机械工业出版社",
		PublishedAt: "2020-01-01",
		Total:       100,
	},
	{
		ISBN:        "9787111566424",
		Name:        "Go语言编程入门",
		Authors:     []string{"王五", "赵六"},
		Press:       "清华大学出版社",
		PublishedAt: "2020-02-01",
		Total:       100,
	},
	{
		ISBN:        "9787111566431",
		Name:        "Go语言高级编程",
		Authors:     []string{"孙七", "周八"},
		Press:       "人民邮电出版社",
		PublishedAt: "2020-03-01",
		Total:       100,
	},
	{
		ISBN:        "9787111566448",
		Name:        "Go语言从入门到精通",
		Authors:     []string{"吴九", "郑十"},
		Press:       "电子工业出版社",
		PublishedAt: "2020-04-01",
		Total:       100,
	},
	{
		ISBN:        "978-7111566455",
		Name:        "Go语言实战45讲",
		Authors:     []string{"王五", "赵六"},
		Press:       "清华大学出版社",
		PublishedAt: "2020-05-01",
		Total:       100,
	},
}

func init() {
	mapBooks = make(map[string]*Book)
	for _, book := range books {
		mapBooks[book.ISBN] = book
	}
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(books)
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	book, ok := mapBooks[mux.Vars(r)["isbn"]]
	if !ok {
		http.NotFound(w, r)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(book)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", BooksHandler)
	r.HandleFunc("/books/{isbn}", BookHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// github.com/GoAdminGroup/go-admin
