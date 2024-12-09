package config

type Notification struct {
	NotificationQueue            QueueConfig `yaml:"notification_queue"`
	NotificationRoutingKey       string      `yaml:"notification_routing_key"`
	NotificationToAuthQueue      QueueConfig `yaml:"notification_to_auth_queue"`
	NotificationToAuthRoutingKey string      `yaml:"notification_to_auth_routing_key"`
}
