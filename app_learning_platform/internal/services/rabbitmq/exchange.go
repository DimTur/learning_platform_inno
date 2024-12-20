package rabbitmq

import (
	"context"

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

// Publish send message to exchange
func (c *RMQClient) Publish(ctx context.Context, exchange, routingKey string, body []byte) error {
	return c.publishWithContext(ctx, exchange, routingKey, body)
}

// publishWithContext send message with context
func (c *RMQClient) publishWithContext(ctx context.Context, exchange, routingKey string, body []byte) error {
	errChan := make(chan error, 1)

	go func() {
		errChan <- c.publishCH.Publish(
			exchange,
			routingKey,
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
