package routers

import (
	"beego_test/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/happy", &controllers.HappyController{})
	beego.Router("/sad", &controllers.SadController{})
	beego.Router("/one", &controllers.OneController{})
	// beego.AutoRouter(&controllers.OneController{})
}
