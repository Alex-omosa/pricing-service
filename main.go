package main

import (
	"context"
	"log"

	"github.com/Alex-omosa/pricing-service/app"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"

	"go.uber.org/zap"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var tracer = otel.Tracer("gin-server")
var traceProvider *sdktrace.TracerProvider
var logger *zap.Logger
var appconfig app.AppConfig
var mongoClient *mongo.Client
var natsConn *nats.Conn
var redisClient *redis.Client

func init() {
	app.InitializeConfig(&appconfig)                          // Initialize the app configuration
	logger = app.InitializeLogger()                           // Initialize the logger
	traceProvider = app.InitializeTracer()                    // Initialize the tracer
	mongo, err := app.CreateMongodbClient(appconfig.MongoURL) // Initialize the mongo client
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", zap.Error(err))
	}
	mongoClient = mongo

	nats, err := app.ConnectNats(appconfig.NatsUrl) // Initialize the nats connection
	if err != nil {
		logger.Fatal("Failed to connect to Nats", zap.Error(err))
	}

	natsConn = nats

	//-----------------Initialize the redis client----------------
	redisClient = redis.NewClient(&redis.Options{ //Redis client
		Addr:     appconfig.RedisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if redisClient == nil {
		logger.Fatal("Failed to connect to Redis")
	}

}

func main() {
	app := app.App{
		Logger:      logger,
		Tracer:      tracer,
		MongoClient: mongoClient,
		NatsConn:    natsConn,
		RedisClient: redisClient,
	}

	err := app.Start()
	if err != nil {
		logger.Error("Failed to initialize app", zap.Error(err))
	}

	cleanup()
}

func cleanup() {
	logger.Info("Shutting down...")
	if err := traceProvider.Shutdown(context.Background()); err != nil {
		log.Printf("Error shutting down tracer provider: %v", err)
	}
}
