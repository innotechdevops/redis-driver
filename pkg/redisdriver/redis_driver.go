package redisdriver

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const DefaultPort = "6379"
const DefaultDatabase = 0

// Config is a model for connect RedisDB
type Config struct {
	Pass         string
	Host         string
	DatabaseName string
	Port         string
}

// RedisDriver is the interface
type RedisDriver interface {
	Connect() *redis.Client
}

type redisDB struct {
	Conf Config
}

func (r *redisDB) Connect() *redis.Client {
	db, err := strconv.Atoi(r.Conf.DatabaseName)
	if err != nil {
		db = DefaultDatabase
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Conf.Host, r.Conf.Port),
		Password: r.Conf.Pass,
		DB:       db,
	})
	log.Print("Redis Connected.")
	return client
}

// New for create redis driver
func New(config Config) RedisDriver {
	return &redisDB{
		Conf: config,
	}
}

// ConfigEnv for create redis driver
func ConfigEnv() Config {
	return Config{
		Host:         os.Getenv("REDIS_HOST"),
		Pass:         os.Getenv("REDIS_PASS"),
		DatabaseName: os.Getenv("REDIS_DATABASE"),
		Port:         os.Getenv("REDIS_PORT"),
	}
}
