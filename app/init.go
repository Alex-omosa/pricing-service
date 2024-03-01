package app

import (
	"context"
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"go.uber.org/zap"
)

func InitializeLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}

func InitializeConfig(config interface{}) {
	err := envconfig.Process("", config)
	if err != nil {
		fmt.Println("Failed to process env var", zap.Error(err))
	}
}

func InitializeTracer() *sdktrace.TracerProvider {
	tp, err := InitTracer()
	if err != nil {
		fmt.Println("Failed to initialize tracer", zap.Error(err))
	}
	return tp
}

// CreateMongodbClient connects to the MongoDB instance
// and returns a reference to the client object or an error
func CreateMongodbClient(mongodbUri string) (*mongo.Client, error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(mongodbUri).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		// logger.Fatal(err.Error())
		println(err.Error())
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		// logger.Fatal(err.Error())
		println(err.Error())
		return nil, err
	}

	// fmt.Println("Successfully connected to MongoDB!")
	return client, nil
}

// ConnectNats connects to the NATS instance
func ConnectNats(natsUrl string) (*nats.Conn, error) {
	// Connect to NATS server
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
