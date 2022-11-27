package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fetchCompliment())
}

var c []Compliments

func fetchCompliment() string {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(cap(c))
	return (c[idx].Compliment)
}

var data map[string]string

type Compliments struct {
	Id         string `json:"id"`
	Compliment string `json:"compliment"`
}

func main() {
	jsonFile, err := os.Open("compliments.json")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(b, &data)
	for key, element := range data {
		c = append(c, Compliments{Id: key, Compliment: element})
	}
	fetchCompliment()
	fmt.Println("Server Running")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))

}
