package routers

import (
	"net/http"

	"../controllers"
)

func init() {
	t := new(controllers.TestController)
	t.I = t
	register("/test.api", t)
}

func register(pattern string, service controllers.IService) {
	http.HandleFunc("/test.api", service.HandleHTTPRequest)
}
