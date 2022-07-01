package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

var (
	templateFile = template.Must(template.ParseFiles("templates/index.html"))
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleUpload(w, r)
		return
	}
	templateFile.ExecuteTemplate(w, "index.html", nil)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) //10 MB

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := "./uploaded_files/" + path.Base(fileHeader.Filename)
	dest, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", home)

	log.Println("Started server on port 5007")
	if err := http.ListenAndServe(":5007", nil); err != nil {
		log.Fatal(err)
	}
}
