FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o fiber-prefork-server

EXPOSE 8080

CMD ["./fiber-prefork-server"]
