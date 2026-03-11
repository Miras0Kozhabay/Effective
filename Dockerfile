FROM golang:1.25.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o app ./cmd

CMD ["./app"]