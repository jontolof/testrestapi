# Dockerfile
FROM golang:1.24-alpine3.21

# Installera nödvändiga paket
RUN apk update && apk add --no-cache git

# Installera Air
RUN go install github.com/air-verse/air@latest

# Definiera en arbetskatalog
WORKDIR "/app"

# Kopiera först endast modulfilerna för att cacha beroenden
COPY go.mod go.sum ./

RUN go mod download

# Kopera resten av filerna i projektet
COPY . .

# --nocache 
# RUN go build -o ./main ./cmd/testrestapi/main.go

# CMD ["./main"]
CMD ["air", "-c", ".air.toml"]

EXPOSE 8080