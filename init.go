package clashofclients

import (
	"github.com/garyburd/redigo/redis"
)

type Config struct {
	Name          string
	RedisPort     int
	RedisPassword string
	SessionSecret string
	SessionMaxAge int
}

type game struct {
	session string
}

var defaults Config = Config{
	Name:          "clashofclients",
	RedisPort:     6379,
	SessionSecret: "secret123",
	RedisPassword: "",
	SessionMaxAge: 2592000,
}

type ClashOfClients struct {
	Cfg       Config
	redisPool *redis.Pool
	game      game
}

func New(configs ...Config) *ClashOfClients {

	var c Config
	if len(configs) == 0 {
		c = Config{}
	} else {
		c = configs[0]
	}

	coc := ClashOfClients{
		Cfg: c,
	}

	coc.prepareConfig()

	return &coc
}

func (coc *ClashOfClients) prepareConfig() {

	if len(coc.Cfg.Name) == 0 {
		coc.Cfg.Name = defaults.Name
	}
	if len(coc.Cfg.SessionSecret) == 0 {
		coc.Cfg.SessionSecret = defaults.SessionSecret
	}
	if len(coc.Cfg.RedisPassword) == 0 {
		coc.Cfg.RedisPassword = defaults.RedisPassword
	}
	if coc.Cfg.RedisPort == 0 {
		coc.Cfg.RedisPort = defaults.RedisPort
	}
	if coc.Cfg.SessionMaxAge == 0 {
		coc.Cfg.SessionMaxAge = defaults.SessionMaxAge
	}
}
