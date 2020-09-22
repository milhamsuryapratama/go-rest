package main

import (
	"fmt"
	"go-rest/kelas"
	"go-rest/siswa"
	"net/http"
)

func main() {
	http.HandleFunc("/kelas", kelas.GetAll)
	http.HandleFunc("/kelas/create", kelas.CreateKelas)
	http.HandleFunc("/kelas/delete", kelas.DeleteKelas)

	http.HandleFunc("/siswa", siswa.GetAll)
	http.HandleFunc("/siswa/create", siswa.CreateSiswa)
	http.HandleFunc("/siswa/delete", siswa.DeleteSiswa)
	fmt.Println("Server running")
	http.ListenAndServe(":9000", nil)
}
