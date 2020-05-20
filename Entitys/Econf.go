package Entitys

// 类型 不能空
type Fconf struct {
	Database Database
	Redis    Redis
	Admin    Admin
	Logs     Logs
}

type Database struct {
	Server         string
	Port           string
	Connection_max string
	DBName         string
	DBuser         string
	DBPasswd       string
	Charset        string
}

type Redis struct {
	Server string
	Port   string
	passwd string
}

type Admin struct {
	Admin  string
	Passwd string
}

type Logs struct {
	Logpath string
}
