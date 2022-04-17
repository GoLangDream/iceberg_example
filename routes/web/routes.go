package web

import (
	. "github.com/GoLangDream/iceberg/web"
)

func init() {
	ApplicationRouterDraw = func(r Router) {
		r.GET("/hello", "home#index")

		r.GET("/set_session", "home#set_session")
		r.GET("/get_session", "home#get_session")

		r.GET("/set_cookie", "home#set_cookie")
		r.GET("/get_cookie", "home#get_cookie")
	}
}
