// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/joshuakwan/prometheus-operator/controllers"
	"github.com/astaxie/beego"
)

func init() {
	namespaceRoot := "/api/v1"
	alertmanagerNamespaceRoot := "/alertmanager/"
	prometheusNamespaceRoot := "/prometheus/"

	nsAlertmanager := beego.NewNamespace(namespaceRoot,
		beego.NSNamespace(alertmanagerNamespaceRoot,
			beego.NSInclude(
				&controllers.AlertmanagerController{},
			),
		),
	)
	beego.AddNamespace(nsAlertmanager)

	nsPrometheus := beego.NewNamespace(namespaceRoot,
		beego.NSNamespace(prometheusNamespaceRoot,
			beego.NSInclude(
				&controllers.PrometheusController{},
			),
		),
	)
	beego.AddNamespace(nsPrometheus)
}
