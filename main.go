package main

import (
	_ "weatherapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

