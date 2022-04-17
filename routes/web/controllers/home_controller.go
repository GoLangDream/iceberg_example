package controllers

import (
	"github.com/GoLangDream/iceberg/web"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (h *HomeController) Index() {
	h.Text("hello word")
}

func (h *HomeController) SetSession() {
	h.Session("test_session", "a")
}

func (h *HomeController) GetSession() {
	testSession := h.Session("test_session")
	if testSession == nil {
		h.Text("session 没有初始化")
	} else {
		h.Text("session 的值为" + testSession.(string))
	}
}

func (h *HomeController) SetCookie() {
	h.Cookie("test_cookie", "test_cookie_value")
}

func (h *HomeController) GetCookie() {
	testCookie := h.Cookie("test_cookie")
	if testCookie == "" {
		h.Text("testCookie 没有初始化")
	} else {
		h.Text("session 的值为" + testCookie)
	}
}
