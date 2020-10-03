ARG IMAGE_TAG=fpm-alpine

FROM composer:latest AS composer

FROM php:${IMAGE_TAG}

LABEL maintainer="Chotow <i@choyri.com>"

ENV COMPOSER_ALLOW_SUPERUSER=1
COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN set -x \
    && if [ -z "${ON_HUB}" ]; then \
        sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories; \
    fi \
    && apk update \
    && apk add \
        libzip-dev \
        libjpeg-turbo-dev \
        libwebp-dev \
        freetype-dev \
        imagemagick-dev \
        gmp-dev \
    && apk add --virtual .build-deps $PHPIZE_DEPS \
    \
    && pecl channel-update pecl.php.net && pecl install redis imagick \
    \
    && docker-php-ext-enable redis imagick opcache \
    && docker-php-ext-configure gd --with-jpeg --with-webp --with-freetype \
    && docker-php-ext-install -j$(nproc) gd exif zip gmp mysqli pdo_mysql \
    \
    && apk del .build-deps \
    && apk add tzdata \
    \
    && rm -rf /tmp/pear ~/.pearrc /var/cache/apk/*

RUN set -x \
    && cp /usr/share/zoneinfo/${TZ:-Asia/Shanghai} /etc/localtime \
    && echo "${TZ:-Asia/Shanghai}" > /etc/timezone

COPY ./php.ini      /usr/local/etc/php/
COPY ./conf.d/      /usr/local/etc/php/conf.d/
COPY ./php-fpm.d/   /usr/local/etc/php-fpm.d/

WORKDIR /srv
