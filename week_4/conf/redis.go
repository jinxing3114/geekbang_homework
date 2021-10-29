package conf

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Pass string `yaml:"pass"`
	Db   int    `yaml:"db"`
}
