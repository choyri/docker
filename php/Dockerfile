ARG IMAGE_TAG=fpm-alpine

FROM php:$IMAGE_TAG

LABEL maintainer="Chotow <i@choyri.com>"

ARG IMAGE_TAG

RUN set -eux; \
    \
    echo "Image Tag: $IMAGE_TAG"; \
    \
    sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories; \ 
    apk update; \
    \
    apk add --virtual .build-deps $PHPIZE_DEPS; \
    \
    pecl install redis; \
    docker-php-ext-enable redis; \
    \
    if [ $(expr match $IMAGE_TAG 8) == 0 ]; then \
        apk add imagemagick-dev; \
        pecl install imagick; \
        docker-php-ext-enable imagick; \
    fi; \
    \
    apk add \
        libjpeg-turbo-dev \
        libwebp-dev \
        freetype-dev \
    ; \
    docker-php-ext-configure gd --with-jpeg --with-webp --with-freetype; \
    docker-php-ext-install -j$(nproc) gd; \
    \
    apk add \
        libzip-dev \
        gmp-dev \
    ; \
    docker-php-ext-install -j$(nproc) \
        zip \
        gmp \
        mysqli \
        pdo_mysql \
        exif \
    ; \
    \
    docker-php-ext-enable opcache; \
    \
    apk del .build-deps; \
    \
    # https://wiki.alpinelinux.org/wiki/Setting_the_timezone
    apk add tzdata; \
    cp /usr/share/zoneinfo/${TZ:-Asia/Shanghai} /etc/localtime; \
    echo "${TZ:-Asia/Shanghai}" > /etc/timezone; \
    date; \
    apk del tzdata; \
    \
    rm -rf /var/cache/apk/*; \
    \
    php -v

COPY ./php.ini      /usr/local/etc/php/
COPY ./conf.d/      /usr/local/etc/php/conf.d/
COPY ./php-fpm.d/   /usr/local/etc/php-fpm.d/

WORKDIR /srv
