package config

type Chat struct {
	ChatIDExchange   ExchangeConfig `yaml:"chat_id_exchange"`
	ChatIDQueue      QueueConfig    `yaml:"chat_id_queue"`
	ChatIDRoutingKey string         `yaml:"chat_id_routing_key"`
}
