package controllers

import (
	"github.com/astaxie/beego"
)

type HappyController struct {
	beego.Controller
}

type SadController struct {
	beego.Controller
}

func (happy *HappyController) Get() { //重载了一下他的原来的get方法
	happy.Ctx.WriteString("I'm happy")
}

func (sad *SadController) Get() {
	sad.Ctx.WriteString("I'm sad")
}
