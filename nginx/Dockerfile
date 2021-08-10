ARG BASE_IMAGE_TAG=alpine


FROM nginx:${BASE_IMAGE_TAG} AS builder

ARG BASE_IMAGE_TAG

RUN set -eux; \
    \
    VERSION=$(echo $NGINX_VERSION); \
    echo "Base Image Tag: ${BASE_IMAGE_TAG}"; \
    echo "Version: ${VERSION}"; \
    \
    sed -i 's/http[s]:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories; \ 
    apk update; \
    \
    apk add --no-cache --virtual .build-deps \
        gcc \
        libc-dev \
        make \
        pcre-dev \
        zlib-dev \
        git \
    ; \
    wget https://nginx.org/download/nginx-${VERSION}.tar.gz -O nginx.tar.gz; \
    tar -zxf nginx.tar.gz; \
    cd nginx-${VERSION}; \
    \
    git clone --depth 1 --recursive https://github.com/google/ngx_brotli.git /tmp/ngx_brotli; \
    nginx -V 2>&1 | grep -o -- --.* | xargs -I {} ./configure {} --with-compat --add-dynamic-module=/tmp/ngx_brotli; \
    make modules; \
    cp objs/ngx_http_brotli*.so /etc/nginx/modules/; \
    \
    # https://github.com/mozilla/ssl-config-generator/issues/60
    curl https://ssl-config.mozilla.org/ffdhe2048.txt > /etc/nginx/dhparam.pem; \
    \
    nginx -v


FROM nginx:${BASE_IMAGE_TAG}

LABEL maintainer="Chotow <i@choyri.com>"

RUN set -eux; \
    \
    # https://wiki.alpinelinux.org/wiki/Setting_the_timezone
    apk add tzdata; \
    cp /usr/share/zoneinfo/${TZ:-Asia/Shanghai} /etc/localtime; \
    echo "${TZ:-Asia/Shanghai}" > /etc/timezone; \
    date; \
    apk del tzdata; \
    \
    mkdir /etc/nginx/conf.d/sites-enabled

COPY --from=builder /etc/nginx/dhparam.pem              /etc/nginx/dhparam.pem
COPY --from=builder /etc/nginx/modules/ngx_http_brotli* /etc/nginx/modules/
COPY nginx.conf /etc/nginx/nginx.conf
COPY conf.d/    /etc/nginx/conf.d/

EXPOSE 443
