package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.TplName = "OneAdmin.html"
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
}
