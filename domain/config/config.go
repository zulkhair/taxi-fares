package config

type Config struct {
	Log Log `yaml:"log"`
}
type Log struct {
	File string `yaml:"file"`
}
