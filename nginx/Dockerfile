FROM nginx:1.17-alpine

LABEL maintainer="chotow <i@choyri.com>"

ENV NGINX_VERSION 1.17.3

RUN set -x \
    && sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && tempDir="$(mktemp -d)" \
    && chown nobody:nobody $tempDir \
    && apk add --virtual .build-deps \
        gcc \
        libc-dev \
        make \
        pcre-dev \
        zlib-dev \
        git \
    && cd ${tempDir} \
    && git clone https://github.com/eustas/ngx_brotli.git \
    && cd ngx_brotli && git submodule update --init && cd ../ \
    && wget https://nginx.org/download/nginx-$NGINX_VERSION.tar.gz -O nginx.tar.gz \
    && tar -zxf nginx.tar.gz \
    && cd nginx-$NGINX_VERSION && ./configure --with-compat --add-dynamic-module=../ngx_brotli \
    && make modules \
    && cp objs/ngx_http_brotli*.so /etc/nginx/modules/ \
    && rm -rf "$tempDir" \
    && apk del .build-deps

RUN cp /usr/share/zoneinfo/${TZ:-Asia/Shanghai} /etc/localtime && echo "${TZ:-Asia/Shanghai}" > /etc/timezone \
    && mkdir /etc/nginx/proxycache/ \
    && mkdir /etc/nginx/conf.d/stream/

COPY nginx.conf /etc/nginx/nginx.conf
COPY nginx.stream.example.conf /etc/nginx/conf.d/stream/redis.conf.example

EXPOSE 80 443

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]
