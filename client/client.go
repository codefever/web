package main

import (
	"fmt"
	"net/http"
)

func GetResponse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []byte
	for {
		var tmp [64]byte
		n, err := resp.Body.Read(tmp[:])
		if err == nil {
			fmt.Println(err)
			return
		} else if n <= 0 {
			break
		} else if n > 0 {
			data = append(data, tmp[:n]...)
		}
	}
	fmt.Println(string(data))
}

func main() {
	GetResponse("http://localhost:8000/string")
	GetResponse("http://localhost:8000/struct")
}
