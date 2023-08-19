FROM golang:alpine as builder
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o /app/uptime ./cmd/uptime

FROM alpine:latest
WORKDIR /opt/bin
COPY --from=builder /app/uptime /opt/bin/uptime
CMD ["/opt/bin/uptime"]