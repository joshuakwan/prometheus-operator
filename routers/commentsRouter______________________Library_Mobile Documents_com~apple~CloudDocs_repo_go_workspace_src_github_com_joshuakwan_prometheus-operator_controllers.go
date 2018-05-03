package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "GetTotalConfiguration",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "GetConfiguration",
			Router: `/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "UpdateGlobalConfiguration",
			Router: `/global`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "DeleteGlobalConfiguration",
			Router: `/global/:key`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "AddInhibitionRule",
			Router: `/inhibit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "DeleteInhibitionRule",
			Router: `/inhibit/:index`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "AddReceiver",
			Router: `/receivers`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "DeleteReceiver",
			Router: `/receivers/:name`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "AddRoute",
			Router: `/routes`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:AlertmanagerController"],
		beego.ControllerComments{
			Method: "DeleteRoute",
			Router: `/routes/:index`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:PrometheusController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:PrometheusController"],
		beego.ControllerComments{
			Method: "GetTotalConfiguration",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:PrometheusController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:PrometheusController"],
		beego.ControllerComments{
			Method: "GetConfiguration",
			Router: `/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:PrometheusController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/prometheus-operator/controllers:PrometheusController"],
		beego.ControllerComments{
			Method: "AddScrapeConfig",
			Router: `/scrapes`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
