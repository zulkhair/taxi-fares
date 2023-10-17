package config

type Config struct {
	Log   Log `yaml:"log"`
	Fares int `yaml:"fares"`
}
type Log struct {
	File string `yaml:"file"`
}
