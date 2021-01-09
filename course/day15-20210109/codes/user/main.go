package main

import (
	"fmt"
	"log"
	"strings"
	_ "user/routers"
	"user/services"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

func main() {
	fmt.Printf("%q\n", beego.AppConfig.Strings("Token"))
	fmt.Println(beego.AppConfig.String("RunMode"))
	fmt.Println(beego.AppConfig.String("db::Db_Host"))
	dsn := beego.AppConfig.String("db::dsn")
	fmt.Print(dsn)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err)
	}

	orm.RunSyncdb("default", false, true)

	if user := services.GetUserByName("admin"); user == nil {
		log.Print("创建管理员账号")
		services.AddUser("admin", "123@abc", "西安", true)
	}

	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.Log.FileLineNum = true
	beego.BeeLogger.DelLogger("console")
	beego.SetLogger("file", `{
		"level": 7,
		"filename": "logs/web.log",
		"daily": true,
		"maxdays": 15
	}`)
	beego.AddFuncMap("upper", strings.ToUpper)
	beego.AddFuncMap("kk", func(t string) string {
		return "kk:" + t
	})
	beego.Run()
}
