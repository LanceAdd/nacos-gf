# nacos-gf

## install

```shell
go get -u -v github.com/LanceAdd/nacos-gf latest
```

config.yaml

```yaml
nacos:
  cloud:
    config: true
    registry: true
  config:
    ip: "127.0.0.1"
    port: 8848
    nameSpaceId: "public"
    group: "default"
    dataId: "config.yaml"
    username: "nacos"
    password: "nacos"
  registry:
    ip: "127.0.0.1"
    port: 8848
    nameSpaceId: "public"
    group: "default"
    username: "nacos"
    password: "nacos"
```