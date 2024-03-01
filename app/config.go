package app

type AppConfig struct {
	MongoURL      string `envconfig:"MONGO_URL" default:"mongodb://user:password@localhost:27017"`
	MongoPassword string `envconfig:"MONGO_PASSWORD" default:"MgTBMhVz2FBOH2rE"`
	MongoUsername string `envconfig:"MONGO_USERNAME" default:"taxi-app"`
	NatsUrl       string `envconfig:"NATS_URL" default:"http://localhost:4222"`
	NatsTopics    string `envconfig:"NATS_TOPICS" default:"trip.created,trip.updated"`
	RedisUrl      string `envconfig:"REDIS_URL" default:"localhost:6379"`
}
