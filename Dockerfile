FROM golang:1.17-alpine3.15

WORKDIR /app

COPY custom_apps/main.go .

#RUN go build -o app

EXPOSE 8080

CMD ["go run main.go"]
