package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loginInfo/models"
)

type JSONStruct struct {
	Id int
	Name string
	Price string
	Home string
	Date string
}

type OrderSearchController struct {
	beego.Controller
}

func(this *OrderSearchController)OrderSearch(){
	name := this.Input().Get("name")
	beego.Info(name)

	o := orm.NewOrm()
	goods := models.OrderInfo{}
	goods.Name = name

	err := o.Read(&goods,"Name")
	if err!=nil{
		beego.Info("查询失败",err)
		this.Ctx.WriteString("查询失败")
		return
	}

	mystruct := &JSONStruct{goods.Id,goods.Name,goods.Price,goods.Home,goods.Date}
	this.Data["json"] = mystruct
	this.ServeJSON()
	beego.Info(goods.Id,goods.Name,goods.Price,goods.Home,goods.Date)
	this.Ctx.WriteString("查询成功")
}

func(this *OrderSearchController)OrderUpdate(){
	name := this.Input().Get("name")
	price := this.Input().Get("price")
	home := this.Input().Get("home")
	date := this.Input().Get("date")
	beego.Info(name,price,home,date)
	beego.Info(name)
	o := orm.NewOrm()
	goods := models.OrderInfo{}
	goods.Name = name
	err := o.Read(&goods,"Name")
	if err!=nil{
		beego.Info("查询失败")
		this.Ctx.WriteString("没有该商品")
		return
	}
	goods.Name = name
	goods.Price = price
	goods.Home = home
	goods.Date = date
	_,err = o.Update(&goods)
	if err != nil{
		beego.Info("更新错误",err)
		this.Ctx.WriteString("更新错误")
		return
	}
	this.Ctx.WriteString("更新成功")

}

func(this *OrderSearchController)OrderDelete(){
	gname := this.Input().Get("name")
	//hello := this.Input().Get("id")
	//将字符串转换为整形
	//phoneNumber, _ := strconv.Atoi(hello)
	//beego.Info(name)
	o := orm.NewOrm()
	goods := models.OrderInfo{}
	//goods.Name = gname
	//goods.Id = phoneNumber
	goods.Name = gname
	err := o.Read(&goods,"Name")
	if err == nil{
		if num, err := o.Delete(&goods); err == nil {
			fmt.Println(num)
		}

	}

}

//func(this *OrderSearchController)One(){
//	var user User
//	use := models.UserInfo{}
//
//	o := orm.NewOrm()
//	err := o.QueryTable("user_info").Filter("user","123").One(&user)
//	if err == orm.ErrMultiRows {
//		// 多条的时候报错
//		fmt.Printf("Returned Multi Rows Not One")
//	}
//	if err == orm.ErrNoRows {
//		// 没有找到记录
//		fmt.Printf("Not row found")
//	}
//}