FROM golang:1.20-alpine

WORKDIR /app

COPY custom_apps/main.go .

#RUN go build -o app

EXPOSE 8080

RUN ["/usr/local/go/bin/go build main.go"]
CMD [./main]
