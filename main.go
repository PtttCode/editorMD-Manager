package main

import (
	_ "editorMD/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.SetStaticPath("./static", "")

	beego.Run()
}
