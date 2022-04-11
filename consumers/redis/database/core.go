package database

import (
	"github.com/go-redis/redis/v8"

	"consumer/database/modules"
	"context"
	"fmt"
	"os"
	"strconv"
)

type Redis struct {
	Currencies modules.Currencies
}

func New() (*Redis, func() error, error) {

	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		return nil, nil, fmt.Errorf("parsing 'REDIS_DATABASE' variable: %w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	result, err := rdb.Ping(context.Background()).Result()
	for err != nil {
		return nil, nil, fmt.Errorf("connecting to redis: %w: %s", err, result)
	}

	return &Redis{
		Currencies: modules.NewCurrencies(rdb),
	}, rdb.Close, nil
}
