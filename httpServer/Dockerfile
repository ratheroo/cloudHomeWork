FROM golang:1.17 AS builder
MAINTAINER GengDongzhe "ratherout@163.com"
ENV VERSION="g1.0"
LABEL multi.mylabel="homework from gengdongzhe"
COPY . .
RUN go build -o httpserver main.go

FROM scratch
COPY --from=builder ./httpserver ./httpserver
ENTRYPOINT ["./httpserver"]
