package main

import (
    consul_config "github.com/dou.go.common/consul-config"
    "fmt"
)

func main() {
    config := consul_config.GetConfig()
    url := config.String("consul.url")
    fmt.Println(url)
}
