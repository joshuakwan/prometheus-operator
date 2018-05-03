package controllers

import (
	"net/http"
	"log"
	"github.com/joshuakwan/prometheus-operator/models/prometheus"
	"github.com/joshuakwan/prometheus-operator/models/alertmanager"
)

func init() {
	checkPrometheus()
	checkAlertmanager()
}

func checkPrometheus() {
	log.Println("check the readiness of prometheus")
	if err := checkURLLiveness(prometheusUrl); err != nil {
		log.Println("prometheus unreachable, error:", err)
		panic(err)
	}
	log.Println("prometheus is running at", prometheusUrl)
}

func checkAlertmanager() {
	log.Println("check the readiness of alertmanager")
	if err := checkURLLiveness(alertmanagerUrl); err != nil {
		log.Println("alertmanager unreachable, error:", err)
		panic(err)
	}
	log.Println("alertmanager is running at", alertmanagerUrl)
}

func checkURLLiveness(url string) error {
	_, err := http.Get(url)
	return err
}

func reloadPrometheusService(serviceUrl string) {
	url := serviceUrl + "/-/reload"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Fatalln(serviceUrl, "cannot be reloaded")
	} else {
		if resp.StatusCode == 200 {
			log.Println(serviceUrl, "reloaded")
		} else {
			log.Fatalln(serviceUrl, "fails to get reloaded")
		}
	}
}

func refreshPrometheus() {
	log.Println("write prometheus config to disk")
	if err := prometheus.SaveConfigToFile(prometheusLiveConfig, prometheusConfigFilename); err != nil {
		log.Fatalln("fails to write prometheus config to disk", err)
	} else {
		reloadPrometheusService(prometheusUrl)
	}
}

func refreshAlertmanager() {
	log.Println("write alertmanager config to disk")
	if err := alertmanager.SaveConfigToFile(alertmanagerLiveConfig, alertmanagerConfigFilename); err != nil {
		log.Fatalln("fails to write alertmanager config to disk", err)
	} else {
		reloadPrometheusService(alertmanagerUrl)
	}
}
