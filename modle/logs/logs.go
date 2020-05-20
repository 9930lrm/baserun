package logs

import (
	"baserun/conf"
	"log"
	"os"
)

var Logger *log.Logger
var f *os.File
var err error

////初始化 日志
func init() {
	f, err = os.OpenFile(conf.Config.Logs.Logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Logger.Fatal(err)

	}
	Logger = log.New(f, "", log.LstdFlags)

}

// 提示信息
func LogInfo(strinfo string) {
	Logger.SetPrefix("[Info]: ")
	Logger.Printf(" %s", strinfo)
}

// 错误消息
func LogErr(strerr interface{}) {
	Logger.SetPrefix("[Error]: ")
	Logger.Printf(" %s", strerr)
}
