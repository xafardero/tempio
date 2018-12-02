package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	port := "8080"

	http.HandleFunc("/", tempIOHandler)

	http.ListenAndServe(":"+port, nil)
}

func tempIOHandler(w http.ResponseWriter, r *http.Request) {
	tmpIOJson, error := parseTempIORequest(r)

	if error != nil {
		panic(error)
	}

	fileHandle, error := os.OpenFile("tempIO.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer fileHandle.Close()

	if error != nil {
		panic(error)
	}

	logWriter := bufio.NewWriter(fileHandle)

	fmt.Fprintln(logWriter, tmpIOJson)

	logWriter.Flush()

	w.WriteHeader(http.StatusCreated)
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
