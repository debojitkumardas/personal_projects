package data_model

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

/*
type Director struct {
    First_name string `json:"first_name" bson:"first_name"`
    Last_name string `json:"last_name" bson:"last_name"`
}
*/

type Movie struct {
    Id primitive.ObjectID `json:"id" bson:"_id"`
    ISBN string `json:"isbn" bson:"isbn"`
    Title string `json:"title" bson:"title"`
    Year string `json:"year" bson:"year"`  // year of release
    // Director *Director `json:"director" bson:"director"`
    // Theaters_only bool
    // Platforms string
}
