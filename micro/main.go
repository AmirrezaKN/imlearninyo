package main

import (
	"fmt"
	"io/ioutil"
	_ "log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "oopsi", http.StatusBadRequest)
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("oopsi"))
		return
	}

	fmt.Fprintf(w, "Hello %s\n", data)
}
