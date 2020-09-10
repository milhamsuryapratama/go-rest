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
	db := connectDb()

	if r.Method == "GET" {
		isKelas, err := db.Query("SELECT * FROM kelas")

		if err != nil {
			panic(err.Error())
		}

		k := Kelas{}
		kelas := []Kelas{}

		for isKelas.Next() {
			var id int
			var namaKelas string

			err := isKelas.Scan(&id, &namaKelas)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
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

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}

	http.Error(w, "", http.StatusBadRequest)

	defer db.Close()
}
