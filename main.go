package main 

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is a demo page")
	fmt.Println("homepage")
}

func handleRequests() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
