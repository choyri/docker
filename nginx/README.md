# choyri/nginx

基于原生 [nginx](https://hub.docker.com/_/nginx)，自定义优化。


## Feature

- 增加 [ngx_brotli](https://github.com/google/ngx_brotli)
- 自定义 [nginx.conf](./nginx.conf)
- 时区指定为 `Asia/Shanghai`


## Tip

- `/srv/`
  - 默认 [`root`](https://nginx.org/en/docs/http/ngx_http_core_module.html#root) 目录
- `/etc/nginx/conf.d/`
  - `default.conf` 默认主机
  - `general.conf` 综合配置项，供导入
  - `security.conf` 安全配置项，供导入
  - `security_with_hsts.conf` 带 HSTS 的安全配置项，供导入
- `/etc/nginx/conf.d/sites-enabled/`
  - 站点目录

`reuseport` 参数对于一组 `host:port` 仅需要声明一次，不可重复。


## Example

```yml
# docker-compose.yml

version: "3.8"

services:
  nginx:
    image: choyri/nginx
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./conf:/etc/nginx/conf.d/sites-enabled:ro
      - ./ssl:/etc/nginx/ssl:ro
      - ./log:/var/log/nginx
      - ./srv:/srv
```

```
# /etc/nginx/conf.d/sites-enabled/example.com.conf

server {
    listen                  443 ssl http2 reuseport;
    listen                  [::]:443 ssl http2 reuseport;
    server_name             example.com;

    include                 /etc/nginx/conf.d/general.conf;
    include                 /etc/nginx/conf.d/security_with_hsts.conf;

    access_log              /var/log/nginx/example.com.access.log;
    error_log               /var/log/nginx/example.com.error.log warn;

    ssl_certificate         /etc/nginx/ssl/example.com.cer;
    ssl_certificate_key     /etc/nginx/ssl/example.com.key;
    ssl_trusted_certificate /etc/nginx/ssl/chain.pem; # 中间证书 & 根证书

    location / {
        root    /usr/share/nginx/html;
    }
}

server {
    listen      80;
    listen      [::]:80;
    server_name example.com;

    location / {
        return 301 https://example.com$request_uri;
    }
}
```
