package modules

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewCurrencies(client *redis.Client) Currencies {
	return Currencies{client: client}
}

type Currencies struct {
	client *redis.Client
}

func (u *Currencies) Set(abbreviation, price string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	err := u.client.Set(ctx, abbreviation, price, 0).Err()
	return err
}

func (u *Currencies) Get(abbreviation string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	price, err := u.client.Get(ctx, abbreviation).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}

	return price, nil
}
