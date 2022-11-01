package redisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type RedisClient struct {
	rdb *redis.Client
	ctx context.Context
}

var lock = &sync.Mutex{}

func New() *RedisClient {
	return &RedisClient{}
}

func (rd *RedisClient) GetInstanceRedis() *redis.Client {
	godotenv.Load()
	if rd.rdb == nil {
		addr := fmt.Sprintf("%s:%s", os.Getenv("REDIDSHOST"), os.Getenv("REDISPORT"))
		lock.Lock()
		rd.rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: os.Getenv("REDISPASSWORD"), // no password set
			DB:       0,                          // use default DB
		})
		rd.ctx = context.Background()
		lock.Unlock()
	}
	return rd.rdb
}

func (rd *RedisClient) Get(key string, model interface{}) error {
	redis := rd.GetInstanceRedis()
	renotes, err := redis.Get(rd.ctx, key).Result()
	// fmt.Println(renotes)
	if err != nil {
		return err
	}

	// ttl, err := redis.TTL(rd.ctx, key).Result()
	// fmt.Println(ttl)
	err = json.Unmarshal([]byte(renotes), &model)
	// fmt.Println("return to models", err)
	if err != nil {

		return err
	}

	return nil
}
func (rd *RedisClient) Set(key string, model interface{}, expired time.Duration) error {
	redis := rd.GetInstanceRedis()
	val, err := json.Marshal(&model)
	// fmt.Println(key)
	if err != nil {
		return err
	}

	err = redis.Set(rd.ctx, key, string(val), expired).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rd *RedisClient) Remove(key string) error {
	redis := rd.GetInstanceRedis()
	return redis.Do(rd.ctx, "DEL", key).Err()
}
