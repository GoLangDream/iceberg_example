package admin

import "github.com/GoLangDream/iceberg/web"

func init() {
	web.RegisterController(UserController{})
}

type UserController struct {
	*web.BaseController
}

func (c UserController) Index() {
	c.Text("admin user index action")
}
