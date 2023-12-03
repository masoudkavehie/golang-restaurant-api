FROM golang:1.16.3-alpine3.13

WORKDIR /app

COPY . .

RUN go get -d -v ./...

EXPOSE 8000

CMD ["./api"]