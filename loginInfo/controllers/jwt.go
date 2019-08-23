package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"loginInfo/models"
)



// ------------- 这是封装的类库。理论上应该写在另外一个go文件中，这里偷懒了
type BaseController struct {
	beego.Controller
}

func (this *BaseController) Success(data interface{}) {
	this.Data["json"] = beego.M{"code": 200,"data":data }
	// TODO:问题2：这里可以传参，用途是什么
	//接收UpdateUser函数里面的user参数
	//把要输出的数据this.Data["json"] = beego.M{"code": 200, "data": user}
	//调用ServeJSON()进行渲染，就可以把数据进行JSON序列化输出。
	this.ServeJSON()
	//this.StopRun()
}

func (this *BaseController) Error(code int, msg ...string) {
	this.Data["json"] = beego.M{"code": code, "msg": msg}
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) GetUserFromToken() *models.User {
	tokenStr := this.Ctx.Input.Header("token")
	if tokenStr == "" {
		return nil
	} else {
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return []byte("secret"), nil
		})
		if !token.Valid {
			return nil
		} else {
			var user models.User
			// TODO:问题7：token.Claims.(jwt.MapClaims) 这是什么意思
			//强转成jwt.mapclaims类型
			// TODO:问题8：.One(&user) 这里为什么需要加 “&”符号
			//函数参数是值传递, 如果不加&修改的user的拷贝,不会修改原始的user
			if err := orm.NewOrm().QueryTable(new(models.User)).Filter("Id", token.Claims.(jwt.MapClaims)["userId"]).One(&user); err != nil {
				return nil
			}

			// orm执行原生sql
			// var userList []models.User
			// orm.NewOrm().Raw("select a1.*,a2.* from user a1 left join teacher a2 on a1.teacher_id=a2.id limit ?, ?;", 0, 10).QueryRows(&userList)
			// logs.debug(userList)

			// TODO:问题9：这里为什么需要加 “&”符号
			//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
			return &user

		}
	}
}



