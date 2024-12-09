package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	RabbitMQ RabbitMQ `yaml:"rabbit_mq"`
}

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	return c, nil
}
