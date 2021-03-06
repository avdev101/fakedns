FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --directory="go build -o ./faker cmd/faker/main.go" --command=./faker
