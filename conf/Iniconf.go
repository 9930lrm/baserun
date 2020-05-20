package conf

import (
	"baserun/Entitys"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

var Config *Entitys.Fconf

func init() {
	fmt.Println("开始读取配置文件.")

	//    初始化配置文件
	if _, err := toml.DecodeFile("D:\\00.OneDrve\\OneDrive\\goProject\\01.base\\baserun\\conf\\app.toml", &Config); err != nil {
		log.Fatal(err)
	}
	fmt.Println("配置文件读取成功.")

}
