package main

import (
	"encoding/json"
	"net/http"
)

type Args struct {
	Arguments []string
}

type Home struct {
	Success string
}

type APIFunc func(w http.ResponseWriter, r *http.Request)

func returnJSON(f func() interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r  *http.Request) {
		resp := f()
		jsonObj, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonObj)
	}
}

func home() interface{} {
	homeResp := Home{"Test"}
	return homeResp
}

func main() {
	http.HandleFunc("/", returnJSON(home))
	http.ListenAndServe(":3000", nil)
}