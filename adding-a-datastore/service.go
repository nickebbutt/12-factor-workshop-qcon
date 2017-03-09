package main

import (
	"context"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Service interface {
	GetDeal(ctx context.Context, id int) (Deal, error)
}

type Deal struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

func NewDealService(db *mgo.Session) Service {
	return &dealService{
		db: *db,
	}
}

type dealService struct {
	db mgo.Session
}

func (s *dealService) GetDeal(ctx context.Context, id int) (Deal, error) {
	c := s.db.DB("test").C("deals")
	r := Deal{}
	err := c.Find(bson.M{"id": id}).One(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}
