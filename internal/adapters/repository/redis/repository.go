package redis

import (
	"context"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/domain"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/ports"
	"github.com/go-redis/redis/v8"
	errs "github.com/pkg/errors"
	"log"
)

var ctx = context.Background()

type redisRepository struct {
	client *redis.Client
}

func (r *redisRepository) Find(code string) (*domain.Redirect, error) {
	val, err := r.client.Get(ctx, code).Result()
	if err != nil {
		return nil, errs.Wrapf(domain.RedirectNotFound, "redis.Repository.Find()")
	}

	redirect := &domain.Redirect{Code: code, URL: val}

	return redirect, nil
}

func (r *redisRepository) Save(redirect *domain.Redirect) error {
	_, err := r.client.Set(ctx, redirect.Code, redirect.URL, 0).Result()
	if err != nil {
		return errs.Wrapf(domain.RedirectSaveError, "redis.Repository.Save()")
	}

	return nil
}

func NewRedisRepository(url string) (ports.RedirectRepository, error) {
	repo := &redisRepository{}
	client, _ := newRedisClient(url)

	log.Println("connected to redis")
	repo.client = client
	return repo, nil
}

func newRedisClient(url string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		return nil, errs.Wrapf(domain.ConnectionError, "redis.Repository")
	}

	return rdb, err
}