package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func check(err interface{}) {
	if err != nil {
		log.Fatalln(err)
	}
}

const userid = "rosetta"
const password = "code"

func main() {
	http.HandleFunc("/p", func(rw http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(rw, "Hello, %q", html.EscapeString(r.URL.Path))
		log.Println(r.URL.Path)
		var payload map[string]interface{}
		payload = make(map[string]interface{})
		payload["token"] = "ABCDEFG"

		requestBody, err := json.Marshal(payload)

		timeout := time.Duration(5 * time.Second)
		client := http.Client{
			Timeout: timeout,
		}
		request, err := http.NewRequest("POST", "http://localhost:3005/dec", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.SetBasicAuth(userid, password)

		check(err)

		resp, err := client.Do(request)
		check(err)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		check(err)

		log.Printf("%s >> %s", resp.Status, string(body))
	})
	port := ":8888"
	log.Printf("Start Server at %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
