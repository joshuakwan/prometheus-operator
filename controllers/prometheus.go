package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"github.com/joshuakwan/prometheus-operator/models/prometheus"
	"encoding/json"
)

var (
	prometheusUrl            = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::prometheus_url")
	prometheusConfigFilename = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::prometheus_config")
	prometheusLiveConfig     = getPrometheusConfig(prometheusConfigFilename)
)

func getPrometheusConfig(filename string) *prometheus.Config {
	log.Println("load prometheus configuration from " + filename)
	cfg, err := prometheus.LoadFile(filename)
	if err != nil {
		panic(err)
	}
	return cfg
}

type PrometheusController struct {
	beego.Controller
}

// @Title GetTotalConfiguration
// @Description Get total configuration
// @router / [get]
func (c *PrometheusController) GetTotalConfiguration() {
	c.Data["json"] = prometheusLiveConfig
	c.ServeJSON()
}

// @Title GetConfiguration
// @Description Get current configuration
// @router /:key [get]
// @Param key path string true "the configuration key"
func (c *PrometheusController) GetConfiguration() {
	key := c.GetString(":key")
	log.Println("get the configuration:", key)
	if key == "" {
		c.Data["json"] = prometheusLiveConfig
	} else {
		switch key {
		case "global":
			c.Data["json"] = prometheusLiveConfig.Global
		case "rules":
			c.Data["json"] = prometheusLiveConfig.RuleFiles
		case "scrapes":
			c.Data["json"] = prometheusLiveConfig.ScrapeConfigs
		case "alerting":
			c.Data["json"] = prometheusLiveConfig.Alerting
		default:
			c.CustomAbort(400, "The path should be global, receivers, routes, inhibit or leave empty")
		}
	}

	c.ServeJSON()
}

// @Title AddScrapeConfig
// @Description Add a new scrape configuration
// @router /scrapes [post]
func (c *PrometheusController) AddScrapeConfig() {
	body := c.Ctx.Input.RequestBody
	log.Println(string(body))

	var newScrape prometheus.ScrapeConfig
	if err := json.Unmarshal(body, &newScrape); err != nil {
		c.CustomAbort(400, err.Error())
	}

	prometheusLiveConfig.ScrapeConfigs = prometheus.AddScrapeConfig(prometheusLiveConfig.ScrapeConfigs, &newScrape)
	go refreshPrometheus()
}
