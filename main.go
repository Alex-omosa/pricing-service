package main

import (
	"runtime"

	app "github.com/Alex-omosa/go-shared/app"
	db "github.com/Alex-omosa/go-shared/db"
	trippb "github.com/Alex-omosa/protos/protos/trip-service"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const mongoURI = "mongodb+srv://trip-service:j0zUuoeXMGHMmtrj@taxi-app.bbjceer.mongodb.net/?retryWrites=true&w=majority&appName=taxi-app"

var logger *zap.Logger
var mongoClient *mongo.Client
var redisClient *redis.Client
var natsConn *nats.Conn

func init() {

	// Initialize the logger
	logger = app.InitializeLogger()

	mongoClient, _ = db.CreateMongodbClient(mongoURI)

	natsConn = app.ConnectNats("nats://localhost:4222")

}

type App struct {
	Logger      *zap.Logger
	MongoClient *mongo.Client
	RedisClient *redis.Client
}

func main() {
	app := App{
		Logger:      logger,
		MongoClient: mongoClient,
		RedisClient: redisClient,
	}

	srv, err := micro.AddService(natsConn, micro.Config{
		Name:        "trip-service",
		Version:     "1.0.0",
		Description: "This is a trip service",
	})
	if err != nil {
		logger.Error("Error adding service", zap.Error(err))
	}

	srv.AddEndpoint("payment", micro.HandlerFunc(func(msg micro.Request) {
		trip := trippb.Trip{}

		proto.Unmarshal(msg.Data(), &trip)
	}))
	app.Logger.Info("Application started successfully")
	runtime.Goexit()
}
