package db

import (
    "errors"
    "github.com/gookit/config/v2"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var (
    DSN string = "dsn"
)

func GetGormFromkey(cfg *config.Config, key string) (db *gorm.DB, err error) {
    if _,ok := cfg.GetValue(key); false == ok {
        err = errors.New("key not found: " + key)
        return
    }
    dsnKey := key + "." + DSN
    dsnVal := cfg.String(dsnKey)
    db, err = gorm.Open(mysql.Open(dsnVal), &gorm.Config{})
    return
}
