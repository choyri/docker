FROM golang:alpine AS builder

ENV FRP_VERSION 0.29.1

RUN FRP_PATH=/go/src/github.com/fatedier/frp \
    && sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add gcc git libc-dev make \
    && git clone https://github.com/fatedier/frp -b "v$FRP_VERSION" $FRP_PATH \
    && cd $FRP_PATH \
    && sed -i "s/go build -o/go build --ldflags '-linkmode external -extldflags \"-static\"' -o/g" Makefile \
    && make \
    && mv bin/frps /usr/sbin/ \
    && mv conf/frps.ini /etc/

FROM scratch

COPY --from=builder /usr/sbin/frps /
COPY --from=builder /etc/frps.ini /

CMD ["/frps", "-c", "/frps.ini"]
