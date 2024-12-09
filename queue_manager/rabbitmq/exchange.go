package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// DeclareExchange announce exchange
func (c *RMQClient) DeclareExchange(
	name, kind string,
	durable, autoDelete, internal, noWait bool,
	args map[string]interface{},
) error {
	// Change args to amqp.Table
	tableArgs := amqp.Table(args)

	return c.adminCH.ExchangeDeclare(
		name,       // exchange name
		kind,       // exchange type
		durable,    // durable
		autoDelete, // auto-delete
		internal,   // internal
		noWait,     // no-wait
		tableArgs,  // arguments as amqp.Table
	)
}

// BindQueueToExchange binds the queue to the exchanger
func (c *RMQClient) BindQueueToExchange(queueName, exchangeName, routingKey string) error {
	return c.adminCH.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false, // no-wait
		nil,   // arguments
	)
}
