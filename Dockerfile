FROM golang:1.22 as builder
LABEL authors="alex"

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk update && apk add pass
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main pm
ENTRYPOINT ["./pm"]
