FROM golang:1.24.2-alpine

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]