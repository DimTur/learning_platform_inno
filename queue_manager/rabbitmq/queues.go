package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// DeclareQueue announces the queue
func (c *RMQClient) DeclareQueue(
	name string,
	durable, autoDelete, exclusive, noWait bool,
	args map[string]interface{}) (amqp.Queue, error) {

	tableArgs := amqp.Table(args)

	return c.adminCH.QueueDeclare(
		name,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		tableArgs,
	)
}
