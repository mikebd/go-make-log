FROM golang:1.21-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD config ./config
ADD main ./main
ADD makeLog ./makeLog

RUN go build -o go-make-log ./main

ENTRYPOINT ["./go-make-log"]
