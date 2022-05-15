package web

import "os"

type RoutesApplication struct {
}

func (app *RoutesApplication) HomePath() string {
	path, _ := os.Getwd()
	return path
}
