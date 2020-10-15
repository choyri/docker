ARG IMAGE_TAG=alpine

FROM nginx:${IMAGE_TAG} AS builder

ARG VERSION

RUN set -x \
    && if [ "$VERSION" == "" ] ; then echo 'arg VERSION does not exist'; exit 1; fi \
    \
    && if [ -z "${ON_HUB}" ]; then \
        sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories; \
    fi \
    && apk update \
    && apk add --virtual .build-deps \
        gcc \
        libc-dev \
        make \
        pcre-dev \
        zlib-dev \
        git \
        openssl \
    \
    && tempDir="$(mktemp -d)" \
    && chown nobody:nobody $tempDir \
    && cd ${tempDir} \
    \
    && git clone --depth 1 https://github.com/google/ngx_brotli.git \
    && cd ngx_brotli \
    && git submodule update --init \
    && cd ../ \
    \
    && wget https://nginx.org/download/nginx-${VERSION}.tar.gz -O nginx.tar.gz \
    && tar -zxf nginx.tar.gz \
    && cd nginx-${VERSION} \
    && nginx -V 2>&1 | grep -o -- --.* | xargs -I {} ./configure {} --with-compat --add-dynamic-module=../ngx_brotli \
    && make modules \
    && cp objs/ngx_http_brotli*.so /etc/nginx/modules/ \
    \
    && openssl dhparam -out /etc/nginx/dhparam.pem 2048


FROM nginx:${IMAGE_TAG}

ARG TZ

LABEL maintainer="Chotow <i@choyri.com>"

RUN set -x \
    && cp /usr/share/zoneinfo/${TZ:-Asia/Shanghai} /etc/localtime \
    && echo "${TZ:-Asia/Shanghai}" > /etc/timezone \
    \
    && mkdir /etc/nginx/conf.d/sites-enabled

COPY --from=builder /etc/nginx/dhparam.pem                  /etc/nginx/dhparam.pem
COPY --from=builder /usr/lib/nginx/modules/ngx_http_brotli* /etc/nginx/modules/
COPY nginx.conf /etc/nginx/nginx.conf
COPY conf.d/    /etc/nginx/conf.d/

EXPOSE 443
