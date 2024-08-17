package bootstrap

import (
	"github.com/Adosh74/geek-blog/pkg/config"
	"github.com/Adosh74/geek-blog/pkg/html"
	"github.com/Adosh74/geek-blog/pkg/routing"
	"github.com/Adosh74/geek-blog/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoute()

	routing.Serve()
}
