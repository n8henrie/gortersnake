FROM golang:1.16-buster as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main
RUN ls -A

FROM scratch
WORKDIR /root
COPY --from=builder /app/main /app/main
CMD ["/app/main"]