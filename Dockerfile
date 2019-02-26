FROM golang:1.11.4 as builder
WORKDIR /go/src/tempIO
COPY . .
RUN go get -u github.com/spf13/viper
RUN go get -u github.com/go-sql-driver/mysql
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tempIO .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/tempIO .
CMD ["./tempIO"]
