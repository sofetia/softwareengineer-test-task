# syntax=docker/dockerfile:1

FROM golang:1.21-bookworm AS build

WORKDIR /klaus-softwareengineering-test-task

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o klaus .

FROM alpine:latest

WORKDIR /klaus-softwareengineering-test-task

EXPOSE 9999

COPY database.db .

COPY --from=build /klaus-softwareengineering-test-task/klaus .

CMD ["./klaus"]
