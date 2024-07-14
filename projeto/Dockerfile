FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add libc6-compat

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]
