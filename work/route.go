package work

import "Go-Web-Frame-Ghost/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}