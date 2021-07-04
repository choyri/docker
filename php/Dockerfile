ARG BASE_IMAGE_TAG=fpm-alpine

FROM php:$BASE_IMAGE_TAG

LABEL maintainer="Chotow <i@choyri.com>"

ARG BASE_IMAGE_TAG
ARG DOCKER_TAG=unknown

RUN set -eux; \
    \
    echo "Base Image Tag: $BASE_IMAGE_TAG"; \
    echo "Docker Tag: $DOCKER_TAG"; \
    \
    sed -i 's/http[s]:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories; \ 
    apk update; \
    \
    apk add --virtual .build-deps $PHPIZE_DEPS; \
    \
    pecl install redis; \
    docker-php-ext-enable redis; \
    \
    apk add imagemagick-dev; \
    pecl install imagick; \
    docker-php-ext-enable imagick; \
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
    if echo $DOCKER_TAG | grep -q "dev"; then \
        pecl install xdebug; \
        docker-php-ext-enable xdebug; \
    else \
        docker-php-ext-enable opcache; \
    fi; \
    \
    apk del .build-deps; \
    \
    # for imagick
    apk add libgomp; \
    \
    # https://wiki.alpinelinux.org/wiki/Setting_the_timezone
    apk add tzdata; \
    cp /usr/share/zoneinfo/${TZ:-Asia/Shanghai} /etc/localtime; \
    echo "${TZ:-Asia/Shanghai}" > /etc/timezone; \
    date; \
    apk del tzdata; \
    \
    rm -rf /var/cache/apk/*; \
    rm -rf /tmp/pear; \
    \
    php -v

COPY ./php.ini      /usr/local/etc/php/
COPY ./conf.d/      /usr/local/etc/php/conf.d/
COPY ./php-fpm.d/   /usr/local/etc/php-fpm.d/

WORKDIR /srv
