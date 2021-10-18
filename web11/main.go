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

//func

func main() {
	user := User{Name: "Hong", Email: "guswns0218@naver.com", Age: 26}
	user2 := User{Name: "Park", Email: "1004brian@gmail.com", Age: 28}
	//users:= []Users{user,user2}
	tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge:{{.Age}}\n")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, user)
	tmpl.Execute(os.Stdout, user2)
}
