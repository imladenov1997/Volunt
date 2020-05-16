package main

import (
	"encoding/json"
	"fmt"
	"github.com/imladenov1997/volunt/app/response"
	"github.com/imladenov1997/volunt/components"
	"github.com/imladenov1997/volunt/graphql/queries"
	"net/http"

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
	bill := *components.NewTotalBill("Euro", 34.64)
	fmt.Printf("%+v", bill)


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