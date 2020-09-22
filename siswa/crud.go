package siswa

import (
	"encoding/json"
	database "go-rest/Database"
	"go-rest/kelas"
	"go-rest/response"
	"net/http"

	// mysql exported ...
	_ "github.com/go-sql-driver/mysql"
)

// Siswa is ...
type Siswa struct {
	ID     int
	Nama   string
	Alamat string
	Jk     string
	Kelas  kelas.Kelas
}

// func connectDb() (db *sql.DB) {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "go-rest"
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// GetAll is ...
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.ConnectDb()

	s := Siswa{}

	if r.Method == "GET" {
		isSiswa, err := db.Query("SELECT siswa.id, siswa.nama, siswa.jk, siswa.alamat, kelas.nama_kelas FROM siswa JOIN kelas ON kelas.id = siswa.kelas_id")

		siswa := []Siswa{}

		for isSiswa.Next() {
			var id int
			var nama, jk, alamat, namaKelas string

			err := isSiswa.Scan(&id, &nama, &jk, &alamat, &namaKelas)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				// return
			}

			s.ID = id
			s.Nama = nama
			s.Jk = jk
			s.Alamat = alamat
			// k.ID = kelasID
			// k.NamaKelas = namaKelas
			s.Kelas.ID = id
			s.Kelas.NamaKelas = namaKelas

			siswa = append(siswa, s)
		}

		var res, error = json.Marshal(siswa)

		if error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
	}

	http.Error(w, "", http.StatusBadRequest)

	defer db.Close()
	return
}

// CreateSiswa ...
func CreateSiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.ConnectDb()

	res := response.Response{}

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		_, error := db.Exec("INSERT INTO siswa (nama, jk, alamat, kelas_id) VALUES (?, ?, ?, ?)", r.Form.Get("nama"), r.Form.Get("jk"), r.Form.Get("alamat"), r.Form.Get("kelas_id"))

		if error != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		res = response.Response{
			Status: 200,
			Pesan:  "Sukses!!!",
		}

		var result, er = json.Marshal(res)
		if er != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)

	defer db.Close()
}

// DeleteSiswa ...
func DeleteSiswa(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDb()

	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM siswa WHERE id = ? ", id)

	if err != nil {
		// panic(err.Error())
		http.Error(w, "", http.StatusBadRequest)
	}

	res := response.Response{
		Status: 200,
		Pesan:  "Sukses!!!",
	}

	var result, er = json.Marshal(res)
	if er != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)

	defer db.Close()
}
