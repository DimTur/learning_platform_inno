package config

type Spfu struct {
	SpfuQueue      QueueConfig `yaml:"spfu_queue"`
	SpfuRoutingKey string      `yaml:"spfu_routing_key"`
}
