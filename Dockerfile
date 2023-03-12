FROM golang:buster as build
WORKDIR /go/src/github.com/pastukhov/prometheus-mdns-sd/
COPY go.mod go.sum ./
RUN go mod download
COPY mdns.go ./
RUN GO111MODULES=on CGO_ENABLED=0 GOOS=linux go build -a -o mdns .

FROM debian:buster-slim
WORKDIR /cmd/
COPY --from=build /go/src/github.com/pastukhov/prometheus-mdns-sd/mdns .
ENTRYPOINT ["./mdns"]
