package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SigninKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT(user_id string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    claims["sub"] = user_id
    claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

    token_string, err := token.SignedString(SigninKey)

    if err != nil {
	_ = fmt.Errorf("Something went wrong: %s\n", err.Error())
	return "", err
    }

    return token_string, err
}

func GenerateToken(w http.ResponseWriter, _ *http.Request) {
    user_id := "johndoe"
    // fmt.Print("Enter userid: ")
    // fmt.Scanln(&user_id)

    valid_token, err := GetJWT(user_id)
    if err != nil {
        fmt.Println("Failed to generate token")
    }
    // fmt.Println(valid_token)
    fmt.Fprintf(w, string(valid_token))
}

func main() {
    http.HandleFunc("/movie/token", GenerateToken)
    log.Fatal(http.ListenAndServe("localhost:9001", nil))
}
