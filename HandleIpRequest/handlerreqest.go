package HandleIpRequest

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var rdb *redis.Client
var counter Counter
var c *gin.Context

func incRequestCount(addressIP string) error {
	err := rdb.Watch(
		func(tx *redis.Tx) error {
			counter.Add(addressIP, 1)
			_ = tx.SetNX(addressIP, 0, time.Duration(1)*time.Second)
			count, err := counter.Get(addressIP)

			if count > 3 {
				err = errors.New("rate limited")
			}
			if err != nil {
				return err
			}
			return nil
		}, addressIP)
	return err
}
func useRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		key := "RATE_LIMIT_COUNT_" + clientIp
		err := incRequestCount(key)
		if err != nil {
			c.AbortWithStatus(403)
			return
		}
		c.Next()
	}
}

/*
	func (c *Counter) HandleRequest(addr string) (int64, error) {
		err := errors.New("")
		c.Add(addr, 1)
		count, _ := c.Get(addr)
		if count > 3 {
			return count, err
		}
		return count, nil

}
*/
/*func HandleRequest(tx *redis.Tx) error {
	counter.Add(addressIP, 1)
	_ = tx.SetNX(addressIP, 0, time.Duration(1)*time.Second)
	count, err := counter.Get(addressIP)
	if count > 3 {

	}
}*/
