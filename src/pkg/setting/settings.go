package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var   (
	 Cfg *ini.File

	 RunMode string

	 HTTPPort int
	 ReadTimeout time.Duration
	 WriteTimeout time.Duration

	 PageSize int
	 JwtSecret string
)
func  LoadBase()  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func  LoadApp()  {
	sec,err:=Cfg.GetSection("app")
	if err!=nil{
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func  init()  {
	var error error
	Cfg,error = ini.Load("F:\\GinDemo\\src\\conf\\app.ini")
	if error!=nil{
		log.Fatal(2,"Fail to parse 'conf/app.ini':%v",error)
	}
	LoadBase()
	LoadApp()
	LoadServer()
}













