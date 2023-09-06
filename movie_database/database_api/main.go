package main

import (
	"context"
	"database_api/data_control"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    new_router := chi.NewRouter()

    // for CRUD operations
    dc := data_control.NewDataControl(GetClient())
    // new_router.Get("/movie", dc.GetMovies)
    new_router.Get("/movie/{id}", dc.ReadMovieData)
    new_router.Post("/movie", dc.CreateMovieData)
    new_router.Delete("/movie/{id}", dc.DeleteMovieData)

    // for querying the data
    new_router.Get("/movie/query/{id}", data_control.IsAuthorized(dc.QueryMovieData))

    // starting the server
    http.ListenAndServe("localhost:8080", new_router)
}

func GetClient() *mongo.Client {
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        panic(err)
    }

    return client
}
