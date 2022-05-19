package admin

import "github.com/GoLangDream/iceberg/web"

func init() {
	web.RegisterController(UserController{})
}

type UserController struct {
	*web.BaseController
}

func (c *UserController) Init() {
	println("init controller")

	c.BeforeActon("Index", c.checkAction)
}

func (c *UserController) Index() {
	c.Text("controller is " + c.ControllerName() + " action is " + c.ActionName())
}

func (c *UserController) checkAction() {
	println("check action")
}
