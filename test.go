package main

import (
    "fmt" // пакет для форматированного ввода вывода
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	 fmt.Printf("%T\n", "www")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}