FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o foundry

EXPOSE 8080

ENTRYPOINT ["/app/foundry"]

