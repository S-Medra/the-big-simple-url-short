package main

import (
	"fmt"
	"net/http"
)

func main() {
	shortener := &URLShortener{
		urls: make(map[string]string),
	}

	//READ FORM
	http.HandleFunc("/", shortener.HandleForm)

	//SHORTEN
	http.HandleFunc("/shorten", shortener.HandleShorten)

	//REDIRECT
	http.HandleFunc("/short/", shortener.HandleRedirect)

	fmt.Println("URL Shortener is running on :" + port)
	http.ListenAndServe(":"+port, nil)
}
