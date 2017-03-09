FROM golang:1.6

COPY main.go /go/src/github.com/microservices-workshop/deals/main.go

RUN go get gopkg.in/mgo.v2
RUN go install github.com/microservices-workshop/deals

ENTRYPOINT /go/bin/deals

EXPOSE 8888
