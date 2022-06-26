package initializers

import (
	iceberg "github.com/GoLangDream/iceberg/web"
	"routes/web"
)

func init() {
	iceberg.RouterDraw = web.RouterDraw
}
