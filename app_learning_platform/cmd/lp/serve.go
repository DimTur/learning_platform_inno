package lp

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/DimTur/lp_learning_platform/internal/app"
	"github.com/DimTur/lp_learning_platform/internal/app/consumers"
	ssogrpc "github.com/DimTur/lp_learning_platform/internal/clients/sso/grpc"
	"github.com/DimTur/lp_learning_platform/internal/config"
	"github.com/DimTur/lp_learning_platform/internal/services/rabbitmq"
	"github.com/DimTur/lp_learning_platform/internal/services/redis"
	attstorage "github.com/DimTur/lp_learning_platform/internal/services/storage/postgresql/attempts"
	channelstorage "github.com/DimTur/lp_learning_platform/internal/services/storage/postgresql/channels"
	lessonstorage "github.com/DimTur/lp_learning_platform/internal/services/storage/postgresql/lessons"
	pagestorage "github.com/DimTur/lp_learning_platform/internal/services/storage/postgresql/pages"
	planstorage "github.com/DimTur/lp_learning_platform/internal/services/storage/postgresql/plans"
	questionstorage "github.com/DimTur/lp_learning_platform/internal/services/storage/postgresql/questions"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
)

func NewServeCmd() *cobra.Command {
	var configPath string

	c := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Start gRPS LP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

			ctx, cancel := signal.NotifyContext(cmd.Context(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
			defer cancel()
			var wg sync.WaitGroup

			cfg, err := loadConfig(configPath)
			if err != nil {
				return err
			}

			storagePool, err := initDBConnection(ctx, cfg)
			if err != nil {
				return err
			}
			defer storagePool.Close()

			channelStorage := channelstorage.NewChannelStorage(storagePool)
			planStorage := planstorage.NewPlansStorage(storagePool)
			lessonStorage := lessonstorage.NewLessonsStorage(storagePool)
			pageStorage := pagestorage.NewPagesStorage(storagePool)
			questionStorage := questionstorage.NewQuestionsStorage(storagePool)
			attemptStorage := attstorage.NewAttemptsStorage(storagePool)

			ssoClient, err := ssogrpc.New(
				ctx,
				log,
				cfg.Clients.SSO.Address,
				cfg.Clients.SSO.Timeout,
				cfg.Clients.SSO.RetriesCount,
			)
			if err != nil {
				return err
			}

			// Init Redis
			rAttempts := &redis.RedisAttempts{
				Host:     cfg.Redis.Host,
				Port:     cfg.Redis.Port,
				DB:       cfg.Redis.AttemptsDB,
				Password: cfg.Redis.Password,
			}
			redisAttempts, err := redis.NewRedisClient(*rAttempts)
			if err != nil {
				log.Error("failed to close redis", slog.Any("err", err))
			}

			validate := validator.New()

			// Init RabbitMQ
			rmq, err := initRabbitMQ(cfg)
			if err != nil {
				log.Error("failed init rabbit mq", slog.Any("err", err))
			}

			application, err := app.NewApp(
				channelStorage,
				planStorage,
				lessonStorage,
				pageStorage,
				questionStorage,
				attemptStorage,
				redisAttempts,
				rmq,
				rmq,
				ssoClient,
				cfg.GRPCServer.Address,
				log,
				validate,
			)
			if err != nil {
				return err
			}

			startConsumers(ctx, cfg, rmq, channelStorage, planStorage, log, &wg)

			grpcCloser, err := application.GRPCSrv.Run()
			if err != nil {
				return err
			}

			log.Info("server listening:", slog.Any("port", cfg.GRPCServer.Address))
			<-ctx.Done()
			wg.Wait()

			rmq.Close()
			grpcCloser()

			return nil
		},
	}

	c.Flags().StringVar(&configPath, "config", "", "path to config")
	return c
}

func loadConfig(configPath string) (*config.Config, error) {
	return config.Parse(configPath)
}

func initDBConnection(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Storage.User,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.DBName,
	)
	return pgxpool.New(ctx, dsn)
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

func startConsumers(
	ctx context.Context,
	cfg *config.Config,
	rmq *rabbitmq.RMQClient,
	channelStorage *channelstorage.ChannelPostgresStorage,
	planStorage *planstorage.PlansPostgresStorage,
	log *slog.Logger,
	wg *sync.WaitGroup,
) {
	channelsConsumer := consumers.NewConsumeChannel(rmq, channelStorage, log)
	plansConsumer := consumers.NewConsumePlan(rmq, planStorage, rmq, log)
	learnersConsumer := consumers.NewConsumeSharedLearnersWithPlan(rmq, planStorage, rmq, log)

	wg.Add(3)
	go func() {
		defer wg.Done()
		if err := channelsConsumer.Start(
			ctx,
			cfg.RabbitMQ.Channel.ChannelConsumer.Queue,
			cfg.RabbitMQ.Channel.ChannelConsumer.Consumer,
			cfg.RabbitMQ.Channel.ChannelConsumer.AutoAck,
			cfg.RabbitMQ.Channel.ChannelConsumer.Exclusive,
			cfg.RabbitMQ.Channel.ChannelConsumer.NoLocal,
			cfg.RabbitMQ.Channel.ChannelConsumer.NoWait,
			cfg.RabbitMQ.Channel.ChannelConsumer.ConsumerArgs.ToMap(),
		); err != nil {
			log.Error("failed to start share channels consumer", slog.Any("err", err))
		}
	}()

	go func() {
		defer wg.Done()
		if err := plansConsumer.Start(
			ctx,
			cfg.RabbitMQ.Plan.PlanConsumer.Queue,
			cfg.RabbitMQ.Plan.PlanConsumer.Consumer,
			cfg.RabbitMQ.Plan.PlanConsumer.AutoAck,
			cfg.RabbitMQ.Plan.PlanConsumer.Exclusive,
			cfg.RabbitMQ.Plan.PlanConsumer.NoLocal,
			cfg.RabbitMQ.Plan.PlanConsumer.NoWait,
			cfg.RabbitMQ.Plan.PlanConsumer.ConsumerArgs.ToMap(),
		); err != nil {
			log.Error("failed to start share plans consumer", slog.Any("err", err))
		}
	}()

	go func() {
		defer wg.Done()
		if err := learnersConsumer.Start(
			ctx,
			cfg.RabbitMQ.Spfu.SpfuConsumer.Queue,
			cfg.RabbitMQ.Spfu.SpfuConsumer.Consumer,
			cfg.RabbitMQ.Spfu.SpfuConsumer.AutoAck,
			cfg.RabbitMQ.Spfu.SpfuConsumer.Exclusive,
			cfg.RabbitMQ.Spfu.SpfuConsumer.NoLocal,
			cfg.RabbitMQ.Spfu.SpfuConsumer.NoWait,
			cfg.RabbitMQ.Spfu.SpfuConsumer.ConsumerArgs.ToMap(),
		); err != nil {
			log.Error("failed to start spfu consumer", slog.Any("err", err))
		}
	}()
}
