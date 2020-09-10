package main

import (
	"fmt"
	"go-rest/kelas"
	"net/http"
)

func main() {
	http.HandleFunc("/kelas", kelas.GetAll)
	http.HandleFunc("/kelas/create", kelas.CreateKelas)
	fmt.Println("Server running")
	http.ListenAndServe(":9000", nil)
}
