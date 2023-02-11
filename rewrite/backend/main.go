package main

import (
	"net/http"
)

func main() {
	finish := make(chan bool)

	app1 := http.NewServeMux()
	app1.HandleFunc("/ready", ready1)

	app2 := http.NewServeMux()
	app2.HandleFunc("/ready", ready2)

	go func() {
		http.ListenAndServe(":18081", app1)
	}()

	go func() {
		http.ListenAndServe(":18082", app2)
	}()

	<-finish
}

func ready1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("app1"))
}

func ready2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("app2"))
}
