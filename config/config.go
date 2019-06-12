package config

type Config struct {
	Name        string
	Environment string
}

func Newconfig() *Config {
	viper.echoSetConfigName("config")
}
