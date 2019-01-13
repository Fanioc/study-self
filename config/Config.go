package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

var Conf *toml.Tree

/**
 * 返回单例实例
 * @method New
 */
func New() *toml.Tree {
	config, err := toml.LoadFile("./config/config.toml")

	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}
	Conf = config
	return config
}

func init() {
	New()
}

func Get(key string) interface{} {
	return Conf.Get(key)
}
