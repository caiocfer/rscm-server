FROM --platform=arm64 golang:1.18-alpine3.15 AS build
WORKDIR /app
COPY . .
RUN GOARCH=arm64 go build -o main main.go
FROM alpine:3.15
WORKDIR /app

COPY --from=build  /app/main .

ENV DB_USER=root
ENV DB_PASS=root
ENV DB_NAME=rscm
ENV DB_PROTOCOL="tcp(140.238.182.164:32415)"

ENV API_PORT=5001



EXPOSE 8080
CMD ["/app/main"]