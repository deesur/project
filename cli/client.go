package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	call("http://localhost:8080/employee", "POST")
}

func call(urlPath, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	data := url.Values{}
	data.Set("name", "John")
	data.Set("age", "18")
	data.Add("hobbies", "sports")
	data.Add("hobbies", "music")
	encodedData := data.Encode()
	fmt.Println(encodedData)
	req, err := http.NewRequest(method, urlPath, strings.NewReader(encodedData))
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	return nil
}
