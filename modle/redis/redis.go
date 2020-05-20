package redis

import (
	"baserun/conf"
	red "github.com/gomodule/redigo/redis"
	"time"
)

type Redis struct {
	pool *red.Pool
}

var redis *Redis

func init() {
	initRedis()

}

func initRedis() {
	redis = new(Redis)
	redis.pool = &red.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				conf.Config.Redis.Server+":"+conf.Config.Redis.Port,
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				red.DialDatabase(0),
				//red.DialPassword(""),
			)
		},
	}

}

// 执行redis值
func Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}

//获取 redis的结果值
func GetValue(value interface{}, err error) (str string, err2 error) {
	str, err2 = red.String(value, err)
	return
}
