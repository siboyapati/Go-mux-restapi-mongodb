package main

import (
	"fmt"
	"log"
	"net/http"
	. "test/routers"
)

func init() {
	fmt.Println("hello world from main function!!!!")
}

func main() {
	router := NewRouter()
	log.Printf("[Info] Server listening on the port : 3000")
	server := http.ListenAndServe(":3000", router)
	log.Fatal(server)

}
