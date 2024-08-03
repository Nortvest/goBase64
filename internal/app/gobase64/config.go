package gobase64

type Config struct {
	BindAdrr string `toml:"bind_adrr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{}
}
