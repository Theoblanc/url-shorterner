package infrastructure

import (
	"github.com/Theoblanc/url-shortener/config"
	"github.com/go-redis/redis"
)

// GetRedisCilent resdis client
func GetRedisCilent(config config.Interface) *redis.Client {
	address := config.Redis().Address()
	password := config.Redis().Password()

	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})
}
