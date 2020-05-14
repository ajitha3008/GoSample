// Learn Go in 3 hours
// Section 2, Video 1
// Our First Go Program
package main

import (
	"fmt"
	"net/http"
)

//code copied from https://github.com/PacktPublishing/Learn-Go-in-3-Hours/blob/master/Section2/v1/web_server.go
//http://localhost:8080/hello?name=ajitha
/*
	All Go programs start running from a function called `main` in a package called `main`
*/
func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}
