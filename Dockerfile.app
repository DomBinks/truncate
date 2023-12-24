FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go build cmd/main.go

EXPOSE 443

CMD ["./main"]