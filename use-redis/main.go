package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go-learn/use-redis/redistool"
)

func main() {
	conn := redistool.RedisPool.Get()
	defer conn.Close()

	_, e := conn.Do("SET", "username", "ft")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	r, e := redis.String(conn.Do("GET", "username"))
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(r)
}
