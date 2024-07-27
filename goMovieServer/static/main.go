package main

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director`
}

type Director struct {
	Firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

var movie []Movie

func main() {
	fileServer

}
