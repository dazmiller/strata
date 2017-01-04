package models

import "github.com/astaxie/beego/orm"

type ItemsFin struct {
	Per  string `orm:"column(Per);null"`
	Cost string `orm:"column(Cost);null"`
	ID   int64  `orm:"column(rowid);null"`
	Desc string `orm:"column(Desc);null"`
}

type ItemsService struct {
	Type string `orm:"column(Type);null"`
	Cost string `orm:"column(Cost);null"`
	ID   int64  `orm:"column(rowid);null"`
	Desc string `orm:"column(Desc);null"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(ItemsFin), new(ItemsService))
}

func Getfinancialitems() []ItemsFin {

	var l []ItemsFin
	o := orm.NewOrm()
	qs := o.QueryTable(new(ItemsFin))
	qs.All(&l)
	return l
}

func Getserviceitems() []ItemsService {

	var l []ItemsService
	o := orm.NewOrm()
	qs := o.QueryTable(new(ItemsService))
	qs.All(&l)
	return l
}
