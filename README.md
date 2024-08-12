# nacos-gf

## install

```shell
go get -u -v github.com/LanceAdd/nacos-gf latest
```

## config.yaml

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
    port: 8848 s
    nameSpaceId: "public"
    group: "default"
    username: "nacos"
    password: "nacos"
    clusterName: "default"
```
## Usage

```go
package boot

import (
	"github.com/LanceAdd/nacos-gf"
)

func init() {
	// The parameters of LoadRemoteConfig are custom functions that are used to perform some custom operations after the remote configuration file is updated
	nacos.LoadRemoteConfig(ReInitAuth, ReInitSecurity, ReInitSnapShot)
	nacos.LoadRegisterConfig()
}

```

```go
package main

import (
    _ "example/boot"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	server := g.Server("example")
    server.Run()
}

```

