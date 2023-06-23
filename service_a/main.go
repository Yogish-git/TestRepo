package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from service_a main Yogish")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8081", nil)
}
