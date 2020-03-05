package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	woeid := getWoeidMap()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		place, _ := r.URL.Query()["place"]
		fmt.Fprintf(w, woeid[strings.ToLower(strings.Join(place, ""))])
	})
	http.ListenAndServe(":8081", nil)
}

func getWoeidMap() map[string]string {

	result := make(map[string]string)

	file, _ := os.Open("/opt/files/woeid.txt")

	reader := bufio.NewReader(file)

	for {
		bytesRead, errs := reader.ReadString('\n')
		if errs != nil {
			break
		}
		bytesRead = strings.ToLower(bytesRead)
		strArr := strings.Split(bytesRead, ",")
		result[strings.TrimSpace(strArr[0])+","+strings.TrimSpace(strArr[1])] = strings.TrimSpace(strArr[2])
	}
	return result
}
