package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pkgTest2"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	PkgTest  pkgTest.RouterGroup
	PkgTest2 pkgTest2.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
