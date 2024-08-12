# nacos-gf

## install
```shell
go get -u -v github.com/LanceAdd/nacos-gf latest
```

## config.yaml
```yaml
nacos:
  cloud:
    config: false
    registry: true
  config:
    ip: "172.16.23.99"
    port: 8848
    nameSpaceId: "c1e15245-03c1-4cba-9c40-ddc59e4d4167"
    group: "coneall"
    dataId: "sso.yml"
    username: "nacos"
    password: "nacos"
  registry:
    ip: "172.16.23.99"
    port: 8848
    nameSpaceId: "c1e15245-03c1-4cba-9c40-ddc59e4d4167"
    group: "coneall"
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
	server := g.Server()
	server.Run()
}

```