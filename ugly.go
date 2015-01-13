package main
// START OMIT
import ("fmt"
"net/http")

type User struct {
Name string `json:"name"`
Age int `json:"age"`
}

func main() {
	http.HandleFunc("/", Hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {fmt.Println("boom")}
}

func Hello(w    http.ResponseWriter,r *http.Request) {
	user := User{
Name: "Alice",Age: 2}
fmt.Fprintf(w, "%s, %d", user.Name, user.Age)
}
// END OMIT