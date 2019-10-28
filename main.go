package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Articles []Article

var articles = Articles{
	Article{Title: "judul pertama", Description: "Deskripsi pertama"},
	Article{Title: "judul kedua", Description: "Deskripsi kedua"},
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/article", getArticle)
	http.HandleFunc("/post-article", postArticle)
	http.ListenAndServe(":5000", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testing"))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

// Handle form Body and response write
// func postArticle(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		body, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, "Can't Read Body", http.StatusInternalServerError)
// 		}
// 		w.Write([]byte(string(body)))

// 	} else {
// 		http.Error(w, "Invalid Request method", http.StatusMethodNotAllowed)
// 	}
// }

// Handle form json and response json
func postArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newArticle Article
		err := json.NewDecoder(r.Body).Decode(&newArticle)

		if err != nil {
			fmt.Printf("Ada Error", err)
		}

		articles = append(articles, newArticle)

		json.NewEncoder(w).Encode(articles)

	} else {
		http.Error(w, "Invalid Request method", http.StatusMethodNotAllowed)
	}
}
