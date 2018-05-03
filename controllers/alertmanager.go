package controllers

import (
	"github.com/astaxie/beego"
	prom_alertmanager "github.com/prometheus/alertmanager/config"
	"log"
	"encoding/json"
	"github.com/joshuakwan/prometheus-operator/models/alertmanager"
	"strconv"
)

var (
	alertmanagerUrl            = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::alertmanager_url")
	alertmanagerConfigFilename = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::alertmanager_config")
	alertmanagerLiveConfig     = getAlertmanagerConfig(alertmanagerConfigFilename)
)

func getAlertmanagerConfig(filename string) *prom_alertmanager.Config {
	log.Println("load alertmanager configuration from " + filename)
	cfg, _, err := prom_alertmanager.LoadFile(filename)
	if err != nil {
		panic(err)
	}
	return cfg
}

type AlertmanagerController struct {
	beego.Controller
}

// @Title GetTotalConfiguration
// @Description Get total configuration
// @router / [get]
func (c *AlertmanagerController) GetTotalConfiguration() {
	c.Data["json"] = alertmanagerLiveConfig
	c.ServeJSON()
}

// @Title GetConfiguration
// @Description Get current configuration
// @router /:key [get]
// @Param key path string true "the configuration key"
func (c *AlertmanagerController) GetConfiguration() {
	key := c.GetString(":key")
	log.Println("get the configuration:", key)
	if key == "" {
		c.Data["json"] = alertmanagerLiveConfig
	} else {
		switch key {
		case "global":
			c.Data["json"] = alertmanagerLiveConfig.Global
		case "receivers":
			c.Data["json"] = alertmanagerLiveConfig.Receivers
		case "routes":
			c.Data["json"] = alertmanagerLiveConfig.Route
		case "inhibit":
			c.Data["json"] = alertmanagerLiveConfig.InhibitRules
		default:
			c.CustomAbort(400, "The path should be global, receivers, routes, inhibit or leave empty")
		}
	}

	c.ServeJSON()
}

// @Title UpdateGlobalConfiguration
// @Description update global settings, partial update is supported
// @router /global [put]
func (c *AlertmanagerController) UpdateGlobalConfiguration() {
	currentConfig := alertmanagerLiveConfig.Global
	var newGlobal prom_alertmanager.GlobalConfig

	body := c.Ctx.Input.RequestBody
	log.Println(string(body))

	err := json.Unmarshal(body, &newGlobal)
	if err != nil {
		c.CustomAbort(400, "Invalid JSON object")
	}

	log.Println(newGlobal)

	alertmanager.Update(currentConfig, &newGlobal)
	log.Println(currentConfig)

	go refreshAlertmanager()

	c.Data["json"] = currentConfig
	c.ServeJSON()
}

// @Title Delete
// @Description delete an item in global settings with the specified key
// @router /global/:key [delete]
func (c *AlertmanagerController) DeleteGlobalConfiguration() {
	currentConfig := alertmanagerLiveConfig.Global
	key := c.GetString(":key")
	log.Println("delete the global configuration:", key)

	alertmanager.Delete(currentConfig, key)

	go refreshAlertmanager()

	c.Data["json"] = currentConfig
	c.ServeJSON()
}

// @Title AddInhibitionRule
// @Description add an inhibition rule
// @router /inhibit [post]
func (c *AlertmanagerController) AddInhibitionRule() {
	currentConfig := alertmanagerLiveConfig.InhibitRules

	var newRule prom_alertmanager.InhibitRule

	body := c.Ctx.Input.RequestBody
	log.Println(string(body))

	err := json.Unmarshal(body, &newRule)
	if err != nil {
		c.CustomAbort(400, "Invalid JSON object")
	}
	log.Println(&newRule)

	alertmanagerLiveConfig.InhibitRules = alertmanager.AddInhibitRule(currentConfig, &newRule)
	log.Println(alertmanagerLiveConfig.InhibitRules)

	go refreshAlertmanager()

	c.Data["json"] = alertmanagerLiveConfig.InhibitRules
	c.ServeJSON()
}

