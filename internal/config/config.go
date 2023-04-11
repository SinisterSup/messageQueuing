package config

type Config struct {
	MongoDB  MongoDBConfig
	RabbitMQ RabbitMQConfig
}

type MongoDBConfig struct {
	URI string
}

type RabbitMQConfig struct {
	URI string
}

func LoadConfig() *Config {
	return &Config{
		MongoDB: MongoDBConfig{
			URI: "mongodb://localhost:27017",
		},
		RabbitMQ: RabbitMQConfig{
			URI: "amqp://guest:guest@localhost:5672/",
		},
	}
}
