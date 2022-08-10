package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Company struct {
	Name string `json:"name"`
}

type Employee struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	LastName string   `json:"lastName"`
	Age      string   `json:"age"`
	Salary   string   `json:"salary"`
	Company  *Company `json:"company"`
}

var employees []Employee

func main() {
	router := mux.NewRouter()

	employees = append(employees, Employee{
		ID:       "1",
		Name:     "Yaroslav",
		LastName: "Nykolyshyn",
		Age:      "19",
		Salary:   "200$",
		Company:  &Company{Name: "Google"},
	})
	employees = append(employees, Employee{
		ID:       "2",
		Name:     "Ruslan",
		LastName: "Nykolyshyn",
		Age:      "19",
		Salary:   "200$",
		Company:  &Company{Name: "Phokal"},
	})
	router.HandleFunc("/employee", getEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/employee", createEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", updateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:i], employees[i+1])
			var emp Employee
			_ = json.NewDecoder(r.Body).Decode(&emp)
			emp.ID = params["id"]
			employees = append(employees, emp)
			json.NewEncoder(w).Encode(emp)
		}
	}
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	emp.ID = strconv.Itoa(rand.Intn(10000000))
	employees = append(employees, emp)
	json.NewEncoder(w).Encode(emp)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range employees {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range employees {

		if item.ID == params["id"] {
			employees = append(employees[:i], employees[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
