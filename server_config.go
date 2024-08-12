package nacos

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type LocalNacosConfig struct {
	Ip          string `v:"required#[nacos.config.ip] can not be empty"`
	Port        uint64 `v:"required#[nacos.config.port] can not be empty"`
	NameSpaceId string `v:"required#[nacos.config.name-space-id] can not be empty"`
	Group       string `v:"required#[nacos.config.group] can not be empty"`
	DataId      string `v:"required#[nacos.config.data-id] can not be empty"`
	Username    string `v:"required#[nacos.config.username] can not be empty"`
	Password    string `v:"required#[nacos.config.password] can not be empty"`
}

func LoadRemoteConfig(actions ...func()) {
	ctx := gctx.GetInitCtx()
	enable, err := g.Cfg().Get(ctx, "nacos.cloud.config")
	if err != nil {
		panic(err)
	}
	if enable == nil {
		panic("[nacos.cloud.config] can not be empty, must be true or false")
	}
	if !enable.Bool() {
		return
	}

	configMap, err := g.Cfg().Get(ctx, "nacos.config")
	config := LocalNacosConfig{}
	if err = gconv.Scan(configMap, &config); err != nil {
		panic(err)
	}
	if err = g.Validator().Data(config).Run(ctx); err != nil {
		panic(err)
	}
	serverConfig := constant.ServerConfig{
		IpAddr: config.Ip,
		Port:   config.Port,
	}
	clientConfig := constant.ClientConfig{
		CacheDir:    "/tmp/nacos",
		LogDir:      "/tmp/nacos",
		NamespaceId: config.NameSpaceId,
		Username:    config.Username,
		Password:    config.Password,
	}
	configParam := vo.ConfigParam{
		DataId: config.DataId,
		Group:  config.Group,
	}

	adapter, err := NewConfig(ctx, Config{
		ServerConfigs: []constant.ServerConfig{serverConfig},
		ClientConfig:  clientConfig,
		ConfigParam:   configParam,
		Watch:         true,
	}, actions...)
	if err != nil {
		g.Log().Fatalf(ctx, `%+v`, err)
	}
	g.Cfg().SetAdapter(adapter)
	g.Log().Infof(ctx, "[SUCCESS] Load nacos remote config from: %s:%d; config: [nameSpaceId: %s; group: %s; dataId: %s]", config.Ip, config.Port, config.NameSpaceId, config.Group, config.DataId)
}
