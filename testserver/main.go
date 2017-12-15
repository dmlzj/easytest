package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request : ", r.Method)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Request Error:", err)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		fmt.Println("Request :", string(body))
		w.Header().Set("Content-type", "application/json")
		w.Write(body)
	})
	http.ListenAndServe(":8008", nil)
}
