package nacos

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var (
	_ gsvc.Registry = &Registry{}
)

// RegistryConfig is the configuration object for nacos client.
type RegistryConfig struct {
	ServerConfigs []constant.ServerConfig `v:"required"` // See constant.ServerConfig
	ClientConfig  *constant.ClientConfig  `v:"required"` // See constant.ClientConfig
}

// NewRegistry new a registry with address and opts
func NewRegistry(address []string, opts ...constant.ClientOption) (reg *Registry) {
	clientConfig := constant.NewClientConfig(opts...)

	if len(clientConfig.NamespaceId) == 0 {
		clientConfig.NamespaceId = "public"
	}

	serverConfigs := make([]constant.ServerConfig, 0, len(address))
	for _, endpoint := range address {
		tmp := gstr.Split(endpoint, ":")
		ip := tmp[0]
		port := gconv.Uint64(tmp[1])
		if port == 0 {
			port = 8848
		}
		serverConfigs = append(serverConfigs, *constant.NewServerConfig(ip, port))
	}
	ctx := gctx.New()
	reg, err := NewWithConfig(ctx, RegistryConfig{
		ServerConfigs: serverConfigs,
		ClientConfig:  clientConfig,
	})

	if err != nil {
		panic(gerror.Wrap(err, `create nacos client failed`))
	}
	return
}

// NewConfig creates and returns registry with Config.
func NewWithConfig(ctx context.Context, config RegistryConfig) (reg *Registry, err error) {
	// Data validation.
	err = g.Validator().Data(config).Run(ctx)
	if err != nil {
		return nil, err
	}

	nameingClient, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  config.ClientConfig,
		ServerConfigs: config.ServerConfigs,
	})
	if err != nil {
		return
	}
	return NewWithClient(nameingClient), nil
}

// NewWithClient new the instance with INamingClient
func NewWithClient(client naming_client.INamingClient) *Registry {
	r := &Registry{
		client:      client,
		clusterName: "DEFAULT",
		groupName:   "DEFAULT_GROUP",
	}
	return r
}

// SetClusterName can set the clusterName. The default is 'DEFAULT'
func (reg *Registry) SetClusterName(clusterName string) *Registry {
	reg.clusterName = clusterName
	return reg
}

// SetGroupName can set the groupName. The default is 'DEFAULT_GROUP'
func (reg *Registry) SetGroupName(groupName string) *Registry {
	reg.groupName = groupName
	return reg
}
