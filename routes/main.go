package main

import (
	"github.com/GoLangDream/iceberg"
	_ "routes/initializers"
	_ "routes/web/controllers"
	_ "routes/web/controllers/admin"
)

func main() {
	iceberg.InitApplication()
	iceberg.StartApplication()
}
