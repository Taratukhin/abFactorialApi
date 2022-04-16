FROM golang:1.16

COPY . /go/src/app

WORKDIR /go/src/app/cmd/server

RUN go build -o server server.go

EXPOSE 8989

CMD ["./server"]
