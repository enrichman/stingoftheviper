package main

type Config struct {
	Verbose int
	Name    string

	StingConfig StingConfig
}

type StingConfig struct {
	Name string
}

func NewConfig() Config {
	return Config{
		Name: "default_name",

		StingConfig: StingConfig{
			Name: "default_sting_name",
		},
	}
}
