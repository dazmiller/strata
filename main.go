package main

import (
	_ "stratacompare/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:stratacompare.db")
}

func main() {
	orm.Debug = true
	beego.Run()
}
