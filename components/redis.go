package components

import (
	"codetube.cn/service-source/config"
	"github.com/go-redis/redis/v8"
)

// Redis 连接
var Redis = newRedis()

// CommonRedis 公共缓存连接
var CommonRedis *redis.Client

// CourseRedis 课程缓存连接
var CourseRedis *redis.Client

// 数据库连接列表
type redisConnection struct {
	Common *redis.Client // 公共缓存使用的连接
	Course *redis.Client // 课程缓存使用的连接
}

// 创建数据库连接列表
func newRedis() *redisConnection {
	return &redisConnection{
		Common: redis.NewClient(&redis.Options{
			Addr:     config.ServiceConfig.Redis["common"].Addr(),
			Password: config.ServiceConfig.Redis["common"].Password,
			DB:       config.ServiceConfig.Redis["common"].Db,
		}),
		Course: redis.NewClient(&redis.Options{
			Addr:     config.ServiceConfig.Redis["course"].Addr(),
			Password: config.ServiceConfig.Redis["course"].Password,
			DB:       config.ServiceConfig.Redis["course"].Db,
		}),
	}
}

// RedisInit 初始化Redis连接
func (c *redisConnection) RedisInit() (err error) {
	CommonRedis = c.Common
	CourseRedis = c.Course
	//其他连接...
	return
}
