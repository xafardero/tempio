package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	port := "8080"

	http.HandleFunc("/", tempIOSaveHandler)
	http.HandleFunc("/now", tempIONowHandler)

	http.ListenAndServe(":"+port, nil)
}

func tempIOSaveHandler(w http.ResponseWriter, r *http.Request) {
	tmpIOJson, error := parseTempIORequest(r)

	if error != nil {
		panic(error)
	}

	save("salon", tmpIOJson)

	w.WriteHeader(http.StatusCreated)
}

func tempIONowHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	value := read("salon", currentTime.Format("2006-01-02_15:04"))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, value)
}

func parseTempIORequest(r *http.Request) (string, error) {
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	tmpIO := tempIO{
		DateTime:    time.Now(),
		Temperature: r.Form.Get("temperature"),
		Humidity:    r.Form.Get("humidity"),
	}

	return tmpIO.json()
}
