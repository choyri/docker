# choyri/php

基于原生 [php](https://hub.docker.com/_/php)，自定义优化。


## Feature

- 增加 [composer](https://getcomposer.org/)
- 启用若干扩展
- 自定义 [php.ini](./php.ini)
- 增加额外的 `php-fpm` [配置文件](./php-fpm.d/www.user.conf)
- 时区指定为 `Asia/Shanghai`


## Tip

`php.ini` 中未设置 `error_log`，且 `php-fpm` 配置文件置空了 `php_admin_flag[log_errors]`，该操作是为了让 php 的错误信息输出到 php-fpm worker。
