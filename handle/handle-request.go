package handle

import (
	"controllers/User"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var rdb *redis.Client
var c *gin.Context

func incRequestCount(key string, rateLimit int64, second int64) error {
	err := rdb.Watch(func(tx *redis.Tx) error {
		_ = tx.SetNX(key, 0, time.Duration(second)*time.Second)
		count, err := tx.Incr(key).Result()
		if count > rateLimit {
			err = errors.New("rate limited")
		}
		if err != nil {
			return err
		}
		return nil
	}, key)
	return err
}

func useRateLimit(rateLimit int64, second int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		key := "RATE_LIMIT_COUNT_" + clientIp
		err := incRequestCount(key, rateLimit, second)
		if err != nil {
			c.AbortWithStatus(403)
			return
		}
		c.Next()
	}
}

func HandlerRequest() {
	router := gin.Default()
	router.Use(useRateLimit(3, 10))
	router.POST("/User", User.AddNewUser)

	rdb = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379", // use default Addr
		Password:    "",               // no password set
		DB:          0,                // use default database
		PoolTimeout: time.Minute,      // since we user transaction so it can take a long time
	})
	router.Run(":8080")
}