// @Title DeleteInhibitionRule
// @Description delete an inhibition rule at a certain index (start at 0)
// @router /inhibit/:index [delete]
func (c *AlertmanagerController) DeleteInhibitionRule() {
	currentConfig := alertmanagerLiveConfig.InhibitRules

	index, err := strconv.Atoi(c.GetString(":index"))
	if err != nil {
		c.CustomAbort(400, "Invalid index value "+c.GetString(":index"))
	} else {
		rules, err := alertmanager.RemoveInhibitRule(currentConfig, index)
		if err != nil {
			c.CustomAbort(400, "Index "+strconv.Itoa(index)+" not in the right range")
		} else {
			alertmanagerLiveConfig.InhibitRules = rules
			log.Println(alertmanagerLiveConfig.InhibitRules)
			go refreshAlertmanager()
			c.Data["json"] = alertmanagerLiveConfig.InhibitRules
		}
	}

	c.ServeJSON()
}

// @Title AddReceiver
// @Description add a new receiver
// @router /receivers [post]
func (c *AlertmanagerController) AddReceiver() {
	currentConfig := alertmanagerLiveConfig.Receivers
	log.Println(len(currentConfig))

	var newReceiver prom_alertmanager.Receiver

	body := c.Ctx.Input.RequestBody
	log.Println(string(body))

	err := json.Unmarshal(body, &newReceiver)
	if err != nil {
		c.CustomAbort(400, "Invalid JSON object")
	}
	log.Println(&newReceiver)

	receivers, err := alertmanager.AddReceiver(currentConfig, &newReceiver)

	if err != nil {
		c.CustomAbort(400, "Receiver "+newReceiver.Name+" already exists")
	} else {
		alertmanagerLiveConfig.Receivers = receivers
		go refreshAlertmanager()
		c.Data["json"] = alertmanagerLiveConfig.Receivers
	}

	c.ServeJSON()
}

// @Title DeleteReceiver
// @Description delete a receiver by name
// @router /receivers/:name [delete]
func (c *AlertmanagerController) DeleteReceiver() {
	currentConfig := alertmanagerLiveConfig.Receivers

	name := c.GetString(":name")
	log.Println("receiver's name to delete: " + name)

	receivers, err := alertmanager.RemoveReceiver(currentConfig, name)

	if err != nil {
		c.CustomAbort(400, "Receiver "+name+" not found")
	} else {
		alertmanagerLiveConfig.Receivers = receivers
		go refreshAlertmanager()
		c.Data["json"] = alertmanagerLiveConfig.Receivers
	}

	c.ServeJSON()
}

// @Title AddRoute
// @Description add a new sub route (now supports only layer 1)
// @router /routes [post]
func (c *AlertmanagerController) AddRoute() {
	currentConfig := alertmanagerLiveConfig.Route

	var newSubroute prom_alertmanager.Route

	body := c.Ctx.Input.RequestBody
	log.Println(string(body))

	err := json.Unmarshal(body, &newSubroute)
	if err != nil {
		c.CustomAbort(400, "Invalid JSON object")
	}
	log.Println(&newSubroute)

	alertmanagerLiveConfig.Route = alertmanager.AddSubroute(currentConfig, &newSubroute)
	log.Println(alertmanagerLiveConfig.Route)

	go refreshAlertmanager()

	c.Data["json"] = alertmanagerLiveConfig.Route
	c.ServeJSON()
}

// @Title DeleteRoute
// @Description delete a sub route at a certain index (starts at 0, now supports only layer 1)
// @router /routes/:index [delete]
func (c *AlertmanagerController) DeleteRoute() {
	currentConfig := alertmanagerLiveConfig.Route

	index, err := strconv.Atoi(c.GetString(":index"))
	if err != nil {
		c.CustomAbort(400, "Invalid index value "+c.GetString(":index"))
	} else {
		route, err := alertmanager.RemoveSubroute(currentConfig, index)
		if err != nil {
			c.CustomAbort(400, "Index "+strconv.Itoa(index)+" not in the right range")
		} else {
			alertmanagerLiveConfig.Route = route
			log.Println(alertmanagerLiveConfig.Route)
			go refreshAlertmanager()
			c.Data["json"] = alertmanagerLiveConfig.Route
		}
	}

	c.ServeJSON()
}
