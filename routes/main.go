package main

import (
	"github.com/GoLangDream/iceberg"
	"routes/initializers"
	_ "routes/web/controllers"
	_ "routes/web/controllers/admin"
)

func main() {
	iceberg.InitApplication(&initializers.RoutesApplicationConfig{})
	iceberg.StartApplication()
}
