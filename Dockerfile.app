FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go build cmd/main.go

EXPOSE 8080

CMD ["./main"]