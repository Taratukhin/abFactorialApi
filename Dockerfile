FROM golang:1.18.1-alpine AS build

COPY . /go/src/app

WORKDIR /go/src/app/cmd/server

RUN CGO_ENABLED=0 go test -v ./../../...

RUN go build -o server server.go

FROM alpine

WORKDIR /

COPY --from=build /go/src/app/cmd/server/server /server

EXPOSE 8989

CMD ["/server"]