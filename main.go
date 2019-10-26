package main

import "net/http"

func main() {
	http.HandleFunc("/", getHome)
	http.ListenAndServe(":5000", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testing"))
}
