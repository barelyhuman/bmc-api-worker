package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/posts/", func(rw http.ResponseWriter, r *http.Request) {
		link := r.URL.Path[len("/posts/"):]
		fmt.Println("link", link)
		if link != "" {
			payload, _ := GetPostBySlug(link)
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(payload)
		} else {
			payload := GetAllPosts()
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(payload)
		}
	})

	port := ":3000"

	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	fmt.Println("Listening on port" + port)

	http.ListenAndServe(port, nil)
}

func cleanLink(link string) string {
	var cleaned string
	cleaned = link[len("https:/"):]
	fmt.Println(cleaned)
	return cleaned
}
