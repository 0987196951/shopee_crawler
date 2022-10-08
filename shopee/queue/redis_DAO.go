package queue

import (
	"log"

	"github.com/go-redis/redis"
	"shopee.rd/config"
	"shopee.rd/utils"
)

var glb_Redis_DAO Redis_DAO

type Redis_DAO struct {
	Redis_config    config.Redis_config
	Redis_client    *redis.Client
	Logger          log.Logger
	Max_failure     int
	Options         *redis.Options
	MaxRetryBackOff int
	Notifier        utils.Notifier
}
