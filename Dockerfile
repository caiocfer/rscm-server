FROM --platform=arm64 golang:1.18-alpine3.15 AS build
WORKDIR /app
COPY . .
RUN GOARCH=arm64 go build -o main main.go
FROM alpine:3.15
WORKDIR /app

COPY --from=build  /app/main .

EXPOSE 8080
CMD ["/app/main"]