package data_mod

type Director struct {
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
}

type Movie struct {
    Id string `json:"id"`
    ISBN string `json:"isbn"`
    Title string `json:"title"`
    Director *Director `json:"director"`
}
