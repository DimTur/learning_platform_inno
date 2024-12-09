package config

type RabbitMQ struct {
	UserName string  `yaml:"username"`
	Password string  `yaml:"password"`
	Host     string  `yaml:"host"`
	Port     int     `yaml:"port"`
	Channel  Channel `yaml:"channel"`
	Plan     Plan    `yaml:"plan"`
	Spfu     Spfu    `yaml:"spfu"`
}

type ConsumerConfig struct {
	Queue        string       `yaml:"queue"`
	Consumer     string       `yaml:"consumer"`
	AutoAck      bool         `yaml:"autoAck"`
	Exclusive    bool         `yaml:"exclusive"`
	NoLocal      bool         `yaml:"noLocal"`
	NoWait       bool         `yaml:"noWait"`
	ConsumerArgs ConsumerArgs `yaml:"args"`
}

type ConsumerArgs struct {
	XConsumerTtl       int32 `yaml:"x-consumer-timeout"`
	XConsumerPrefCount int32 `yaml:"x-consumer-prefetch-count"`
}

func (c ConsumerArgs) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"x-consumer-timeout":        c.XConsumerTtl,
		"x-consumer-prefetch-count": c.XConsumerPrefCount,
	}
}
