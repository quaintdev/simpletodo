package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/todo/add", handleAdd)

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
