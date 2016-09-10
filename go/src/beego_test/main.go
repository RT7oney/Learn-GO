package main

import (
	"beego_test/models"
	_ "beego_test/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true) //自动建表的功能
	beego.Run()
}
