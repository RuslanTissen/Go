package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u User) getAllInfo() string { // string => what we want to get by return
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, .8, []string{"Football", "skate", "Dance"}}
	// fmt.Fprintf(w, `<h1>Main Text</h1>
	// <b>Main Text</b>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Page.")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":5000", nil)
}

func main() {
	// var bob User = ...
	// bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: .8}
	handleRequest()
}
