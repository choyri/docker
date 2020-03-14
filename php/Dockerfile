ARG IMAGE_TAG=fpm-alpine

FROM composer:latest AS composer

FROM php:${IMAGE_TAG}

ENV COMPOSER_ALLOW_SUPERUSER=1
COPY --from=composer /usr/bin/composer /usr/bin/composer

RUN sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add \
        libzip-dev \
        libjpeg-turbo-dev \
        libwebp-dev \
        freetype-dev \
        imagemagick-dev \
    && apk add --virtual .build-deps $PHPIZE_DEPS \
    && pecl channel-update pecl.php.net && pecl install redis imagick \
    && docker-php-ext-enable redis imagick opcache \
    && docker-php-ext-configure gd --with-jpeg --with-webp --with-freetype \
    && docker-php-ext-install -j$(nproc) gd exif zip mysqli pdo_mysql \
    && apk del .build-deps \
    && apk add tzdata \
    && rm -rf /tmp/pear ~/.pearrc /var/cache/apk/*

RUN cp /usr/share/zoneinfo/${TZ:-Asia/Shanghai} /etc/localtime && echo "${TZ:-Asia/Shanghai}" > /etc/timezone \
    && mkdir -p /var/log/php \
    && chown www-data:www-data /var/log/php

COPY ./php.ini      /usr/local/etc/php/php.ini
COPY ./conf.d/      /usr/local/etc/php/conf.d/
COPY ./php-fpm.d/   /usr/local/etc/php-fpm.d/

WORKDIR /srv
