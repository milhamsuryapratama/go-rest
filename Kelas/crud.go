package kelas

import (
	"database/sql"
	"encoding/json"
	"net/http"

	// mysql exported ...
	_ "github.com/go-sql-driver/mysql"
)

// Kelas is ...
type Kelas struct {
	ID        int
	NamaKelas string
}

func connectDb() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go-rest"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// GetAll is ...
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connectDb()

	k := Kelas{}

	if r.Method == "GET" {
		isKelas, err := db.Query("SELECT * FROM kelas")

		kelas := []Kelas{}

		for isKelas.Next() {
			var id int
			var namaKelas string

			err := isKelas.Scan(&id, &namaKelas)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				// return
			}

			k.ID = id
			k.NamaKelas = namaKelas

			kelas = append(kelas, k)
		}

		var res, error = json.Marshal(kelas)

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

// CreateKelas ...
func CreateKelas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connectDb()

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		_, error := db.Exec("INSERT INTO kelas (nama_kelas) VALUES (?)", r.Form.Get("nama_kelas"))

		if error != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "", http.StatusBadRequest)

	defer db.Close()
}

// DeleteKelas ...
func DeleteKelas(w http.ResponseWriter, r *http.Request) {
	db := connectDb()

	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM kelas WHERE id = ? ", id)

	if err != nil {
		// panic(err.Error())
		http.Error(w, "", http.StatusBadRequest)
	}

	defer db.Close()
	// w.Write()
}
