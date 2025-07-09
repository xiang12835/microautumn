package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "microautumn/models"
	_ "microautumn/web/web-routers"
	"path"
	"runtime"
)

func callerSourcePath() string {
	_, callerPath, _, _ := runtime.Caller(1)
	return path.Dir(callerPath)
}

func main() {

	curpath := callerSourcePath()
	static_path := path.Join(curpath, "..", "/", "static")
	template_path := path.Join(curpath, "/web-controllers/templates")

	beego.LoadAppConfig("ini", path.Join(curpath, "..", "/conf/app.conf"))
	beego.SetStaticPath("/static", static_path)
	beego.SetViewsPath(template_path)

	fmt.Println(beego.AppConfig.Int("HttpPort"))
	fmt.Println("[static path]", static_path)
	fmt.Println("[template path]", template_path)

	beego.Run()
}
