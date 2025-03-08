# Dockerfile
FROM golang:alpine

WORKDIR "/testrestapi"

COPY . .

RUN go build -o ./main ./cmd/testrestapi/main.go

CMD ["./main"]

EXPOSE 8080


