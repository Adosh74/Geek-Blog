package bootstrap

import (
	"github.com/Adosh74/geek-blog/pkg/config"
	"github.com/Adosh74/geek-blog/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoute()

	routing.Serve()
}
