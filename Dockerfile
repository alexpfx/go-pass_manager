FROM golang:1.22 as builder
LABEL authors="alex"

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM scratch
COPY --from=builder /app/main /pm
ENTRYPOINT ["/pm"]