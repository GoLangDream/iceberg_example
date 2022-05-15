package main

import (
	"github.com/GoLangDream/iceberg"
	"routes/initializers"
	_ "routes/web/controllers"
)

func main() {
	iceberg.InitApplication(&initializers.RoutesApplicationConfig{})
	iceberg.StartApplication()
}
