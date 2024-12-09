package config

type Plan struct {
	PlanQueue      QueueConfig `yaml:"plan_queue"`
	PlanRoutingKey string      `yaml:"plan_routing_key"`
}
