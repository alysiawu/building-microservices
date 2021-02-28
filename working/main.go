package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nicholasjackson/building-microservices-youtube/tree/episode_2/product-api"
)

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Hello World")
	// 	d, err := ioutil.ReadAll(r.Body)
	// 	// log.Printf("Data %s\n", d)
	// 	if err != nil {
	// 		http.Error(rw, "Oops", http.StatusBadRequest)
	// 		// rw.WriteHeader(http.StatusBadRequest)
	// 		// rw.Write([]byte("oops"))
	// 		return
	// 	}

	// 	fmt.Fprintf(rw, "Hello %s", d)

	// })

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Goodby World")
	// })
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9092", nil)
}
