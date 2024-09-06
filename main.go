package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                 string
	age                  uint16
	money                int16
	avg_grade, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("Username is: %s. He is %d. ", u.name, u.age)
}

func (u *User) setNewName(new_name string) {
	u.name = new_name
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 21, 31, 4.5, 0.8}
	bob.setNewName("Martin")
	fmt.Fprintf(w, bob.getAllInfo())
}

func contact_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coooooooooontacts")
}

func friends_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Frieeeeeeeeeends")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contact_page)
	http.HandleFunc("/friends", friends_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
