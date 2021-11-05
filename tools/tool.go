package tools

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// GetPort 通过yaml返回服务器运行端口
func GetPort(filepath string) int {
	type Config struct {
		Port int `yaml:"port"`
	}
	var conf Config
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return conf.Port
}