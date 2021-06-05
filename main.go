package main

import (
	"editorMD/controllers"
	_ "editorMD/routers"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.SetStaticPath("./static", "")
	web.Router("/savemd", &controllers.DocController{})

	beego.Run()
}
