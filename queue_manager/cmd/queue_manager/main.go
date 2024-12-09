package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/DimTur/lp_queue_manager/config"
	"github.com/DimTur/lp_queue_manager/rabbitmq"
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	var configPath string

	c := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "init rabbitmq queues, exchanges and bind it",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

			cfg, err := loadConfig(configPath)
			if err != nil {
				return err
			}

			// Init RabbitMQ
			rmq, err := initRabbitMQ(cfg)
			if err != nil {
				log.Error("failed init rabbit mq", slog.Any("err", err))
			}

			// Declare Share exchange
			if err := declareExchange(rmq, cfg.RabbitMQ.Share.ShareExchange); err != nil {
				log.Error("failed to declare Share exchange", slog.Any("err", err))
			}

			// Declare OTP exchange
			if err := declareExchange(rmq, cfg.RabbitMQ.Otp.OtpExchange); err != nil {
				log.Error("failed to declare OTP exchange", slog.Any("err", err))
			}

			// Declare Chat exchange
			if err := declareExchange(rmq, cfg.RabbitMQ.Chat.ChatIDExchange); err != nil {
				log.Error("failed to declare Chat exchange", slog.Any("err", err))
			}

			// Declare and bind OTP queue
			if err := declareQueueAndBind(
				rmq,
				cfg.RabbitMQ.Otp.OtpQueue,
				cfg.RabbitMQ.Otp.OtpExchange.Name,
				cfg.RabbitMQ.Otp.OtpRoutingKey,
			); err != nil {
				log.Error("failed to declare and bind OTP Queue", slog.Any("err", err))
			}

			// Declare and bind Chat queue
			if err := declareQueueAndBind(
				rmq,
				cfg.RabbitMQ.Chat.ChatIDQueue,
				cfg.RabbitMQ.Chat.ChatIDExchange.Name,
				cfg.RabbitMQ.Chat.ChatIDRoutingKey,
			); err != nil {
				log.Error("failed to declare and bind Chat Queue", slog.Any("err", err))
			}

			// Declare and bind Notification queue
			if err := declareQueueAndBind(
				rmq,
				cfg.RabbitMQ.Notification.NotificationQueue,
				cfg.RabbitMQ.Share.ShareExchange.Name,
				cfg.RabbitMQ.Notification.NotificationRoutingKey,
			); err != nil {
				log.Error("failed to declare and bind Notification Queue", slog.Any("err", err))
			}

			// Declare and bind notification_to_auth queue
			if err := declareQueueAndBind(
				rmq,
				cfg.RabbitMQ.Notification.NotificationToAuthQueue,
				cfg.RabbitMQ.Share.ShareExchange.Name,
				cfg.RabbitMQ.Notification.NotificationToAuthRoutingKey,
			); err != nil {
				log.Error("failed to declare and bind notification_to_auth Queue", slog.Any("err", err))
			}

			// Declare and bind Spfu queue
			if err := declareQueueAndBind(
				rmq,
				cfg.RabbitMQ.Spfu.SpfuQueue,
				cfg.RabbitMQ.Share.ShareExchange.Name,
				cfg.RabbitMQ.Spfu.SpfuRoutingKey,
			); err != nil {
				log.Error("failed to declare and bind Spfu Queue", slog.Any("err", err))
			}

			// Declare and bind Plan queue
			if err := declareQueueAndBind(
				rmq,
				cfg.RabbitMQ.Plan.PlanQueue,
				cfg.RabbitMQ.Share.ShareExchange.Name,
				cfg.RabbitMQ.Plan.PlanRoutingKey,
			); err != nil {
				log.Error("failed to declare and bind Plan Queue", slog.Any("err", err))
			}

			// Declare and bind Channel queue
			if err := declareQueueAndBind(
				rmq,
				cfg.RabbitMQ.Channel.ChannelQueue,
				cfg.RabbitMQ.Share.ShareExchange.Name,
				cfg.RabbitMQ.Channel.ChannelRoutingKey,
			); err != nil {
				log.Error("failed to declare and bind Channel Queue", slog.Any("err", err))
			}

			return nil
		},
	}

	c.Flags().StringVar(&configPath, "config", "", "path to config")
	return c
}

func loadConfig(configPath string) (*config.Config, error) {
	return config.Parse(configPath)
}

func initRabbitMQ(cfg *config.Config) (*rabbitmq.RMQClient, error) {
	rmqUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.UserName,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	return rabbitmq.NewClient(rmqUrl)
}

func declareExchange(rmq *rabbitmq.RMQClient, exchangeConfig config.ExchangeConfig) error {
	return rmq.DeclareExchange(
		exchangeConfig.Name,
		exchangeConfig.Kind,
		exchangeConfig.Durable,
		exchangeConfig.AutoDeleted,
		exchangeConfig.Internal,
		exchangeConfig.NoWait,
		exchangeConfig.Args.ToMap(),
	)
}

func declareQueueAndBind(rmq *rabbitmq.RMQClient, queueConfig config.QueueConfig, exchangeName, routingKey string) error {
	// Announcement of the queue
	if _, err := rmq.DeclareQueue(
		queueConfig.Name,
		queueConfig.Durable,
		queueConfig.AutoDeleted,
		queueConfig.Exclusive,
		queueConfig.NoWait,
		queueConfig.Args.ToMap(),
	); err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", queueConfig.Name, err)
	}

	// Binding a queue to an exchange
	if err := rmq.BindQueueToExchange(
		queueConfig.Name,
		exchangeName,
		routingKey,
	); err != nil {
		return fmt.Errorf("failed to bind queue %s to exchange %s: %w", queueConfig.Name, exchangeName, err)
	}

	return nil
}
