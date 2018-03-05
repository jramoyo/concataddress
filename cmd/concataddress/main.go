package main

import (
	"log"
	"net/http"

	"github.com/jramoyo/concataddress/internal/app/concataddress"
)

func main() {
	http.HandleFunc("/", concataddress.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
