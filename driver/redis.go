package driver

import "github.com/go-redis/redis"

var RDB *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	initClient()
}
