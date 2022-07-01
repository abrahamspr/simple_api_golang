package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

func main(){
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	http.HandleFunc("/api/mahasiswa", user)
	fmt.Printf("Web Server berjalan di port 8085")
	http.ListenAndServe(":8085", nil)
}

type lepkom struct {
	Nama string `json:"nama_mahasiswa"`
	Kursus string `json:"kursus_mahasiswa"`
	Foto string `json:"foto_mahasiswa"`
}

var data_mahasiswa = []lepkom{
	{
		Nama : "Abraham",
		Kursus : "Golang Beginner",
		Foto : "img/komeng.jfif",
	},
	{
		Nama : "Komeng",
		Kursus : "Golang Intermediate",
		Foto : "img/komeng.jfif",
	},
}

func user(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")

	if r.Method == http.MethodGet {
		result, err := json.Marshal(data_mahasiswa)

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
}