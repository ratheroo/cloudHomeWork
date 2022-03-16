FROM golang:1.17 AS builder
MAINTAINER GengDongzhe "ratherout@163.com"
ENV VERSION="g1.0"
LABEL multi.mylabel="homework from gengdongzhe"
#当前文件夹的文件复制到镜像里
COPY . .
#编译main文件并运行
RUN go build -o httpserver main.go

#二次构建，由生成的httpserver直接构建
FROM scratch
COPY --from=builder /httpserver/httpserver /

#镜像入口文件
ENTRYPOINT ["./httpserver"]
