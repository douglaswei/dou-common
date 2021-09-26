package consul_config

import (
    "fmt"
    "testing"
)

func TestFuc(t *testing.T) {
    config := GetConfig()
    url := config.String("consul.url")
    fmt.Println(url)
}