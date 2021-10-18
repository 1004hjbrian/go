package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
)

var rd *render.Render //전역변수 rd

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "junyoung", Email: "kfinnet123@gmail.com"}

	rd.JSON(w, http.StatusOK, user)
	//w.Header().Add("Content-type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(user)
	//fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) //user 의심
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, err)
	// rd.Text(w, http.StatusBadRequest, err.Erroro())  //rd.Text작성으로 한줄로 끝
	// 	return
	rd.HTML(w, http.StatusOK, "hello.tmpl", "HyeonJun")
	//tmpl.ExecuteTemplate(w, "hello.tmpl", "HyeonJun")
}

func main() {
	rd = render.New(render.Options{
		Directory:  "template",
		Extensions: []string{".html", ".tmpl"},
	})
	mux := pat.New() //gorilla.Mux.

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)
	//go get github/gorilla/pat 간편하면서도 강력한 데이터 전송Method 위치변경(뒤에서 앞으로) 기능제공
	http.ListenAndServe(":3000", mux)
}
