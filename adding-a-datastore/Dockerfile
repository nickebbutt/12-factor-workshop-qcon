FROM golang:1.6

COPY . /go/src/github.com/microservices-workshop/deals/

RUN go get github.com/gorilla/mux github.com/go-kit/kit/log github.com/go-kit/kit/endpoint github.com/go-kit/kit/transport/http gopkg.in/mgo.v2

RUN go install github.com/microservices-workshop/deals

ENTRYPOINT /go/bin/deals

EXPOSE 8888