package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// func test(rw http.ResponseWriter, req *http.Request) {
// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(body)
// 	var t test_struct
// 	err = json.Unmarshal(body, &t)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(t)
// }

func main() {
	http.HandleFunc("/p", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello, %q", html.EscapeString(r.URL.Path))
		log.Println(r.URL.Path)
	})
	port := ":8888"
	log.Printf("Start Server at %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
