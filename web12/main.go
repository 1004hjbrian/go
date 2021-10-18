package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "junyoung", Email: "kfinnet123@gmail.com"}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) //user 의심
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func main() {
	mux := mux.NewRouter() //Gorilla Mux 사용 , http.NewServeMux대신 Gorilla Mux.NewRouter() 사용

	mux.HandleFunc("/users", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users", addUserHandler).Methods("POST")
	//go get github/gorilla/pat 간편하면서도 강력한 데이터 전송Method 위치변경(뒤에서 앞으로) 기능제공
	http.ListenAndServe(":3000", mux)
}
