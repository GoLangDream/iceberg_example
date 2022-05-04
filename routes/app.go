package main

import (
	"github.com/GoLangDream/iceberg/web"
	_ "routes/web"
	_ "routes/web/controllers"
	_ "routes/web/controllers/admin"
)

func main() {
	web.Start()
}
