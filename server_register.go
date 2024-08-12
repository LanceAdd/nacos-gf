package nacos

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

type RemoteRegistryConfig struct {
	Ip          string `v:"required#[nacos.registry.ip] can not be empty"`
	Port        uint64 `v:"required#[nacos.registry.port] can not be empty"`
	NameSpaceId string `v:"required#[nacos.registry.name-space-id] can not be empty"`
	Group       string `v:"required#[nacos.registry.group] can not be empty"`
	Username    string `v:"required#[nacos.registry.username] can not be empty"`
	Password    string `v:"required#[nacos.registry.password] can not be empty"`
}

func LoadRegisterConfig() {
	ctx := gctx.GetInitCtx()
	enable, err := g.Cfg().Get(ctx, "nacos.cloud.registry")
	if err != nil {
		panic(err)
	}
	if enable == nil {
		panic("[nacos.cloud.registry] can not be empty, must be true or false")
	}
	if !enable.Bool() {
		return
	}
	configMap, err := g.Cfg().Get(ctx, "nacos.registry")
	config := RemoteRegistryConfig{}
	if err = gconv.Scan(configMap, &config); err != nil {
		panic(err)
	}
	if err = g.Validator().Data(config).Run(ctx); err != nil {
		panic(err)
	}
	url := fmt.Sprintf("%s:%d", config.Ip, config.Port)

	address := []string{url}

	registry := NewRegistry(address, constant.WithNamespaceId(config.NameSpaceId), constant.WithUsername(config.Username), constant.WithPassword(config.Password)).SetGroupName(config.Group)
	gsvc.SetRegistry(registry)
	appName, err := g.Cfg().Get(gctx.GetInitCtx(), "app.name")
	if err != nil {
		panic(err)
	}
	g.Log().Infof(ctx, "[SUCCESS] Register server instance [nameSpaceId: %s; group: %s; name: %s] into [ip: %s; port:%d] by [username:%s; password: %s]", config.NameSpaceId, config.Group, appName, config.Ip, config.Port, config.Username, config.Password)

}
