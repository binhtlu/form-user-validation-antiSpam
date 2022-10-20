package tests

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func returnRdb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379", // use default Addr
		Password:    "",               // no password set
		DB:          0,                // use default database
		PoolTimeout: time.Minute,      // since we user transaction so it can take a long time
	})
	return rdb
}

func incRequestCount(key string, rateLimit int64, second int64) error {
	err := returnRdb().Watch(func(tx *redis.Tx) error {
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
		//clientIp := c.ClientIP()
		//key := "RATE_LIMIT_COUNT_" + clientIp
		err := incRequestCount("key", rateLimit, second)
		if err != nil {
			c.AbortWithStatus(403)
			c.JSON(http.StatusBadRequest, "spam")
			return
		}
		c.Next()
	}
}

func TestQP_AnswerAfterSelectJoinWithTokenCreate(t *testing.T) {
	var b bool
	for i := 1; i < 5; i++ {
		err := incRequestCount("key", 3, 10)
		if err != nil {
			b = false
		}
		if err == nil {
			b = true
		}
	}
	if b == false {
		fmt.Println("that bai")
	} else {
		fmt.Println("thanh cong")
	}

}
