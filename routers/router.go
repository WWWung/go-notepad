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

	user := new(controllers.UserController)
	user.I = user
	register("/user.api", user)

	_type := new(controllers.TypeController)
	_type.I = _type
	register("/type.api", _type)
}

func register(pattern string, service controllers.IService) {
	http.HandleFunc(pattern, service.HandleHTTPRequest)
}
