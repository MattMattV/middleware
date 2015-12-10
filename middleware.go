package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var port string = ":9090"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	server := vars["server"]

	resp, err := http.Get("http://" + server + "/")

	if detectError(err) == false {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		detectError(err)
		w.Write(content)
	}

}

func detectError(err error) bool {

	if err != nil {
		log.Println(err)
		return true
	} else {
		return false
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/request/{server}", handleRequest).Methods("GET")

	http.ListenAndServe(port, router)
}
