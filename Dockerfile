FROM golang:1.17-alpine3.15

WORKDIR /app

COPY main.go .

RUN go build -o app

EXPOSE 8080

CMD ["./app"]
