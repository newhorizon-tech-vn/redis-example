package services

import (
	"context"
	"time"

	"github.com/newhorizon-tech-vn/redis-example/cache"
	"github.com/newhorizon-tech-vn/redis-example/models"
	"github.com/newhorizon-tech-vn/redis-example/models/entities"
	"github.com/newhorizon-tech-vn/redis-example/pkg/log"
	"github.com/newhorizon-tech-vn/redis-example/setting"
)

type UserService struct {
	UserId int
}

func (s *UserService) GetUser() (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancel()

	// try to get from cache
	user, err := cache.GetUser(ctx, s.UserId)
	if err == nil {
		return user, nil
	}

	// try to get from storage
	user, err = models.GetUser(s.UserId)
	if err != nil {
		return nil, err
	}

	// update cache
	// if udpate cache failed, we only write log (and push error metrics), then still continue return user data
	if err = cache.SetUser(ctx, user, setting.Setting.Redis.DefaultExpireTime); err != nil {
		log.Error("update cache failed", "user_id", s.UserId, "error", err)
	}

	return user, nil
}
