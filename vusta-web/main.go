package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./web")))
	fmt.Println("http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
