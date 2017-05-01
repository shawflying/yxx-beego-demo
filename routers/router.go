package routers

import (
	"mohoo-activity-bigwheel/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/bigwheel/getName", &controllers.GetNameController{})
    beego.Router("/bigwheel/rule", &controllers.RuleHomeController{})
    beego.Router("/bigwheel/list", &controllers.ListHomeController{})
    beego.Router("/bigwheel/home", &controllers.MainController{})
    beego.Router("/", &controllers.MainController{})
}
