package main

import (
	"net/http"
)

func main() {

	port := "8080"

	http.HandleFunc("/", tempIOSaveHandler)

	http.ListenAndServe(":"+port, nil)
}

func tempIOSaveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	if r.Form.Get("temperature") != "" {
		save(r.Form.Get("temperature"), r.Form.Get("humidity"))
	}

	w.WriteHeader(http.StatusCreated)
}
