package data_control

import (
	"context"
	"database_api/data_model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type DataControl struct {
    client *mongo.Client
}

func NewDataControl(client *mongo.Client) *DataControl {
    return &DataControl{client}
}

/*
func (dc DataControl) GetMovies(w http.ResponseWriter, r *http.Request) {
    set_limit := options.Find().SetLimit(5)
    cursor, err := dc.client.Database("movie-databases").Collection("movies").Find(context.TODO(), bson.M{}, set_limit)
    if err != nil {
	fmt.Println(err)
	return
    }
    defer cursor.Close(context.Background())

    movie_data := []data_model.Movie{}
    if err := cursor.All(context.Background(), &movie_data); err != nil {
	fmt.Println(err)
	return
    }
}
*/

// *************************************** CRUD operations ************************************************
// ********************************************************************************************************
func (dc DataControl) CreateMovieData(w http.ResponseWriter, r *http.Request) {
    movie_data := data_model.Movie{}

    if err := json.NewDecoder(r.Body).Decode(&movie_data); err != nil {
	fmt.Println(err)
	return
    }

    movie_data.Id = primitive.NewObjectID()
    if _, err := dc.client.Database("movie-database").Collection("movies").InsertOne(context.TODO(), movie_data); err != nil {
	fmt.Println("Failed to insert movie data in the database. Error:", err)
	return
    }

    movie_data_json, err := json.Marshal(movie_data)
    if err != nil {
	fmt.Println("Failed to marshall movie_data to json. Error:", err)
    }

    w.Header().Set("Context-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s\n", movie_data_json)
}

func (dc DataControl) ReadMovieData(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    oid, err := primitive.ObjectIDFromHex((id))
    if err != nil {
	fmt.Println("Error converting object-id string. Error:", err)
	w.WriteHeader(http.StatusNotFound)
	return
    }

    movie_data := data_model.Movie{}

    if err := dc.client.Database("movie-database").Collection("movies").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&movie_data); err != nil {
	fmt.Println("Failed to retrieve movie data from the database. Error:", err)
	return
    }

    movie_data_json, err := json.Marshal(movie_data)
    if err != nil {
	fmt.Println("Failed to marshall movie_data to json. Error:", err)
    }

    w.Header().Set("Context-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%s\n", movie_data_json)
}

func (dc DataControl) UpdateMovieData(w http.ResponseWriter, r* http.Request) {
    //
}

func (dc DataControl) DeleteMovieData(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
	fmt.Println("Error converting object-id string. Error:", err)
	w.WriteHeader(http.StatusNotFound)
	return
    }

    if _, err := dc.client.Database("movie-database").Collection("movies").DeleteOne(context.TODO(), bson.M{"_id": oid}); err != nil {
	fmt.Println("Falied to delete movie data from database. Error:", err)
	return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Delete movie ", oid, '\n')
}
// ********************************************************************************************************

// *************************************** .. operations ************************************************
// ********************************************************************************************************
var SigningKey = []byte(os.Getenv("SECRET_SIGNING_KEY"))

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Header["Token"] != nil {
	    token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
		    return nil, fmt.Errorf("Invalid signing method")
		}
		return SigningKey, nil
	    })

	    if err != nil {
		fmt.Fprintf(w, err.Error())
	    }

	    if token.Valid {
		endpoint(w, r)
	    }
	} else {
	    fmt.Fprintf(w, "No authorization token provided")
	}
    })
}

func (dc DataControl) QueryMovieData(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    oid, err := primitive.ObjectIDFromHex((id))
    if err != nil {
	fmt.Println("Error converting object-id string. Error:", err)
	w.WriteHeader(http.StatusNotFound)
	return
    }

    movie_data := data_model.Movie{}

    if err := dc.client.Database("movie-database").Collection("movies").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&movie_data); err != nil {
	fmt.Println("Failed to retrieve movie data from the database. Error:", err)
	return
    }

    movie_data_json, err := json.Marshal(movie_data)
    if err != nil {
	fmt.Println("Failed to marshall movie_data to json. Error:", err)
    }

    w.Header().Set("Context-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%s\n", movie_data_json)
}
