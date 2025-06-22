package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go"))

}

func main() {
	fmt.Println("Server is running on port 8080")

	handler := MyHandler{}

	http.ListenAndServe(":8080", handler)
}
