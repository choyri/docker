ARG BASE_IMAGE_TAG=fpm-alpine


FROM php:$BASE_IMAGE_TAG

ARG DEV

LABEL maintainer="Chotow <i@choyri.com>"

RUN set -eux; \
    \
    sed -i 's/http[s]:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories; \ 
    apk update; \
    \
    apk add --virtual .build-deps $PHPIZE_DEPS; \
    \
    # redis
    pecl install redis; \
    docker-php-ext-enable redis; \
    \
    # imagick
    apk add libgomp imagemagick-dev; \
    pecl install imagick; \
    docker-php-ext-enable imagick; \
    \
    # gd
    apk add \
        libjpeg-turbo-dev \
        libwebp-dev \
        freetype-dev \
    ; \
    if [[ ! $PHP_VERSION =~ "^(7\.4|8)" ]]; then \
        apk add libpng-dev; \
        docker-php-ext-configure gd \
            --with-jpeg-dir=/usr/include/ \
            --with-png-dir=/usr/include/ \
            --with-webp-dir=/usr/include/ \
            --with-freetype-dir=/usr/include/; \
    else \
        docker-php-ext-configure gd --with-jpeg --with-webp --with-freetype; \
    fi; \
    docker-php-ext-install -j$(nproc) gd; \
    \
    # some other
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
    if [ -n ${DEV:-''} ]; then \
        pecl install xdebug; \
        docker-php-ext-enable xdebug; \
    else \
        docker-php-ext-enable opcache; \
    fi; \
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
    rm -rf /tmp/pear; \
    \
    php -v

COPY ./php.ini      /usr/local/etc/php/
COPY ./conf.d/      /usr/local/etc/php/conf.d/
COPY ./php-fpm.d/   /usr/local/etc/php-fpm.d/

WORKDIR /srv
