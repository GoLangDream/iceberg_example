package main

import (
	"github.com/GoLangDream/iceberg/web"
	. "routes/web"
	_ "routes/web/controllers"
	_ "routes/web/controllers/admin"
)

func main() {
	server := web.CreateServer(&RoutesApplication{})
	server.Start()
}
