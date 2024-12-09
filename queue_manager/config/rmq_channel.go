package config

type Channel struct {
	ChannelQueue      QueueConfig `yaml:"channel_queue"`
	ChannelRoutingKey string      `yaml:"channel_routing_key"`
}
