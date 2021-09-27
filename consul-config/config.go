package consul_config

import (
    "github.com/gookit/config/v2"
    "github.com/gookit/config/v2/yaml"
    consulapi "github.com/hashicorp/consul/api"
    "log"
)

var (
    ConsulPath string = "go/ProjInit"
    ConsulUrl string = "consul.url"
)

func GetConfig() *config.Config {
    consulAddress := config.Getenv("consulAddress", "consul-dev:8500")
    return GetConfigByAddress(consulAddress)
}

func GetConfigByAddress(consulAddress string) *config.Config {
    config.AddDriver(yaml.Driver)
    dftCfg := config.Default()
    consulClient, err := consulapi.NewClient(&consulapi.Config{Address: consulAddress})

    if err != nil {
        log.Fatalln("consul连接失败:", err)
    }

    kv, _, err := consulClient.KV().Get(ConsulPath, nil)
    if err != nil {
        log.Fatalln("consul获取配置失败:", err)
    }

    err = dftCfg.LoadSources(config.Yaml, kv.Value)
    if err != nil {
        log.Fatalln("Viper解析配置失败:", err)
    }

    return dftCfg
}
