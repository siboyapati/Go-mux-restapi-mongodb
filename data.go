package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type Article struct {
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// // let's declare a global Articles array
// // that we can then populate in our main function
// // to simulate a database
// var Articles []Article

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnAllArticles")
// 	json.NewEncoder(w).Encode(Articles)
// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	fmt.Fprintf(w, "Key: "+key)
// }

// func handleRequests() {
// 	// creates a new instance of a mux router
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	// replace http.HandleFunc with myRouter.HandleFunc
// 	// myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/all", returnAllArticles)
// 	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
// 	// finally, instead of passing in nil, we want
// 	// to pass in our newly created router as the second
// 	// argument
// 	log.Fatal(http.ListenAndServe(":8080", myRouter))
// }

// func main() {
// 	d := Dog{
// 		Breed: "pug",
// 	}
// 	b, _ := json.Marshal(d)
// 	fmt.Println(string(b))

// 	fmt.Println("Rest API v2.0 - Mux Routers")
// 	Articles = []Article{
// 		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// 	}
// 	handleRequests()
// }
