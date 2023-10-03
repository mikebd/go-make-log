FROM golang:1.21-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD main ./main
ADD makeLog ./makeLog

RUN go build -o go-make-log ./main

CMD ["./go-make-log"]
