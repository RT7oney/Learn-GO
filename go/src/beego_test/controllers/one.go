package controllers

import (
	"github.com/astaxie/beego"
)

type OneController struct {
	beego.Controller
}

func (one *OneController) Get() { //重载了一下他的原来的get方法
	one.Ctx.WriteString("这是九型人格当中的第一种")
}
