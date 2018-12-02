FROM golang:1.8.3 as builder
WORKDIR /go/src/tempIO
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tempIO .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/tempIO .
CMD ["./tempIO"]
