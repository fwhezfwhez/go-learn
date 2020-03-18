package redistool

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool


func init() {
	RedisPool = newPool("localhost:6379", "", 0)
	conn := RedisPool.Get()
	defer conn.Close()
	rs, e := redis.String(conn.Do("ping"))
	if e != nil {
		panic(e.Error())
	}
	fmt.Println(rs)
}
