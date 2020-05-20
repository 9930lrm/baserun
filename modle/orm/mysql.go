package orm

import (
	"baserun/conf"
	"baserun/modle/logs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//初始化数据库

var MysqlDb *gorm.DB
var MysqlDbErr error

//func initMysql()  {
//    dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf.Config.Database.DBuser, conf.Config.Database.DBPasswd, conf.Config.Database.Server, conf.Config.Database.Port, conf.Config.Database.DBName, "utf8")
//
//    // 打开连接失败
//    MysqlDb, MysqlDbErr = gorm.Open()
//    //defer MysqlDb.Close();
//    if MysqlDbErr != nil {
//        log.Println("dbDSN: " + dbDSN)
//        panic("数据源配置不正确: " + MysqlDbErr.Error())
//    }
//
//    // 最大连接数
//    MysqlDb.SetMaxOpenConns(100)
//    // 闲置连接数
//    MysqlDb.SetMaxIdleConns(20)
//    // 最大连接周期
//    MysqlDb.SetConnMaxLifetime(100*time.Second)
//
//    if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
//        panic("数据库链接失败: " + MysqlDbErr.Error())
//    }
//
//fmt.Printf("数据库连接成功。")
//
//}

// 初始化mysql
func init() {
	//dbDSN := fmt.Sprintf("root:oamysqldev@(10.10.20.249:10100)/default?charset=utf8")
	dbDSN := fmt.Sprintf(conf.Config.Database.DBuser + ":" + conf.Config.Database.DBPasswd + "@(" + conf.Config.Database.Server + ":" + conf.Config.Database.Port + ")/" + conf.Config.Database.DBName + "?charset=" + conf.Config.Database.Charset)
	MysqlDb, MysqlDbErr = gorm.Open("mysql", dbDSN)
	if MysqlDbErr != nil {
		logs.LogErr(MysqlDbErr)
		return
	}
	logs.LogInfo("db connet success.")
}
