FROM golang:1.17 AS builder
MAINTAINER GengDongzhe "ratherout@163.com"
ENV VERSION="g1.0"
ENV GO111MODULE=on
ENV CGO_ENABLED=0
LABEL multi.mylabel="homework from gengdongzhe"
#当前文件夹的文件复制到镜像里
WORKDIR /httpServer/
COPY . .
#编译main文件并运行
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

#二次构建，由生成的httpserver直接构建
FROM scratch
COPY --from=builder /httpserver/httpserver /
WORKDIR /httpServer/
#镜像入口文件
ENTRYPOINT ["./httpserver"]
