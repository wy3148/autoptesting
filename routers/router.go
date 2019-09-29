package routers

import (
	"github.com/astaxie/beego"
	"github.com/wy3148/autoptesting/controllers"
)

func init() {
	beego.Router("/contact/:contact_id", &controllers.ContactController{},"get:Get")
	beego.Router("/contact/", &controllers.ContactController{},"post:Post")
}
