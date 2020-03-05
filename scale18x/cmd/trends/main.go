package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "github.com/TylerBrock/colorjson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/itsmurugappan/knative-examples/scale18x/statik"
	"github.com/rakyll/statik/fs"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(statikFS)))

	http.HandleFunc("/", getTwitterTrends)

	http.ListenAndServe(":8080", nil)

}

func getTwitterTrends(res http.ResponseWriter, r *http.Request) {
	place, _ := r.URL.Query()["place"]
	query := fmt.Sprintf("http://woeid.scale18x.svc.cluster.local?place=%s", url.QueryEscape(strings.Join(place, "")))
	response, err := http.Get(query)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	id, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if string(id) == "" {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("trend not found"))
		return
	}
	raw := callTwitter(string(id))

	var out bytes.Buffer
	json.Indent(&out, raw, "", "  ")

	fmt.Println(string(out.Bytes()))
	res.Write(out.Bytes())
}

func callTwitter(id string) []byte {
	token, _ := ioutil.ReadFile("/var/faas/secret/" + "token.txt")
	reqUrl := fmt.Sprintf("https://api.twitter.com/1.1/trends/place.json?id=%s", id)
	fmt.Println(reqUrl)
	fmt.Println(token)
	client := http.Client{}
	request, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Authorization", "Bearer "+string(token))
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	bytesOut, _ := ioutil.ReadAll(resp.Body)
	return bytesOut
}
