FROM golang:1.5
ADD ./go-hello-world /app

ENV NODE_ENV production

EXPOSE 8080

CMD ["/app/go-hello-world"]