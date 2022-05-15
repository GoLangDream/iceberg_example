package initializers

import (
	iceberg "github.com/GoLangDream/iceberg/web"
	"os"
	"routes/web"
)

type RoutesApplicationConfig struct {
}

func (app *RoutesApplicationConfig) HomePath() string {
	path, _ := os.Getwd()
	return path
}

func (app *RoutesApplicationConfig) RouterDraw() func(r *iceberg.Router) {
	return web.RouterDraw
}
