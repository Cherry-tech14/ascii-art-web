package ascii

import (
	"html/template"
	"net/http"
)

type Output struct {
	Result string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	data := Output{}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		banner := r.FormValue("font")
		if text == "" || banner == "" {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		data.Result = LoadBanner(text, banner)
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
