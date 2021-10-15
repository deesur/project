package main

import (
	"fmt"
	"net/http"
)

func privet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Zdarova chelik\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", privet)
	http.HandleFunc("/headers", headers)

	mux := http.NewServeMux()
	mux.HandleFunc("v1/teachers", tHandler)

	sHandler := stHandler{}
	mux.Handle("v1/students", sHandler)

	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":8090", nil)
}

func tHandler(res http.ResponseWriter, req *http.Request) {
	data := []byte("V1 of teacher's called")
	res.WriteHeader(200)
	res.Write(data)
}

type stHandler struct{}

func (h stHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("V1 of student's called")
	res.WriteHeader(200)
	res.Write(data)
}
