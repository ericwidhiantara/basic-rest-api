package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

type Food struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

var (
	Foods []Food
)

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := "Hello From REST API using Golang"

	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error saat translate data", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func GetFoodHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(Foods) < 1 {
		http.Error(w, "Data makanan tidak ditemukan", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(Foods)
	if err != nil {
		http.Error(w, "Error saat translate data", http.StatusInternalServerError)
		return
	}
	w.Write(response)

}

func GetFoodByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	foodId := mux.Vars(r)["id"]
	intVar, _ := strconv.Atoi(foodId)
	for _, food := range Foods {
		if food.ID == intVar {
			response, err := json.Marshal(food)
			if err != nil {
				http.Error(w, "Error saat translate data", http.StatusNotFound)
				return
			}
			w.Write(response)
		}
	}
	// if foodId < 1 {
	// 	http.Error(w, "Data makanan tidak ditemukan", http.StatusNotFound)
	// 	return
	// }

}

func CreateFoodHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var temp Food

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Terdapat kesalahan pada input", http.StatusBadRequest)
		return
	}
	temp.ID = len(Foods) + 1
	Foods = append(Foods, temp)

	response, err := json.Marshal(temp)
	if err != nil {
		http.Error(w, "Error saat translate data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func main() {
	fmt.Println("Hello World")
	route := mux.NewRouter()
	route.HandleFunc("/", CheckHandler).Methods("GET")
	route.HandleFunc("/foods", GetFoodHandler).Methods("GET")
	route.HandleFunc("/foods/{id}", GetFoodByIdHandler).Methods("GET")
	route.HandleFunc("/foods", CreateFoodHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", route))
}
