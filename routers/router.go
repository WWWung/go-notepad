package routers

import (
	"net/http"

	"../controllers"
)

func init() {
	t := new(controllers.TestController)
	t.I = t
	register("/test.api", t)

	c := new(controllers.CaptchaController)
	c.I = c
	register("/captcha.api", c)
}

func register(pattern string, service controllers.IService) {
	http.HandleFunc(pattern, service.HandleHTTPRequest)
}
