package routers

import (
	"editorMD/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/savemd", &controllers.DocController{})
}
