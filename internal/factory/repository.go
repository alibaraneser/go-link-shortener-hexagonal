package factory

import (
	redis2 "github.com/alibaraneser/go-link-shortener-hexagonal/internal/adapters/repository/redis"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/ports"
	"log"
	"os"
)

func Repository() ports.RedirectRepository {
	switch os.Getenv("DB_TYPE") {
	case "redis":
		redisURL := os.Getenv("DB_PATH")
		repo, err := redis2.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
		// can be added mongo db, MySQL, PostreSQL, ...
	}
	return nil
}
