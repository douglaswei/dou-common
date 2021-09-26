package main

import (
    consul_config "github.com/douglaswei/go-common/consul-config"
    "fmt"
)

func main() {
    config := consul_config.GetConfig()
    url := config.String("consul.url")
    fmt.Println(url)
}
