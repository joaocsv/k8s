package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/healtz", Healtz)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/secret", Secret)

	_ = http.ListenAndServe(":8080", nil)
}

func Healtz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
	if duration > 26 {
		w.WriteHeader(500)
		_, err := w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(200)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			return
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	version := os.Getenv("VERSION")

	_, _ = fmt.Fprintf(w, "Name: %s, version: %s", name, version)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/version/versions.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = fmt.Fprintf(w, "Versions: %s.", string(data))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Secret(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("USER")
	version := os.Getenv("PASSWORD")

	_, _ = fmt.Fprintf(w, "User: %s, password: %s", name, version)
}
