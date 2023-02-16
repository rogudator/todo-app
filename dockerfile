FROM golang:1.18-alpine
WORKDIR /build
COPY go.mod . 
RUN go mod download
COPY . .
RUN go build -o /main cmd/main.go
ENTRYPOINT ["/main"]