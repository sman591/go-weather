package controllers

import (
	"github.com/astaxie/beego"
  "net/http"
  "io/ioutil"
  "encoding/xml"
)

type MainController struct {
	beego.Controller
}

type Query struct {
  rss Rss
}

type Rss struct {
  channel Channel
}

type Channel struct {
  title             string
  description       string
  ywlocation        string
  ywunitspeed       string
  ywunittemperature string
  item              Item
}

type Item struct {
  ywcondition Condition `xml:"yweather:condition"`
}

type Condition struct {
  ywtemp string `xml:"temp,attr"`
  ywtext string `xml:"text,attr"`
}

func (this *MainController) Get() {
	this.Data["Website"] = "Go Weather!"
	this.Data["Email"] = "stuart@stuartolivera.com"
  resp, err := http.Get("http://weather.yahooapis.com/forecastrss?w=14623")

  this.Data["WeatherLocation"] = "(failed to get weather)"

  if err == nil {
    defer resp.Body.Close()
    contents, err := ioutil.ReadAll(resp.Body)
    if err == nil {
      var q Query
      xml.Unmarshal(contents, &q)

      println(q.rss.channel.description);

      this.Data["WeatherLocation"]        = q.rss.channel.ywlocation
      this.Data["WeatherTemperature"]     = q.rss.channel.item.ywcondition.ywtemp
      this.Data["WeatherCondition"]       = q.rss.channel.item.ywcondition.ywtext
      this.Data["WeatherUnitSpeed"]       = q.rss.channel.ywunitspeed
      this.Data["WeatherUnitTemperature"] = q.rss.channel.ywunittemperature
    }
  }
	this.TplNames = "index.tpl"
}
