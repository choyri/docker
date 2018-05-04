## frp 官方文档

[English](https://github.com/fatedier/frp/blob/master/README.md) | [中文](https://github.com/fatedier/frp/blob/master/README_zh.md)


## 运行

😆 先试试如何。

```
docker run -itd --rm --name frp choyri/frp
```

不错，把配置复制出来，然后删掉。

```
docker cp frp:/frps_full.ini ./frps.ini \
    && docker stop frp
```

👏 改一下，启动！

「端口请根据配置自行更改」

```
touch frps.log
docker run -itd \
    -p 7000:7000 \
    -p 7080:7080 \
    -p 7500:7500 \
    -v $(pwd)/frps.ini:/frps.ini \
    -v $(pwd)/frps.log:/frps.log \
    --name frp \
    --rm choyri/frp
```


## 复制客户端

```
docker cp frp:/frpc ./frpc
docker cp frp:/frpc_full.ini ./frpc.ini
```
