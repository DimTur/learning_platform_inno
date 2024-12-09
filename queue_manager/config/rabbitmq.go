package config

type RabbitMQ struct {
	UserName     string       `yaml:"username"`
	Password     string       `yaml:"password"`
	Host         string       `yaml:"host"`
	Port         int          `yaml:"port"`
	Share        Share        `yaml:"share"`
	Otp          Otp          `yaml:"otp"`
	Chat         Chat         `yaml:"chat"`
	Notification Notification `yaml:"notification"`
	Spfu         Spfu         `yaml:"spfu"`
	Plan         Plan         `yaml:"plan"`
	Channel      Channel      `yaml:"channel"`
}

type QueueConfig struct {
	Name        string    `yaml:"name"`
	Durable     bool      `yaml:"durable"`
	AutoDeleted bool      `yaml:"auto_deleted"`
	Exclusive   bool      `yaml:"exclusive"`
	NoWait      bool      `yaml:"no_wait"`
	Args        QueueArgs `yaml:"args"`
}

type ExchangeConfig struct {
	Name        string       `yaml:"name"`
	Kind        string       `yaml:"kind"`
	Durable     bool         `yaml:"durable"`
	AutoDeleted bool         `yaml:"auto_deleted"`
	Internal    bool         `yaml:"internal"`
	NoWait      bool         `yaml:"no_wait"`
	Args        ExchangeArgs `yaml:"args"`
}

type QueueArgs struct {
	XMessageTtl int32 `yaml:"x_message_ttl"`
}

type ExchangeArgs struct {
	AltExchange string `yaml:"alternate_exchange"`
}

func (e ExchangeArgs) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"alternate-exchange": e.AltExchange,
	}
}

func (q QueueArgs) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"x-message-ttl": q.XMessageTtl,
	}
}
