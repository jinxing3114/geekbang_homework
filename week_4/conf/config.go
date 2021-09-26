package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var conf *Config

//Config 配置信息
type Config struct {
	Redis  *Redis  `yaml:"redis"`
	Server *Server `yaml:"server"`
	Mode   string
}

//InitConfig 初始化配置信息
func InitConfig() *Config {
	configFile := "conf-local.yml"
	envMode := os.Getenv("ENV_MODE")
	if envMode == "produce" {
		configFile = "conf.yml"
	}
	conf = &Config{}

	if conf != nil {
		return conf
	}

	_, err := os.Lstat(configFile)
	if os.IsNotExist(err) {
		panic(errors.New("yml config file is not exist"))
	}
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(err)
	}
	return conf
}

// GetConf 获取配置信息
func GetConf() *Config {
	return conf
}
