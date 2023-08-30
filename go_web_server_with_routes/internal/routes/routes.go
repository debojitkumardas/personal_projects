package routes

import (
    "fmt"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the homepage!")
}

func APIDataHandler(w http.ResponseWriter, r *http.Request) {
    data := "Some data from the API"
    fmt.Fprintln(w, data)
}

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/", IndexHandler)
    mux.HandleFunc("/api/data", APIDataHandler)

    return mux
}
