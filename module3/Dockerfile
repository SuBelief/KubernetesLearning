FROM golang:1.18 AS builder
ENV GO111MODULE=off \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,dircet"

WORKDIR /module3
COPY . .
RUN go build -o httpserver .
FROM scratch
COPY --from=builder /module3/httpserver .
EXPOSE 80
ENTRYPOINT ["./httpserver"]
