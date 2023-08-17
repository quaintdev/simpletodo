package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	tItems = make(map[string]*Todo)

	http.HandleFunc("/todo/add", handleAddAdvanced)
	http.HandleFunc("/todo/checked", handleTodoChecked)
	http.HandleFunc("/", handleTodoList)
	http.ListenAndServe(":8080", nil)

}

func handleTodoList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html", "./static/todoitem.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	var todos []Todo
	for _, v := range tItems {
		todos = append(todos, *v)
	}
	err = tmpl.Execute(w, todos)
	if err != nil {
		log.Println(err)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Todo struct {
	Id        string
	Item      string
	IsChecked bool
}

func handleTodoChecked(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	for k := range r.URL.Query() {
		if t, ok := tItems[k]; ok {
			t.IsChecked = true
		}
	}
	w.WriteHeader(http.StatusOK)
}

var tItems map[string]*Todo

func handleAddAdvanced(w http.ResponseWriter, r *http.Request) {
	todo := Todo{
		Id:   RandStringRunes(3),
		Item: r.URL.Query()["contains"][0],
	}
	tItems[todo.Id] = &todo
	tmpl, err := template.ParseFiles("./static/todoitem.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "todo", todo)
	if err != nil {
		log.Println(err)
	}
}
