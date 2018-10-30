package controllers

import "net/http"

//IService ..
type IService interface {
	HandleHTTPRequest(w http.ResponseWriter, r *http.Request)
}
