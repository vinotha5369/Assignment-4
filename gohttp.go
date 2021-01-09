package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request)
{
	fmt.Fprintf(w, "Web Application")
}

func about_handler(w http.ResponseWriter, r *http.Request)
{
	fmt.Fprintf(w, "Web Design Handler")
}

func main()
{
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about/",about_handler)
	http.ListenAndServer(":8000", nil)
}