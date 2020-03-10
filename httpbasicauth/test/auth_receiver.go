package main

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"
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

const userPass = "rosetta:code"
const unauth = http.StatusUnauthorized

func dec(rw http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)

	auth := req.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Basic ") {
		log.Print("Invalid authorization", auth)
		http.Error(rw, http.StatusText(unauth), unauth)
		return
	}
	up, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		log.Print("authorization decode error:", err)
		http.Error(rw, http.StatusText(unauth), unauth)
		return
	}
	if string(up) != userPass {
		log.Print("invalid username:password:", string(up))
		http.Error(rw, http.StatusText(unauth), unauth)
		return
	}

	io.WriteString(rw, "Successfully authed")
}

func main() {
	http.HandleFunc("/dec", dec)

	port := ":3005"
	log.Printf("Start Decoder Server2 at %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
