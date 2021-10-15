package main

import (
	"fmt"
	"net/http"
)

type employee struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

func main() {
	createEmployeeHandler := http.HandlerFunc(createEmployee)
	http.Handle("/employee", createEmployeeHandler)
	http.ListenAndServe(":8080", nil)
}

func createEmployee(res http.ResponseWriter, req *http.Request) {
	headerContentType := req.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	req.ParseForm()
	fmt.Println("request.Form::")
	for key, value := range req.Form {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}
	fmt.Println("\nrequest.PostForm::")
	for key, value := range req.Form {
		fmt.Printf("Key:%s, Value:%S\n", key, value)
	}
	fmt.Printf("\nName in Form\n", req.Form["name"])
	fmt.Printf("\nName in PostForm\n", req.PostForm["name"])
	fmt.Printf("\nHobbies in FormValue\n", req.FormValue["hobbies"])

	res.WriteHeader(200)
	return
}
