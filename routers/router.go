package routers

import (
	"net/http"

	"../controllers"
)

func init() {
	t := new(controllers.TestController)
	t.I = t
	register("/test.api", t)

	captcha := new(controllers.CaptchaController)
	captcha.I = captcha
	register("/captcha.api", captcha)

	user := new(controllers.UserController)
	user.I = user
	register("/user.api", user)

	_type := new(controllers.TypeController)
	_type.I = _type
	register("/type.api", _type)

	note := new(controllers.NoteController)
	note.I = note
	register("/note.api", note)
}

func register(pattern string, service controllers.IService) {
	http.HandleFunc(pattern, service.HandleHTTPRequest)
}
