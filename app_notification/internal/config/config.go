package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Server      Server      `yaml:"server"`
	TelegramBot TelegramBot `yaml:"telegram_bot"`
	RabbitMQ    RabbitMQ    `yaml:"rabbit_mq"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type TelegramBot struct {
	TgBotToken string `yaml:"tg_bot_token"`
	TgBotHost  string `yaml:"tg_bot_host"`
	BatchSize  int    `yaml:"batch_size"`
}

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	return c, nil
}
