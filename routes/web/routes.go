package web

import "github.com/GoLangDream/iceberg/web"

func RouterDraw(r *web.Router) {
	r.GET("/hello", "home#index")
	r.GET("/get_first_user", "home#get_first_user")

	r.GET("/set_session", "home#set_session")
	r.GET("/get_session", "home#get_session")

	r.GET("/set_cookie", "home#set_cookie")
	r.GET("/get_cookie", "home#get_cookie")

	r.GET("/get_params/:id", "home#get_params")
	r.GET("/get_query", "home#get_query")

	r.Namespace("admin", func(admin *web.Router) {
		admin.GET("/user", "user#index")
	})

	r.Scope("scope_test", func(scope *web.Router) {
		scope.GET("/hello", "home#index")
	})
}
