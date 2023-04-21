package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangyiming748/goini"
	"github.com/zhangyiming748/pretty"
	"golang.org/x/exp/slog"
	"strings"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

const (
	configurePath = "./conf.ini"
)

//var MyEngine *xorm.Engine

func GetMysqlEngine() *xorm.Engine {
	conf := goini.SetConfig(configurePath)
	user, _ := conf.GetValue("mysql", "user")
	passwd, _ := conf.GetValue("mysql", "passwd")
	database, _ := conf.GetValue("mysql", "database")
	src := strings.Join([]string{user, ":", passwd, "@/", database, "?charset=utf8"}, "")
	pretty.P(src)
	MyEngine, err := xorm.NewEngine("mysql", src)
	if err != nil {
		slog.Error("创建数据库引擎失败", slog.Any("错误信息", err))
		return nil
	}
	err = MyEngine.Ping()
	if err != nil {
		return nil
	} else {
		slog.Info("创建数据库引擎成功")
	}
	MyEngine.SetMapper(names.GonicMapper{})
	return MyEngine
}

//func GetSession() *xorm.Session {
//	return MyEngine.NewSession()
//}
