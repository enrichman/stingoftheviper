package main

type Config struct {
	Name   string
	Number int

	StingConfig StingConfig
}

type StingConfig struct {
	Name string
}

func NewConfig() Config {
	return Config{
		Number: 42,
		Name:   "default_name",

		StingConfig: StingConfig{
			Name: "default_sting_name",
		},
	}
}
