package main

import (
	"baserun/modle/logs"
	"baserun/modle/orm"
	"baserun/modle/redis"
)

func main() {
	logs.LogInfo("app starting....")

	redisTest()

	defer orm.MysqlDb.Close()
}

// ok
func redisTest() {
	//redis.Exec("set","hello","world")
	result, err := redis.Exec("get", "hello")
	if err != nil {
		logs.LogErr(err)
	}
	str, _ := redis.GetValue(result, err)
	logs.LogInfo(str)
}

type aa struct {
}
