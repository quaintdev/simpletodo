package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/todo/add", handleAddAdvanced)
	http.HandleFunc("/todo/checked", handleTodoChecked)
	http.ListenAndServe(":8080", nil)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	id := RandStringRunes(3)
	item := r.URL.Query()["contains"][0]
	html := fmt.Sprintf(`
	<li>
		<input type="checkbox" id="%s">
    	<label for="%s">%s</label>
	</li>`, id, id, item)

	w.Write([]byte(html))
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
}

var tItems []Todo

func handleAddAdvanced(w http.ResponseWriter, r *http.Request) {
	todo := Todo{
		Id:   RandStringRunes(3),
		Item: r.URL.Query()["contains"][0],
	}
	tItems = append(tItems, todo)
	tmpl, err := template.ParseFiles("./static/todoitem.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, todo)
}
