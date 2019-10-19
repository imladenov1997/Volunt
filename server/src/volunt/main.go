package main

import (
	"encoding/json"
	"net/http"
	"volunt/app/response"
	"volunt/graphql/queries"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

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
	homeResp := response.Home{"Test"}
	return homeResp
}

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queries.QueryType,
})

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.HandleFunc("/", returnJSON(home))
	http.Handle("/graph-ql", disableCors(h))
	http.ListenAndServe(":3000", nil)
}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
	if r.Method == "OPTIONS" {
	  w.Header().Set("Access-Control-Max-Age", "86400")
	  w.WriteHeader(http.StatusOK)
	  return
	}
	h.ServeHTTP(w, r)
   })
  }