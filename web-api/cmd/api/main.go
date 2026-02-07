package main

import(
	"fmt"
	"log"
	"net/http"
	"time"
)

//handlers

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Shapes API\n"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running\n"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Werner Sosa\n"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	w.Write([]byte(currentTime))
}

//coustom routes
func quote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Queen(Bohemian) 'Is this the real life? Is this just fantasy? Caught in a landslide, no escape from reality.'\n"))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/quote", quote)

	fmt.Println("Server running on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}