package main

import (
	"github.com/GoLangDream/iceberg/web"
	_ "routes/web"
	_ "routes/web/controllers"
)

func main() {
	web.Start()
}
