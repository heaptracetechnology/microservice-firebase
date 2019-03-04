FROM golang

RUN go get firebase.google.com/go

RUN go get github.com/NaySoftware/go-fcm

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/microservice-firebase

ADD . /go/src/github.com/heaptracetechnology/microservice-firebase

RUN go install github.com/heaptracetechnology/microservice-firebase

ENTRYPOINT microservice-firebase

EXPOSE 3000