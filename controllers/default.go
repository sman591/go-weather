package controllers

import (
	"github.com/astaxie/beego"
  "net/http"
  "io/ioutil"
  "fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "Go Weather!"
	this.Data["Email"] = "stuart@stuartolivera.com"
  resp, err := http.Get("http://weather.yahooapis.com/forecastrss?w=20147")
  if err == nil {
    defer resp.Body.Close()
    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      this.Data["WeatherResponse"] = "failed"
    } else {
      fmt.Printf("%s\n", string(contents))
    }
  } else {
    this.Data["WeatherResponse"] = "failed"
  }
	this.TplNames = "index.tpl"
}
