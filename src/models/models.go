package models

import (
	"GinDemo/src/pkg/setting"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
	DbErr error
)

type  Model struct {
	ID  int  `Gorm:"primary_key" json:"id"`
	Createon int  `json:"createon"`
	Midifiedon int  `json:"midifiedon"`
}

func  init()  {
	var  (
		err  error
		dbType,dbName,user,password,host,tablePerfix string
	)

	sec,err:= setting.Cfg.GetSection("database")
	if err!=nil{
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}
	dbType =sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePerfix = sec.Key("TABLE_PREFIX").String()

	db,DbErr:= gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if DbErr!=nil{
		log.Println("数据库连接错误")
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePerfix + defaultTableName;
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
func CloseDb()  {
	defer  db.Close()
}
