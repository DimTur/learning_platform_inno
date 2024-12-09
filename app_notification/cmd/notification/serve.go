package notification

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/DimTur/lp_notification/internal/app/sender"
	"github.com/DimTur/lp_notification/internal/app/telegram"
	tgclient "github.com/DimTur/lp_notification/internal/clients/telegram"
	"github.com/DimTur/lp_notification/internal/config"
	rabbitmq_store "github.com/DimTur/lp_notification/internal/storage/rabbitmq"
	"github.com/spf13/cobra"
)

func NewServeCmd() *cobra.Command {
	var configPath string

	c := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Start API server",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

			ctx, cancel := signal.NotifyContext(cmd.Context(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
			defer cancel()
			var wg sync.WaitGroup

			cfg, err := config.Parse(configPath)
			if err != nil {
				return err
			}

			tgClient, err := tgclient.NewTgClient(
				cfg.TelegramBot.TgBotHost,
				cfg.TelegramBot.TgBotToken,
				log,
			)
			if err != nil {
				log.Error("failed init tg client", slog.Any("err", err))
			}

			// Init RabbitMQ
			rmq, err := initRabbitMQ(cfg)
			if err != nil {
				log.Error("failed init rabbit mq", slog.Any("err", err))
			}

			// start tg bot
			wg.Add(1)
			go func() {
				defer wg.Done()
				telegram.RunTg(
					ctx,
					tgClient,
					cfg.TelegramBot.BatchSize,
					rmq,
					log,
				)
			}()

			startConsumers(ctx, cfg, rmq, tgClient, log, &wg)

			log.Info("tg bot starting at:", slog.Any("port", cfg.Server.Port))
			<-ctx.Done()
			wg.Wait()

			return nil
		}}

	c.Flags().StringVar(&configPath, "config", "", "path to config")
	return c
}

func initRabbitMQ(cfg *config.Config) (*rabbitmq_store.RMQClient, error) {
	rmqUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.UserName,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	return rabbitmq_store.NewClient(rmqUrl)
}

func startConsumers(
	ctx context.Context,
	cfg *config.Config,
	rmq *rabbitmq_store.RMQClient,
	tgClient *tgclient.TgClient,
	log *slog.Logger,
	wg *sync.WaitGroup,
) {
	otpConsumer := sender.NewConsumeOTP(rmq, tgClient, log)
	shareConsumer := sender.NewConsumeNotification(rmq, tgClient, log)

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := otpConsumer.Start(
			ctx,
			cfg.RabbitMQ.OTP.OTPConsumer.Queue,
			cfg.RabbitMQ.OTP.OTPConsumer.Consumer,
			cfg.RabbitMQ.OTP.OTPConsumer.AutoAck,
			cfg.RabbitMQ.OTP.OTPConsumer.Exclusive,
			cfg.RabbitMQ.OTP.OTPConsumer.NoLocal,
			cfg.RabbitMQ.OTP.OTPConsumer.NoWait,
			cfg.RabbitMQ.OTP.OTPConsumer.ConsumerArgs.ToMap(),
		); err != nil {
			log.Error("failed to start otp consumer", slog.Any("err", err))
		}
	}()

	go func() {
		defer wg.Done()
		if err := shareConsumer.Start(
			ctx,
			cfg.RabbitMQ.Notification.NotificationConsumer.Queue,
			cfg.RabbitMQ.Notification.NotificationConsumer.Consumer,
			cfg.RabbitMQ.Notification.NotificationConsumer.AutoAck,
			cfg.RabbitMQ.Notification.NotificationConsumer.Exclusive,
			cfg.RabbitMQ.Notification.NotificationConsumer.NoLocal,
			cfg.RabbitMQ.Notification.NotificationConsumer.NoWait,
			cfg.RabbitMQ.Notification.NotificationConsumer.ConsumerArgs.ToMap(),
		); err != nil {
			log.Error("failed to start notification consumer", slog.Any("err", err))
		}
	}()
}
