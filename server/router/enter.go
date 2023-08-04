package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/work"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Work    work.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
