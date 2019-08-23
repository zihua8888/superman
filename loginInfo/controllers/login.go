package controllers

import (

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"loginInfo/models"
	_ "time"
)

type LoginController struct {
	BaseController
}

func(this *LoginController)Login(){
	user := this.Input().Get("user")
	pawd := this.Input().Get("pwd")
	beego.Info(user)

	o := orm.NewOrm()
	use := models.UserInfo{}
	use.User = user

	err := o.Read(&use,"user")

	if err != nil{
		beego.Info("用户名失败")
		this.Ctx.WriteString("用户名失败")
		return
	}
	if use.Pwd != pawd{
		beego.Info("密码失败")
		this.Ctx.WriteString("密码失败")
		return
	}


	this.Ctx.WriteString("登入成功")
	token, err := generateToken(user)
	// 这里偷懒了，不去数据库查询
	if err != nil {
		this.Error(500, "token生成失败")
	} else {


		this.Success(token)

		//fmt.Println("hhhhhhhhhhhhhhhhh")
		//this.Ctx.SetCookie("token",token,100,"/")
		this.SetSession("token",token)
		//this.DelSession("token")

	}

	//this.Ctx.WriteString("密码错误")

}



//type UserController1 struct {
//	// 这是匿名内部类的写法。也可以理解为继承
//	BaseController
//}

// this可以换成任意词汇，但是习惯上使用this。
// TODO:问题1：为什么要使用 *，与不使用*有什么差别
// 使用 * 修改函数里面的内容调用者跟随变化   不用*调用者不跟随修改的内容而变化
//func (this *LoginController) Login1() {
//	token, err := generateToken("123456")
//	// 这里偷懒了，不去数据库查询
//	if err != nil {
//		this.Error(500, "token生成失败")
//	} else {
//		this.Success(token)
//	}
//}

func (this *LoginController) User() {
	// 从header中 token 标签 获取token内容。当然，这个key可以是任意值（但是前后端需要统一）。
	// 不建议放到 authorization 字段，TODO:问题5：为什么不建议，会产生什么问题
	//会造成请求拦截的问题
	token := this.Ctx.Input.Header("token")
	if len(token) == 0 {
		this.Error(405, "token is empty")
		//beego.Info("hello")
		//print("hello")
		return   // 这个return没有实际功效 TODO:问题4：为什么
	}

	user := this.GetUserFromToken()
	if user == nil {
		this.Error(405, "无对应信息")
		return
	}

	this.Success(user)
}


func generateToken(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user,
		//"exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
	})
	//对应的字符串请自行生成，最后足够使用加密后的字符串
	return token.SignedString([]byte("123"))
}





func(this *ArticleController)HandleSelece(){

	content := this.Input().Get("select")

	if content == ""{
		beego.Info("输入错误")
		return
	}

	o := orm.NewOrm()
	var heroinfo []models.HeroInfo
	o.QueryTable("HeroInfo").RelatedSel("BookInfos").Filter("BookInfos__author",content).All(&heroinfo)
	beego.Info(heroinfo);
	beego.Info(heroinfo[0].Name)
}










