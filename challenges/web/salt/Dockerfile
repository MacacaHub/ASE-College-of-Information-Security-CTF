FROM golang:1.18-alpine AS builder
RUN apk update
RUN apk add --no-cache build-base

WORKDIR /app
COPY ./app/go.mod .
COPY ./app/go.sum .
RUN  go mod download
COPY ./app/app.go .
RUN CGO_ENABLED=1 go build -o app app.go

FROM alpine:3.15
WORKDIR /app/
COPY --from=builder /app/app ./

EXPOSE 3000
CMD ["./app"]
