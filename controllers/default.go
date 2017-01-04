package controllers

import (
	"stratacompare/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Layout = "site-layout.html"
	c.TplName = "index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["PageHeader"] = "page-header.html"
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Body"] = "index.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.LayoutSections["Scripts"] = "scripts.html"
	c.LayoutSections["PageScripts"] = "page-scripts.html"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.html"
}

func (c *MainController) AllManagers() {

	c.Layout = "site-layout.html"
	c.TplName = "all-managers.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["PageHeader"] = "page-header.html"
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Body"] = "all-managers.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.LayoutSections["Scripts"] = "scripts-all-managers.html"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.html"
}

func (c *MainController) Financials() {

	c.Layout = "site-layout.html"
	c.TplName = "financials.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["PageHeader"] = "page-header.html"
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Body"] = "financials.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.LayoutSections["Scripts"] = "scripts-all-managers.html"
	v := models.Getfinancialitems()
	c.Data["d"] = v
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.html"
}

func (c *MainController) Services() {

	c.Layout = "site-layout.html"
	c.TplName = "services.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["PageHeader"] = "page-header.html"
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Body"] = "services.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.LayoutSections["Scripts"] = "scripts-all-managers.html"
	v := models.Getserviceitems()
	c.Data["d"] = v
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.html"
}

func (c *MainController) ShortCodes() {

	c.Layout = "site-layout.html"
	c.TplName = "shortcodes.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["PageHeader"] = "page-header.html"
	c.LayoutSections["Header"] = "header.html"
	c.LayoutSections["Body"] = "shortcodes.html"
	c.LayoutSections["Footer"] = "footer.html"
	c.LayoutSections["Scripts"] = "scripts-shortcodes.html"
	v := models.Getserviceitems()
	f := models.Getfinancialitems()
	c.Data["services"] = v
	c.Data["fins"] = f

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.html"
}
