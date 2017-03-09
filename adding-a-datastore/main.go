package main

import (
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"context"
)

func main() {
	errc := make(chan error)
	ctx := context.Background()

	db, err := mgo.Dial("deals-db:27017")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	s := NewDealService(db)
	e := MakeEndpoints(s)

	fmt.Println("Deals service starting...")

	initData(db)

	handler := MakeHTTPHandler(ctx, e)
	errc <- http.ListenAndServe(":8888", handler)
}

// Initialize dummy data
func initData(s *mgo.Session) {
	c := s.DB("test").C("deals")
	err := c.Insert(&Deal{Id: 1, Name: "Buy 400 pairs, get one unmatched sock free!"},
		&Deal{Id: 2, Name: "Free shipping anywhere in the Andromeda Galaxy"})
	if err != nil {
		fmt.Printf("Error inserting records in database: %s\n", err.Error())
		panic(err)
	}
}
