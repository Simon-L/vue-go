package main

import (
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received query, writing current time in response")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"time\":\"" + time.Now().Format("15:04:05"+"\"}")))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
