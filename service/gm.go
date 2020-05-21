package service

import (
	"baserun/Entitys"
	"baserun/modle/orm"
)

func GetAllUsers() {

	var users []Entitys.User
	orm.MysqlDb.Find(&users)

}
