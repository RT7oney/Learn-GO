package models

//如何使用orm
//先创建一个结构，然后把结构提交给orm，orm就可以去创建一个数据库
import (
	"github.com/Unknwon/com" //一个通用函数包
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" //只执行这个包里面的初始化函数，不进行这个包里面的其他函数的调用，因为他是一个驱动，再初始化函数中进行驱动的注册
	"os"
	"path"
	"time"
)

const (
	_DB_NAME        = "data/beego_test.db" //在data目录下，但是这个.db不一定存在，所以应该有一个检查这个数据库是否存在的方法
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string    //会默认设置成255长度的
	Created         time.Time `orm:"index"` //使用索引，tag，在反射的时候使用，可以得到tag作为一个说明来使用，而且反射一定要作为导出来使用，如果结构使用orm使用导出字段也就是首字母大写
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Update          time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}
type Comment struct {
	Id       int64
	Tid      int64
	Username string
	Content  string `orm:"size(1000)"`
}

func RegisterDB() {
	//由主进程去调用，告诉orm进行一个创建，注册模型，注册驱动，注册数据库了
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	//orm 需要先注册模型，什么是模型呢，就是刚才的两个结构体
	orm.RegisterModel(new(Category), new(Topic)) //传进去一个指针
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}
