package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"taskhiska/entity"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "Task Progate",
	// 	"content": "Create Task",
	// }

	// data := entity.Task{ID: 1, Name: "Assigntment", Deadline: "Tanggal", Action: "Ya/Tidak"}

	data := []entity.Task{
		{ID: 1, Name: "Danu", Deadline: "15 Maret 2022", Action: "Ya"},
		{ID: 2, Name: "Febry", Deadline: "16 Maret 2022", Action: "Tidak"},
		{ID: 3, Name: "Ibnu", Deadline: "17 Maret 2022", Action: "Tidak"},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Hello World, saya sedang belajar Golang Web")))
}

func TugasprogateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Tugas Progate Go-Language")))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//fmt.Fprintf(w, "Product page : %d", idNumb)
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}

}
