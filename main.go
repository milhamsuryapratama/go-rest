package main

import (
	"fmt"
	"go-rest/kelas"
	"net/http"
)

func main() {
	http.HandleFunc("/kelas", kelas.GetAll)
	fmt.Println("Server running")
	http.ListenAndServe(":9000", nil)
}
