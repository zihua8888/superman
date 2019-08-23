package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loginInfo/models"
	"strconv"
)

type OrderAddController struct {
	beego.Controller
}

// 校验token是否有效
//func CheckToken(token string) bool {
//	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
//		return key, nil
//	})
//	if err != nil {
//		fmt.Println("parase with claims failed.", err)
//		return false
//	}
//	return true
//}


func(this *OrderAddController)OrderAdd(){
	user := this.GetSession("token")
	fmt.Println(this.GetSession("token"))
	if user == nil{
		this.Ctx.WriteString("请登入!")
	}
	id := this.Input().Get("id")
	number, _ := strconv.Atoi(id)
	name := this.Input().Get("name")
	price := this.Input().Get("price")
	home := this.Input().Get("home")
	date := this.Input().Get("date")
	beego.Info(name,price,home,date)
	o := orm.NewOrm()

	order := models.OrderInfo{}
	order.Id = number
	order.Name = name
	order.Price = price
	order.Home = home
	order.Date = date

	_,err :=o.Insert(&order)
	if err != nil{
		beego.Info("添加失败",err)
		return
	}
	this.Ctx.WriteString("添加商品成功")
}

