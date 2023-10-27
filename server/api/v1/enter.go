package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/pkgTest2"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	PkgTestApiGroup  pkgTest.ApiGroup
	PkgTest2ApiGroup pkgTest2.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
