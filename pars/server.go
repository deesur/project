package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/lab1/", parse)
	http.ListenAndServe(":8080", nil)
}

func parse(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	str := r.Form.Get("str")

	fmt.Println("Ответ: ", pal(str))
}

func pal(st string) bool {
	result1 := []byte{}
	result2 := []byte{}
	for i := len(st) - 1; i >= 0; i-- {
		if st[i] >= 65 && st[i] <= 90 || st[i] >= 97 && st[i] <= 122 {
			result1 = append(result1, st[i])
		}
	}
	for i := 0; i <= len(st)-1; i++ {
		if st[i] >= 65 && st[i] <= 90 || st[i] >= 97 && st[i] <= 122 {
			result2 = append(result2, st[i])
		}
	}

	return string(result1) == string(result2)
}
