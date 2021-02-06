package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	svcurl, _ := os.LookupEnv("svc-url")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get(svcurl)
		if err != nil {
			fmt.Println(err)
		}
		if response.Body != nil {
			defer response.Body.Close()
		}
		resp, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "response from local container : %s", string(resp))
	})
	http.ListenAndServe(":8080", nil)
}
