FROM golang:1.14.1-alpine3.11
WORKDIR /go/src/github.com/msiebhur/prometheus-mdns-sd/
COPY mdns.go go.mod go.sum ./
RUN GO111MODULES=on CGO_ENABLED=0 GOOS=linux go build -a -o mdns .

FROM alpine:3.11
WORKDIR /root/
COPY --from=0 /go/src/github.com/msiebhur/prometheus-mdns-sd/mdns .
ENTRYPOINT ["./mdns"]
