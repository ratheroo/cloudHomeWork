FROM alpine
ENV VERSION=1.0
ADD ./bin/httpserver main.go
LABEL multi.label="homework from gengdongzhe"
RUN go build -o /bin/httpserver

FROM scratch
EXPOSE 80
ENTRYPOINT /httpserver