package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ServeHTTP takes http response writer and http request
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := "http://country.io/capital.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)

	// responseJSON, err := json.MarshalIndent(responseString, "", "    ")
	// if err != nil {
	// 	fmt.Println("there was a json error")

	// }
	// fmt.Println()
	// // fmt.Fprint(w, responseJSON)
	// os.Stdout.Write(responseJSON)

}
func main() {
	http.HandleFunc("/", ServeHTTP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
