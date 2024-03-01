package app

import (
	"github.com/Alex-omosa/pricing-service/router"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type App struct {
	Logger      *zap.Logger
	Tracer      trace.Tracer
	MongoClient *mongo.Client
	NatsConn    *nats.Conn
	RedisClient *redis.Client
}

func (a *App) Start() error {
	server := gin.Default()
	router := router.NewRouter(a.Tracer)

	router.AddRoutes(server)

	return server.Run(":8080")
}
