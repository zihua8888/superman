package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegistController struct {
	beego.Controller
}

func(this *RegistController)Insert(){
	user := this.Input().Get("user")
	pwd := this.Input().Get("pwd")
	beego.Info(user,pwd)

	//o := orm.NewOrm()
	//use := models.UserInfo{}
	//use.User = user
	//use.Pwd = pwd
	//
	//_,err := o.Insert(&use)

	var maps []orm.Params
	_,err:=orm.NewOrm().Raw("insert into user_info (user,pwd) values(?,?)",user,pwd).Values(&maps)

	if err != nil{
		beego.Info("插入失败",err)
		return
	}
	this.Ctx.WriteString("注册成功")
}

