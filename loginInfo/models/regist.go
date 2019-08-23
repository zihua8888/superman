package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Id int
	User string
	Pwd string
}

type OrderInfo struct {
	Id int
	Name string
	Price string
	Home string
	Date string
}


func init(){
	orm.RegisterDataBase("default","mysql","root:1234@tcp(127.0.0.1:3306)/logininfo?charset=utf8")
	orm.RegisterModel(new(UserInfo))
	orm.RegisterModel(new(OrderInfo))
	orm.RunSyncdb("default",false,true)
}

