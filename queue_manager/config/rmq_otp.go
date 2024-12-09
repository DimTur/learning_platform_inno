package config

type Otp struct {
	OtpExchange   ExchangeConfig `yaml:"otp_exchange"`
	OtpQueue      QueueConfig    `yaml:"otp_queue"`
	OtpRoutingKey string         `yaml:"otp_routing_key"`
}
