FROM golang:1.18.0-alpine3.15
WORKDIR src/
COPY go.mod go.sum ./
COPY server.go helpers.go database.json ./
RUN go build -v -o /usr/local/bin/app ./...
CMD ["app"]