FROM golang:alpine AS builder

ARG BRANCH
ARG ON_HUB

RUN set -x \
    && if [ -z "${BRANCH}" ]; then BRANCH=master; else BRANCH=v${BRANCH}; fi \
    && if [ -z "${ON_HUB}" ]; then \
        sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories; \
        go env -w GOPROXY="https://goproxy.cn,direct" GOSUMDB="sum.golang.google.cn"; \
    fi \
    && apk --no-cache add git gcc libc-dev tzdata \
    && git clone -b "${BRANCH}" --depth 1 https://github.com/fatedier/frp && cd frp \
    && go build -ldflags '-s -w -linkmode external -extldflags "-static"' -o bin/frps ./cmd/frps

FROM scratch

LABEL maintainer="chotow <i@choyri.com>"

COPY --from=builder /go/frp/bin/frps /
COPY --from=builder /go/frp/conf/frps.ini /
COPY --from=builder /usr/share/zoneinfo/${TZ:-Asia/Shanghai} /etc/localtime

CMD ["/frps", "-c", "/frps.ini"]
