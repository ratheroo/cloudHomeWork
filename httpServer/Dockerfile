FROM alpine
ENV VERSION=1.0

LABEL multi.mylabel="homework from gengdongzhe"
COPY ./bin/httpserver main.go
RUN go build -o /bin/httpserver

FROM scratch
COPY --from=build /bin/httpserver /bin/httpserver
ENTRYPOINT ["/httpserver"]
