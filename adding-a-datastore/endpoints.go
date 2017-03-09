package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetDealEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetDealEndpoint: MakeGetDealEndpoint(s),
	}
}

func MakeGetDealEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getDealRequest)
		p, e := s.GetDeal(ctx, req.ID)
		return getDealResponse{Id: p.Id, Name: p.Name}, e
	}
}

type getDealRequest struct {
	ID int
}

type getDealResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
