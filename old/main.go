package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	Avg_grade, Happiness float64
	Hobbies              []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("Username is: %s. He is %d. ", u.Name, u.Age)
}

func (u *User) setNewName(new_name string) {
	u.Name = new_name
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 21, 31, 4.5, 0.8, []string{"Football", "Basket", "Volley"}}
	//bob.setNewName("Martin")
	//fmt.Fprintf(w, "<b>Main Text</b>")
	templ, _ := template.ParseFiles("templates/home.html")
	templ.Execute(w, bob)
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
