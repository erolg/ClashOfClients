package clashofclients

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

func (coc *ClashOfClients) CheckNickName(nick string) bool {
	p := coc.redisPool.Get()
	defer p.Close()

	result, _ := redis.Bool(p.Do("SISMEMBER", "nicknames", nick))
	return result
}

func (coc *ClashOfClients) CreateGameStore(nick string, email string) (string, string) {
	p := coc.redisPool.Get()
	defer p.Close()

	hash := nick + "_" + strconv.FormatInt(time.Now().Unix(), 10)

	result, _ := redis.String(p.Do("HMSET", hash, "move_counter", 0, "email", email))

	return hash, result
}

func (coc *ClashOfClients) CreateRedisPool() {
	coc.redisPool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":"+strconv.Itoa(coc.Cfg.RedisPort))
			if err != nil {
				return nil, err
			}
			if coc.Cfg.RedisPassword != "" {
				if _, err := c.Do("AUTH", coc.Cfg.RedisPassword); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
