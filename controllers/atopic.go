package controllers

import (
	"github.com/astaxie/beego"
	"ONE/models"
	"net/http"
)

type ATopicController struct {
	beego.Controller
	http.Request
}

func (this *ATopicController) Get() {
	this.TplName = "AdminTopic.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true

	topics, err := models.GetAllTopics("", false, true)

	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *ATopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	// content := this.Request.FormValue("editormd-html-code")
	category := this.Input().Get("category")
	label := this.Input().Get("label")
	summary := this.Input().Get("summary")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, label, summary, content)
	} else {
		err = models.ModifyTopic(tid, title, category, summary, content, label)
	}

	if err != nil {
		beego.Error(err)
	}

	// this.Redirect("/topic", 302)
	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ATopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_add.html"
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}

func (this *ATopicController) View() {
	this.TplName = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = this.Ctx.Input.Params()["0"]

	replies, err := models.GetAllReplies(this.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)

}

func (this *ATopicController) Modify() {
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}

func (this *ATopicController) Delete() {
	this.TplName = "topic.html"
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	// err := models.DeleteTopic(this.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
	return
}
