package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "Hong", Email: "guswns0218@naver.com", Age: 26}
	user2 := User{Name: "Park", Email: "1004brian@gmail.com", Age: 41}
	//users:= []Users{user,user2}
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)
}
