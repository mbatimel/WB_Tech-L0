package config

type Config struct {
	Server Server `yaml:"server"`
	Repo   Repo   `yaml:"repo"`
}
