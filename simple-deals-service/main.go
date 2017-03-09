package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// Stores all the sweet deals
var deals []Deal

// Entrypoint of our (micro)sevice
func main() {
	fmt.Println("Deals service started...")

	initData()

	http.HandleFunc("/deals", dealsHandler)

	http.ListenAndServe(":8888", nil)
}

// Handle deal requests
func dealsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request receiveed")
	idStr := r.FormValue("id")

	var deal Deal
	if idStr == "" {
		fmt.Println("No Id passed in")
		w.WriteHeader(400)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Oops: ", err.Error())
		w.WriteHeader(500)
		return
	}
	deal, err = fetchDeal(id)
	if err != nil {
		fmt.Println("Oops: ", err.Error())
		w.WriteHeader(404)
		return
	}

	err = json.NewEncoder(w).Encode(deal)
	if err != nil {
		fmt.Println("Oops: ", err.Error())
		w.WriteHeader(500)
	}
	fmt.Println("success.")
}

// Retrieve Deal by ID
func fetchDeal(id int) (Deal, error) {
	if id > len(deals) || id <= 0 {
		return Deal{}, errors.New("Invalid Id")
	}
	return deals[id-1], nil
}

// Initialize dummy data
func initData() {
	deals = append(deals, Deal{Id: 1, Name: "Buy 400 pairs, get one unmatched sock free!"})
	deals = append(deals, Deal{Id: 2, Name: "Free shipping anywhere in the Andromeda Galaxy"})
}

type Deal struct {
	Id   int
	Name string
}
