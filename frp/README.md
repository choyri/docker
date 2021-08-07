# choyri/frp [Deprecated]

本镜像已停止维护，请使用官方镜像 [fatedier/frps](https://hub.docker.com/r/fatedier/frps)。

---

最小化的 frp 服务端 Docker 镜像。


## 特点

- 最小化
- 时区指定为 `Asia/Shanghai`
  - 如需修改，需要指定 `TZ` 环境变量后重新构建镜像。


## 使用方法

根据 [这个](https://github.com/fatedier/frp/blob/master/conf/frps_full.ini) 完整的服务端配置文件，撰写自己的配置文件 `frps.ini`。

```shell
# 示例命令
docker run -d \
    -p 7000:7000 \
    -p 7500:7500 \
    -v $(pwd)/frps.ini:/frps.ini \
    -v $(pwd)/frps.log:/frps.log \
    --name frp \
    --rm choyri/frp
```

注意，容器内的 frps 位于根目录，并且启动时指定了同一目录的 `frps.ini`，因此，映射配置文件时目标是 `/frps.ini`。

如果需要映射日志文件，需要确保先在本地创建好文件，否则本地会出现一个指定名称的空文件夹；接着，根据配置文件中的日志路径进行映射。


## 客户端

在 [这里](https://github.com/fatedier/frp/releases) 下载最新版本对应平台的客户端。


## frp 官方文档

[English](https://github.com/fatedier/frp/blob/master/README.md) | [中文](https://github.com/fatedier/frp/blob/master/README_zh.md)
