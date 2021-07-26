FROM golang:1.16-buster as builder
RUN mkdir /app
COPY ./*.go ./go.mod /app/
WORKDIR /app
RUN go mod download
RUN groupadd --system gortersnake && useradd --no-log-init --system --gid gortersnake gortersnake
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main

FROM scratch
WORKDIR /root
COPY --from=builder /etc/passwd /etc/group /etc/
COPY --from=builder /app/main /app/main
USER gortersnake
CMD ["/app/main"]
