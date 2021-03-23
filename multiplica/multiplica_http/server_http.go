package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func Multiplica(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	a := vars["n1"]
	b := vars["n2"]
	number1, _ := strconv.ParseInt(a, 10, 64)
	number2, _ := strconv.ParseInt(b, 10, 64)
	fmt.Fprint(response, number1*number2)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/multiplica/{n1}/{n2}", Multiplica).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

